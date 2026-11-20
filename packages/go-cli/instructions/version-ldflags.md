---
description: Version of a Go command-line tool set at build time with -ldflags -X (--version)
load: on-demand
paths: ["**/cmd/**/*.go", "**/*version*.go", "**/.goreleaser.yaml", "**/.goreleaser.yml"]
---

**Prefer.** A Go command-line tool declares `var version = "dev"` in one package, prints it for `--version`, and release builds set it with `-ldflags "-X <import path>.version=<version>"`. Nothing else holds the version, so a local build says `dev` and a release says its tag.
