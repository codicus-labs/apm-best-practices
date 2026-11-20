---
description: Dockerfile layer order and BuildKit cache mounts
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile"]
---

**Prefer.** A Dockerfile copies the dependency manifests and lock files first, installs dependencies, and copies the source after, so a code change does not reinstall dependencies. Package caches use BuildKit cache mounts (`RUN --mount=type=cache,target=/root/.cache/go-build go build …`).
