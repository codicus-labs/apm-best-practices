---
description: Async Rust on Tokio without blocking calls
load: on-demand
paths: ["**/*.rs"]
---

**Must.** When a crate uses Tokio, an `async fn` never blocks: files go through `tokio::fs`, waiting through `tokio::time::sleep`, CPU-heavy or blocking work through `tokio::task::spawn_blocking`.
