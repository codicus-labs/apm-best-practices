---
description: Containers run as a non-root user
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile"]
---

**Must.** The final stage switches to a non-root user with a numeric UID and GID (`USER 65532:65532`), so Kubernetes can verify `runAsNonRoot`. Files the program only reads are owned by root.
