// Command check validates the packages of this source against docs/format.md
// of codicus-labs/agent-package-manager, until `apm package check` exists.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/bmatcuk/doublestar/v4"
	"github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/tailscale/hujson"
)

var (
	packageName  = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*$`)
	strengthMark = regexp.MustCompile(`^\*\*(Must|Prefer)\.\*\*`)
	firstRelease = semver.MustParse("1.0.0")
	tightens     = regexp.MustCompile("(?:Tightens|Builds on) `bp:([a-z0-9-]+)/([a-z0-9/-]+)`")
)

// config is scripts/check.jsonc: what this source may depend on and what its
// instructions must not say.
type config struct {
	// Sources maps the git URL of another source, exactly as dependencies
	// write it, to a local checkout used to check that the package exists.
	Sources map[string]string `json:"sources"`
	// TightensSource is the URL whose instructions a "Tightens `bp:<pkg>/<id>`"
	// line refers to.
	TightensSource string `json:"tightensSource"`
	// ForbiddenWords fail the check when an instruction contains them.
	ForbiddenWords []string `json:"forbiddenWords"`
	// ChoiceWords are reported as warnings: they often mark a tool choice.
	ChoiceWords []string `json:"choiceWords"`
}

type manifest struct {
	Description  string            `json:"description"`
	Dependencies map[string]string `json:"dependencies"`
}

type summary struct {
	name                     string
	always, onDemand, skills int
	deps                     []string
}

type checker struct {
	root     string
	cfg      config
	errs     []string
	warnings []string
}

func (c *checker) warn(format string, args ...any) {
	c.warnings = append(c.warnings, fmt.Sprintf(format, args...))
}

func (c *checker) fail(format string, args ...any) {
	c.errs = append(c.errs, fmt.Sprintf(format, args...))
}

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	c := &checker{root: root}
	c.loadConfig()
	sums := c.run(root)
	printSummary(sums)
	sort.Strings(c.warnings)
	for _, w := range c.warnings {
		fmt.Fprintln(os.Stderr, "warning:", w)
	}
	if len(c.errs) > 0 {
		sort.Strings(c.errs)
		for _, e := range c.errs {
			fmt.Fprintln(os.Stderr, "error:", e)
		}
		fmt.Fprintf(os.Stderr, "%d errors\n", len(c.errs))
		os.Exit(1)
	}
	fmt.Println("ok")
}

func (c *checker) loadConfig() {
	raw, err := os.ReadFile(filepath.Join(c.root, "scripts", "check.jsonc"))
	if err != nil {
		return
	}
	std, err := hujson.Standardize(raw)
	if err == nil {
		err = json.Unmarshal(std, &c.cfg)
	}
	if err != nil {
		c.fail("scripts/check.jsonc: %v", err)
	}
}

// localSource returns the checkout of the source a URL names, relative to the root.
func (c *checker) localSource(url string) (string, bool) {
	dir, ok := c.cfg.Sources[url]
	if !ok {
		return "", false
	}
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(c.root, dir)
	}
	return dir, true
}

func (c *checker) run(root string) []summary {
	schema, err := compileSchema(filepath.Join(root, "schema", "manifest.schema.json"))
	if err != nil {
		c.fail("schema: %v", err)
		return nil
	}
	pkgDir := filepath.Join(root, "packages")
	entries, err := os.ReadDir(pkgDir)
	if err != nil {
		c.fail("%v", err)
		return nil
	}
	manifests := map[string]manifest{}
	var sums []summary
	skillOwner := map[string]string{}
	for _, e := range entries {
		if !e.IsDir() {
			c.fail("packages/%s: only package directories belong in packages/", e.Name())
			continue
		}
		name := e.Name()
		if !packageName.MatchString(name) {
			c.fail("packages/%s: name must be lowercase letters, digits and dashes", name)
		}
		dir := filepath.Join(pkgDir, name)
		m, ok := c.checkManifest(name, dir, schema)
		if ok {
			manifests[name] = m
		}
		s := summary{name: name}
		s.always, s.onDemand = c.checkInstructions(name, dir)
		s.skills = c.checkSkills(name, dir, skillOwner)
		c.checkMigrations(name, dir)
		c.checkLayout(name, dir)
		if s.always+s.onDemand+s.skills == 0 && len(m.Dependencies) == 0 {
			c.fail("%s: a package without content is a preset and needs dependencies", name)
		}
		for d := range m.Dependencies {
			s.deps = append(s.deps, d)
		}
		sort.Strings(s.deps)
		sums = append(sums, s)
	}
	c.checkDependencies(manifests)
	return sums
}

func compileSchema(path string) (*jsonschema.Schema, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	doc, err := jsonschema.UnmarshalJSON(f)
	if err != nil {
		return nil, err
	}
	comp := jsonschema.NewCompiler()
	if err := comp.AddResource("manifest.schema.json", doc); err != nil {
		return nil, err
	}
	return comp.Compile("manifest.schema.json")
}

func (c *checker) checkManifest(name, dir string, schema *jsonschema.Schema) (manifest, bool) {
	var m manifest
	raw, err := os.ReadFile(filepath.Join(dir, "manifest.jsonc"))
	if err != nil {
		c.fail("%s: %v", name, err)
		return m, false
	}
	std, err := hujson.Standardize(raw)
	if err != nil {
		c.fail("%s/manifest.jsonc: not JSONC: %v", name, err)
		return m, false
	}
	inst, err := jsonschema.UnmarshalJSON(bytes.NewReader(std))
	if err != nil {
		c.fail("%s/manifest.jsonc: %v", name, err)
		return m, false
	}
	if err := schema.Validate(inst); err != nil {
		c.fail("%s/manifest.jsonc: does not fit the schema: %s", name, strings.Join(strings.Fields(fmt.Sprintf("%#v", err)), " "))
		return m, false
	}
	if err := json.Unmarshal(std, &m); err != nil {
		c.fail("%s/manifest.jsonc: %v", name, err)
		return m, false
	}
	return m, true
}

func (c *checker) checkDependencies(manifests map[string]manifest) {
	for name, m := range manifests {
		for dep, rng := range m.Dependencies {
			con, err := semver.NewConstraint(rng)
			if err != nil {
				c.fail("%s: dependency %s has an invalid range %q: %v", name, dep, rng, err)
			}
			if url, pkg, ok := strings.Cut(dep, "#"); ok {
				dir, known := c.localSource(url)
				if !known {
					c.fail("%s: dependency %s names a source this one may not depend on (scripts/check.jsonc)", name, dep)
					continue
				}
				if _, err := os.Stat(filepath.Join(dir, "packages", pkg, "manifest.jsonc")); err != nil {
					c.fail("%s: dependency %s does not exist in %s", name, dep, dir)
				}
				continue
			}
			if _, ok := manifests[dep]; !ok {
				c.fail("%s: dependency %s does not exist", name, dep)
			} else if con != nil && !c.admitsNextRelease(dep, con) {
				c.fail("%s: range %s of %s admits neither its first release, its newest nor its next major", name, rng, dep)
			}
		}
	}
	// Cycles.
	state := map[string]int{}
	var visit func(string, []string)
	visit = func(n string, stack []string) {
		switch state[n] {
		case 1:
			c.fail("dependency cycle: %s", strings.Join(append(stack, n), " -> "))
			return
		case 2:
			return
		}
		state[n] = 1
		for d := range manifests[n].Dependencies {
			if _, ok := manifests[d]; ok {
				visit(d, append(stack, n))
			}
		}
		state[n] = 2
	}
	for n := range manifests {
		visit(n, nil)
	}
}

// admitsNextRelease reports whether a range admits a release the package has
// or is about to get: 1.0.0, its newest tag, or the major after that.
func (c *checker) admitsNextRelease(pkg string, con *semver.Constraints) bool {
	if con.Check(firstRelease) {
		return true
	}
	out, err := exec.Command("git", "-C", c.root, "tag", "--list", pkg+"@*").Output()
	if err != nil {
		return false
	}
	var newest *semver.Version
	for _, t := range strings.Fields(string(out)) {
		v, err := semver.StrictNewVersion(strings.TrimPrefix(t, pkg+"@"))
		if err == nil && (newest == nil || v.GreaterThan(newest)) {
			newest = v
		}
	}
	if newest == nil {
		return false
	}
	next := newest.IncMajor()
	return con.Check(newest) || con.Check(&next)
}

// checkContent applies the source's rules on what instructions may say.
func (c *checker) checkContent(where, body string) {
	lower := strings.ToLower(body)
	for _, w := range c.cfg.ForbiddenWords {
		if regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(w) + `\b`).MatchString(body) {
			c.fail("%s: says %q, which does not belong in this source", where, w)
		}
	}
	for _, w := range c.cfg.ChoiceWords {
		if strings.Contains(lower, strings.ToLower(w)) {
			c.warn("%s: says %q; make sure it describes how to use a tool, not which one to choose", where, w)
		}
	}
	for _, m := range tightens.FindAllStringSubmatch(body, -1) {
		dir, ok := c.localSource(c.cfg.TightensSource)
		if !ok {
			c.fail("%s: tightens bp:%s/%s, but scripts/check.jsonc names no tightensSource", where, m[1], m[2])
			continue
		}
		if _, err := os.Stat(filepath.Join(dir, "packages", m[1], "instructions", m[2]+".md")); err != nil {
			c.fail("%s: tightens bp:%s/%s, which does not exist in %s", where, m[1], m[2], dir)
		}
	}
}

// frontMatter splits a Markdown file into its YAML front matter and body.
func frontMatter(raw []byte) (map[string]any, string, error) {
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	if !strings.HasPrefix(text, "---\n") {
		return nil, "", fmt.Errorf("no front matter")
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		return nil, "", fmt.Errorf("front matter not closed")
	}
	fm := map[string]any{}
	if err := yaml.UnmarshalWithOptions([]byte(text[4:4+end]), &fm, yaml.Strict()); err != nil {
		return nil, "", err
	}
	return fm, strings.TrimSpace(text[4+end+5:]), nil
}

func oneLine(v any) (string, bool) {
	s, ok := v.(string)
	return s, ok && strings.TrimSpace(s) != "" && !strings.Contains(s, "\n")
}

func (c *checker) checkInstructions(pkg, dir string) (always, onDemand int) {
	base := filepath.Join(dir, "instructions")
	if _, err := os.Stat(base); err != nil {
		return 0, 0
	}
	_ = filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(dir, path)
		where := pkg + "/" + rel
		if !strings.HasSuffix(path, ".md") {
			c.fail("%s: only Markdown files belong in instructions/", where)
			return nil
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			c.fail("%s: %v", where, err)
			return nil
		}
		fm, body, err := frontMatter(raw)
		if err != nil {
			c.fail("%s: front matter: %v", where, err)
			return nil
		}
		for k := range fm {
			if k != "description" && k != "load" && k != "paths" {
				c.fail("%s: unknown front matter field %q", where, k)
			}
		}
		if _, ok := oneLine(fm["description"]); !ok {
			c.fail("%s: description must be one non-empty line", where)
		}
		switch fm["load"] {
		case "always":
			always++
			if _, ok := fm["paths"]; ok {
				c.fail("%s: paths are for on-demand instructions only", where)
			}
		case "on-demand":
			onDemand++
			if p, ok := fm["paths"]; ok {
				c.checkPaths(where, p)
			}
			if !strengthMark.MatchString(body) {
				c.fail("%s: body must start with **Must.** or **Prefer.**", where)
			}
		default:
			c.fail("%s: load must be always or on-demand, got %v", where, fm["load"])
		}
		if body == "" {
			c.fail("%s: empty body", where)
		}
		desc, _ := fm["description"].(string)
		c.checkContent(where, desc+"\n"+body)
		return nil
	})
	return always, onDemand
}

func (c *checker) checkPaths(where string, p any) {
	list, ok := p.([]any)
	if !ok || len(list) == 0 {
		c.fail("%s: paths must be a non-empty list", where)
		return
	}
	for _, item := range list {
		s, ok := item.(string)
		if !ok || s == "" {
			c.fail("%s: every path must be a non-empty string", where)
			continue
		}
		if strings.HasPrefix(s, "/") || strings.HasPrefix(s, "./") || strings.Contains(s, "..") {
			c.fail("%s: path %q must be relative to the repository root", where, s)
		}
		if !doublestar.ValidatePattern(s) {
			c.fail("%s: invalid glob %q", where, s)
		}
		if s == "**" || s == "**/*" {
			c.fail("%s: path %q matches every file; leave paths out instead", where, s)
		}
	}
}

func (c *checker) checkSkills(pkg, dir string, owner map[string]string) int {
	entries, err := os.ReadDir(filepath.Join(dir, "skills"))
	if err != nil {
		return 0
	}
	n := 0
	for _, e := range entries {
		where := pkg + "/skills/" + e.Name()
		if !e.IsDir() {
			c.fail("%s: skills/ holds only skill directories", where)
			continue
		}
		n++
		raw, err := os.ReadFile(filepath.Join(dir, "skills", e.Name(), "SKILL.md"))
		if err != nil {
			c.fail("%s: missing SKILL.md", where)
			continue
		}
		fm, _, err := frontMatter(raw)
		if err != nil {
			c.fail("%s/SKILL.md: front matter: %v", where, err)
			continue
		}
		if fm["name"] != e.Name() {
			c.fail("%s/SKILL.md: name must be %q", where, e.Name())
		}
		if _, ok := oneLine(fm["description"]); !ok {
			c.fail("%s/SKILL.md: description must be one non-empty line", where)
		}
		if other, ok := owner[e.Name()]; ok {
			c.fail("skill %s is in both %s and %s", e.Name(), other, pkg)
		}
		owner[e.Name()] = pkg
	}
	return n
}

func (c *checker) checkMigrations(pkg, dir string) {
	entries, err := os.ReadDir(filepath.Join(dir, "migrations"))
	if err != nil {
		return
	}
	for _, e := range entries {
		where := pkg + "/migrations/" + e.Name()
		v, err := semver.StrictNewVersion(strings.TrimSuffix(e.Name(), ".md"))
		if err != nil || !strings.HasSuffix(e.Name(), ".md") || v.Patch() != 0 {
			c.fail("%s: must be <major>.<minor>.0.md", where)
			continue
		}
		raw, err := os.ReadFile(filepath.Join(dir, "migrations", e.Name()))
		if err != nil {
			c.fail("%s: %v", where, err)
			continue
		}
		fm, _, err := frontMatter(raw)
		if err != nil {
			c.fail("%s: front matter: %v", where, err)
			continue
		}
		if _, ok := oneLine(fm["description"]); !ok {
			c.fail("%s: description must be one non-empty line", where)
		}
	}
}

func (c *checker) checkLayout(pkg, dir string) {
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		switch e.Name() {
		case "manifest.jsonc", "instructions", "skills", "migrations":
		default:
			c.fail("%s/%s: not part of a package (manifest.jsonc, instructions/, skills/, migrations/)", pkg, e.Name())
		}
	}
}

func printSummary(sums []summary) {
	fmt.Printf("%-24s %6s %9s %6s  %s\n", "package", "always", "on-demand", "skills", "dependencies")
	for _, s := range sums {
		fmt.Printf("%-24s %6d %9d %6d  %s\n", s.name, s.always, s.onDemand, s.skills, strings.Join(s.deps, ", "))
	}
}
