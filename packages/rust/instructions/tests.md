---
description: Where Rust unit and integration tests live, and async tests with tokio::test
load: on-demand
paths: ["**/*.rs", "**/tests/**/*.rs"]
---

**Prefer.** Unit tests sit in a `#[cfg(test)] mod tests` at the bottom of the file they test; integration tests in the crate's `tests/` directory, with shared helpers in `tests/common/`. Async tests use `#[tokio::test]`.
