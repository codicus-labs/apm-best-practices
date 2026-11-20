---
description: Readiness and liveness probes of Kubernetes containers
load: on-demand
paths: ["**/k8s/**/*.yaml", "**/kubernetes/**/*.yaml"]
---

**Must.** Every long-running container has a `readinessProbe` and a `livenessProbe` on two distinct endpoints of the application: readiness may check dependencies, liveness checks just that the process responds. A container that starts slowly gets a `startupProbe` rather than a long `initialDelaySeconds`.
