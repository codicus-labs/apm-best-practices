---
description: Release builds of a Go command-line tool (tags, static binaries, CGO_ENABLED=0, -trimpath)
load: on-demand
paths: ["**/.goreleaser.yaml", "**/.goreleaser.yml", ".github/workflows/*.yml", ".github/workflows/*.yaml"]
---

**Prefer.** When a Go command-line tool publishes binaries, they are built by CI from a pushed tag `v*` once the checks pass, with `CGO_ENABLED=0` and `-trimpath`, so they are static and carry no paths of the build machine. Every release lists checksums of its archives.
