---
description: Table-driven Go tests with the standard testing package and t.Run subtests
load: on-demand
paths: ["**/*_test.go"]
---

**Prefer.** Tests sit next to the code in `<file>_test.go` and use the standard `testing` package. A function tested with several inputs gets one table of cases, each run as a subtest with `t.Run(tc.name, …)`.
