---
description: Cargo workspace layout for repositories with several crates
load: on-demand
paths: ["Cargo.toml", "**/Cargo.toml"]
---

**Prefer.** A repository with more than one crate is a Cargo workspace: the root `Cargo.toml` has `[workspace]` with `members = ["crates/*"]` and `resolver = "3"`, shared versions sit in `[workspace.dependencies]`, and member crates write `serde = { workspace = true }`.
