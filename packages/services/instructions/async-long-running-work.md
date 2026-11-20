---
description: "Long-running work runs asynchronously: background jobs, workers, queues, work that outlasts a request"
load: on-demand
paths: ["**/jobs/**", "**/*job*.*", "**/workers/**", "**/*worker*.*", "**/queue/**", "**/*queue*.*"]
---

**Prefer.** Work that takes longer than a request runs asynchronously: the request records a job and answers at once with its id, a worker runs the job, and the job's status (pending, running, succeeded, failed) can be read by the caller. Jobs are idempotent, so a worker retries a failed job with backoff up to a fixed number of attempts and then marks it failed with the error.
