---
description: Layout of a Go module with cmd/ and internal/
load: on-demand
paths: ["**/go.mod", "**/cmd/**/*.go", "**/internal/**/*.go"]
---

**Prefer.** Each binary has `cmd/<name>/main.go`, which only wires dependencies and starts the program. All other code lives in `internal/`, by domain (`internal/order`, `internal/postgres`). A `pkg/` directory exists only for code other modules import.
