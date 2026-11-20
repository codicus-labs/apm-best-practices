---
description: Shape of a Docker Compose file
load: on-demand
paths: ["compose.yaml", "compose.*.yaml", "**/docker-compose*.yml", "**/docker-compose*.yaml"]
---

**Must.** The file is named `compose.yaml` (overrides `compose.<name>.yaml`), has no top-level `version:`, gives every service an explicit image tag, and every service others wait for has a `healthcheck`, used as `depends_on: { db: { condition: service_healthy } }`.
