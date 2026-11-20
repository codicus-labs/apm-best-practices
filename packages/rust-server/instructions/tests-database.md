---
description: "Rust database tests with #[sqlx::test] against a database container (Testcontainers)"
load: on-demand
paths: ["**/tests/**/*.rs"]
---

**Prefer.** When Rust tests start their database with Testcontainers, a helper in `tests/common/` starts the container once per test binary and sets `DATABASE_URL` before the first test runs, and database tests use `#[sqlx::test]`, which gives each test its own migrated database.
