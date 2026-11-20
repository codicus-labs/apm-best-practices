---
description: Commands a Go change passes before it counts as done (go test -race, linter, go mod tidy)
load: on-demand
paths: ["**/*.go", "**/go.mod", "**/go.sum"]
---

**Must.** Before a Go change counts as done, `go test -race ./...` and the repository's linter pass (`go vet ./...` in a repository without a linter), and `go mod tidy -diff` reports no difference. CI runs all three and fails when any of them does.
