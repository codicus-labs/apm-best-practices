---
description: Graceful shutdown of a Go HTTP server on SIGTERM, and server timeouts (ReadHeaderTimeout)
load: on-demand
paths: ["**/cmd/**/*.go", "**/*server*.go"]
---

**Must.** A Go service stops on SIGINT and SIGTERM through a context from `signal.NotifyContext`; when it is done, it calls `http.Server.Shutdown` with a context that has a timeout, so in-flight requests finish and a stuck one cannot block the exit. Every `http.Server` sets `ReadHeaderTimeout`, which closes slow-header (Slowloris) connections.
