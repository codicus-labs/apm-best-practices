---
description: context.Context as the first parameter in Go
load: on-demand
paths: ["**/*.go"]
---

**Must.** A function that does I/O or may block takes `ctx context.Context` as its first parameter and passes it on. A context is never stored in a struct field, and `context.Background()` appears only in `main` and tests.
