---
description: Clippy and rustc lint levels set once for the Cargo workspace
load: on-demand
paths: ["Cargo.toml", "**/Cargo.toml", "clippy.toml"]
---

**Prefer.** Lint levels are set once in the workspace's `Cargo.toml` under `[workspace.lints.clippy]` and `[workspace.lints.rust]`, and every member crate opts in with `[lints] workspace = true`; no crate sets levels with `#![deny(...)]` or `#![warn(...)]` at its root. An `#[allow(...)]` names its reason: `#[allow(clippy::cast_possible_truncation, reason = "ids fit in u32")]`.
