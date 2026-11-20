---
description: The .dockerignore keeps the build context small and free of secrets
load: on-demand
paths: ["**/.dockerignore", "**/*.dockerignore", "**/Dockerfile", "**/*.Dockerfile"]
---

**Must.** Every build context has a `.dockerignore` at its root (or a `<Dockerfile name>.dockerignore` next to a Dockerfile that needs its own) that excludes at least `.git`, `.env*`, dependency directories (`node_modules`, `.venv`, `target`) and build output, so neither secrets nor stale artefacts reach the build context.
