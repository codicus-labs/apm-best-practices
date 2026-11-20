---
description: Error types in Rust libraries and binaries, unwrap and expect
load: on-demand
paths: ["**/*.rs"]
---

**Must.** Library crates return typed errors (an enum implementing `std::error::Error`) that callers can match on; binaries return one error type that carries context added at each step (`reading config`). `unwrap()` and `expect()` appear in tests and for invariants the code just established, with `expect` naming the invariant.
