---
description: Environment variables for Docker Compose
load: on-demand
paths: ["compose.yaml", "compose.*.yaml", ".env.example", ".gitignore"]
---

**Prefer.** Values that differ per developer or are secret come from a git-ignored `.env` next to `compose.yaml`, referenced as `${VAR}`; a committed `.env.example` lists every variable with a working local default.
