---
description: Wrapping and comparing errors in Go
load: on-demand
paths: ["**/*.go"]
---

**Must.** An error passed up gets context with `fmt.Errorf("load order %s: %w", id, err)`; it is compared with `errors.Is` or `errors.As`, never by its message. Sentinel errors are package-level `var ErrNotFound = errors.New("…")`.
