# AGENTS.md

The packages of an APM source of best practices. The format is `docs/format.md` of `codicus-labs/agent-package-manager`; follow it exactly. Terms come from its `GLOSSARY.md`: instruction, package, preset, skill (not rule, rulebook, profile).

## What belongs here

Only rules everybody would follow: how to use a language, framework or tool well once a repository uses it. Never which tool, library or service to choose, and nothing about one organisation; those go into an organisation's own source that depends on these packages (`codicus-labs/apm-packages` is one). An instruction about a tool says when it applies ("When the repository uses sqlc, …") and names the tool's files in `paths`, so it reaches only repositories that use the tool.

## Language

Write all code and documentation in this repository in English: instructions, skills, manifests, scripts, commit messages and the README.

## Adding a package

1. Check `README.md` for a package that already covers the stack or concern; extend that one instead.
2. A package is a kind of project (`go-server`, `node-library`), a framework (`nuxt`) or a tool used across kinds with at least two rules (`docker`, `pnpm`). A tool with one rule gets a conditional instruction in the kind package that uses it, or in `base`.
3. Write `packages/<name>/manifest.jsonc` with the `$schema` line, a one-line `description` and `dependencies` on packages of this repository by name, each `"^1.0.0"` until they have later majors. This source depends on no other and has no presets.
4. Add its instructions (below) and a row to the package table in `README.md`.

## Adding an instruction

One concern per file, `packages/<package>/instructions/<concern>.md`:

```markdown
---
description: Pinned base images in Dockerfiles
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile"]
---

**Must.** Every `FROM` names a version tag and its digest …
```

- `description`: one line, the words an agent would search for. Quote it when it contains `: `.
- `load`: `on-demand` by default. `always` only for a short instruction every session in a subscribing repository needs, a few per package.
- `paths`: globs of exactly the files the instruction concerns, relative to the repository root or a subscription's scope, so start them with `**/` unless the file sits at the root. Leave `paths` out when no file pattern fits; the description then finds it.
- The body starts with its strength, **Must.** (departing needs the owner's consent) or **Prefer.** (departing needs a written reason), then states what the code does in terms a reviewer can check in a file. Most instructions here are **Prefer.**; **Must.** is for what is wrong everywhere (security, data loss, broken builds). Name no versions.
- Skills go in `packages/<package>/skills/<name>/SKILL.md` with front matter `name` and `description`; skill names are unique across packages.

## Checking and releasing

Run `apm package check` after every change; it is green before the work counts as done. The instructions for writing packages come from the package `authoring` of `codicus-labs/agent-package-manager`, installed here with `apm install` (`apm.jsonc`). A release is the annotated tag `<package>@<version>`; tag only after the check passes, and leave pushing to the owner. A change to what a released instruction asks is a minor or major release with a migration (see the `authoring` package).

Commits follow Conventional Commits.

<!-- apm:begin -->
<!-- Written by `apm install` from apm-lock.json: change the packages in apm.jsonc, not this block. -->

## From APM packages (apm/authoring@1.0.0)

A package is `packages/<name>/` with a `manifest.jsonc` (description, dependencies), `instructions/**/*.md`, `skills/<skill>/SKILL.md` and `migrations/<version>.md`; subscribing repositories install it with `apm install`. Before writing or changing an instruction, run `apm context --id` on, or read, the instructions on instruction files and on how instructions load; write one with the `write-instruction` skill and a migration with `write-migration`.

A release is the tag `<name>@<version>`: a patch changes wording only, a minor adds or asks more, a major removes, renames or changes meaning. `apm package check` is green before work counts as done; `apm package release <name> <bump>` tags, and pushing the tag publishes.

### On-demand instructions

This repository has more instructions for specific files and topics than this file holds. Before you read, create or change files, run `apm context <file>...` with those files and follow what it prints for every change to them; `apm context --search <words>` finds instructions about a topic, `apm context --list` lists all of them, and `apm context --id <id>` prints one. Instructions that appear next to tool results come from the same place.
<!-- apm:end -->
