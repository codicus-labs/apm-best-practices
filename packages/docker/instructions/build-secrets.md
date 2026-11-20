---
description: Secrets during docker build
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile"]
---

**Must.** A secret a build needs (a registry token) is passed with `RUN --mount=type=secret,id=<name>` and `docker build --secret`. Secrets never go through `ARG`, `ENV` or `COPY`; those end up in the image's layers or history.
