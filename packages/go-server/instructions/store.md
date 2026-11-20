---
description: "Data access in a Go service: one package owns the queries, maps no-rows to its own not-found error, runs transactions"
load: on-demand
paths: ["**/sqlc.yaml", "**/sqlc.yml", "**/sqlc.json", "**/*store*.go", "**/store/**", "**/*repo*.go", "**/db/**/*.go"]
---

**Prefer.** One package of a Go service owns data access: it is the one caller of the generated or hand-written queries, and the rest of the code calls its methods. It maps the driver's no-rows error (`pgx.ErrNoRows`, `sql.ErrNoRows`) to its own `ErrNotFound`, wrapped with context, so callers never import the driver, and runs every multi-statement write in one transaction (with pgx through `pgx.BeginFunc`).
