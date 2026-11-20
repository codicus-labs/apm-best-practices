---
description: "Database migrations with goose in a Go service: embedded SQL files applied on startup"
load: on-demand
paths: ["**/migrations/*.sql", "**/db/**/*.go"]
---

**Prefer.** When a Go service migrates its schema with goose, the SQL migration files are embedded in the binary with `//go:embed` and applied with `goose.NewProvider(…).Up(ctx)` when the database is opened, so every binary and test brings its database up to date by itself and no separate migration step runs before a deploy.
