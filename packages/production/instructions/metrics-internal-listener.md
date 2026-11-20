---
description: "Metrics on a separate internal listener: when exposing Prometheus metrics or adding metric labels"
load: on-demand
paths: ["**/metrics/**", "**/*metrics*.*", "**/k8s/**", "**/kubernetes/**", "**/helm/**", "**/charts/**"]
---

**Prefer.** When an application exposes Prometheus metrics, it serves them on its own port, which the cluster scrapes and the ingress never routes, so the metrics need no credential and never reach the public. Metric labels never carry member or record IDs.
