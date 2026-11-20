---
description: Commands a Rust change passes before it counts as done
load: on-demand
paths: ["**/*.rs", "**/Cargo.toml", "Cargo.toml"]
---

**Must.** Before a Rust change counts as done, `cargo fmt --check`, `cargo clippy --all-targets --all-features -- -D warnings` and `cargo test` pass.
