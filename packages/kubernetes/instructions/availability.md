---
description: Replicas and PodDisruptionBudgets for production Kubernetes workloads
load: on-demand
paths: ["**/k8s/**/*.yaml", "**/kubernetes/**/*.yaml"]
---

**Prefer.** A production `Deployment` runs at least two replicas spread with `topologySpreadConstraints` over `kubernetes.io/hostname`, and has a `PodDisruptionBudget` with `maxUnavailable: 1`, so node maintenance never takes the whole service down.
