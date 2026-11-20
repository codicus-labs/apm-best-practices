---
description: "Go tests against a database container (Testcontainers): one container per test binary, a migrated database per test"
load: on-demand
paths: ["**/*_test.go", "**/*dbtest*/**", "**/*testutil*/**"]
---

**Prefer.** When Go tests start their database with Testcontainers, one helper package exposes `New(t *testing.T)`: it starts the container once per test binary (guarded by `sync.Once`), creates a database under a random name for the calling test, migrates it through the same code that opens the database in production, and drops it in `t.Cleanup` (`DROP DATABASE … WITH (FORCE)` on PostgreSQL).
