---
description: "Go dev tools (code generators, linters) pinned with the tool directive in go.mod and run with go tool"
load: on-demand
paths: ["**/go.mod", "**/tools.go"]
---

**Prefer.** A development tool written in Go that the repository runs, such as a code generator, is listed with a `tool` directive in `go.mod` (added with `go get -tool <module>/cmd/<name>`) and run with `go tool <name>`, so `go.sum` pins it together with the module's dependencies. There is no `tools.go` with blank imports and no `go install` of such a tool in scripts or CI.

Exception: a tool whose documentation advises against building it from source (golangci-lint) is installed as a released binary.
