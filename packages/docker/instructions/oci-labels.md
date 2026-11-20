---
description: OCI labels on built images
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile"]
---

**Prefer.** The final stage sets `org.opencontainers.image.source` to the repository URL and `org.opencontainers.image.revision` from a build argument filled with the commit SHA in CI, so every running image leads back to its code.
