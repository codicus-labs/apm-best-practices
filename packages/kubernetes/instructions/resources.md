---
description: CPU and memory requests and limits of Kubernetes containers
load: on-demand
paths: ["**/k8s/**/*.yaml", "**/kubernetes/**/*.yaml"]
---

**Prefer.** Every container sets `resources.requests.cpu`, `resources.requests.memory` and `resources.limits.memory`, with the memory limit equal to the request, and sets no CPU limit: a CPU limit throttles a pod even when the node has idle CPU.
