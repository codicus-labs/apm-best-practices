---
description: "sqlx usage: checked query macros, committed offline data in .sqlx, migrations at startup"
load: on-demand
paths: ["**/migrations/*.sql", "**/.sqlx/**", "**/*store*.rs", "**/*repo*.rs", "**/db/**/*.rs", "**/database/**/*.rs"]
---

**Prefer.** When a Rust service uses sqlx, queries go through the checked macros `query!` and `query_as!`, their offline data in `.sqlx/` is committed (`cargo sqlx prepare`) so builds need no database, and migrations in `migrations/` run at startup with `sqlx::migrate!()`.
