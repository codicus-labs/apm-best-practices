---
description: Default-deny network policies in Kubernetes namespaces
load: on-demand
paths: ["**/k8s/**/*.yaml", "**/kubernetes/**/*.yaml"]
---

**Prefer.** Each application namespace has a `NetworkPolicy` that denies all ingress by default, plus policies that allow exactly the traffic the application needs: from the ingress controller to the public port, from the metrics scraper to the metrics port.
