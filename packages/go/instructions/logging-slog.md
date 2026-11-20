---
description: Structured logging with log/slog in Go services (JSON handler, key-value attributes)
load: on-demand
paths: ["**/*.go"]
---

**Prefer.** When a Go service logs with `log/slog`, production logs go through `slog.NewJSONHandler` to stdout, and a message carries its data as key-value attributes (`logger.Info("order shipped", "order_id", id)`) rather than formatted into the text.
