---
description: ENTRYPOINT and CMD in exec form
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile"]
---

**Must.** `ENTRYPOINT` and `CMD` use the exec form (`ENTRYPOINT ["/app/server"]`), so the program runs as PID 1 and receives `SIGTERM` directly; the program shuts down gracefully on it. No shell form, no wrapper script that does not `exec`.
