---
description: "Logging with tracing in Rust services: JSON output, RUST_LOG, instrumented request functions"
load: on-demand
paths: ["**/*.rs"]
---

**Prefer.** When a Rust service logs through `tracing`, `tracing-subscriber` writes JSON in production with the level from `RUST_LOG`, and request-level functions carry `#[tracing::instrument(skip(state))]`. `println!` and `eprintln!` appear in CLIs' user-facing output, not in service logs.
