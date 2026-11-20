---
description: Kustomize base and overlays for Kubernetes manifests
load: on-demand
paths: ["**/k8s/**", "**/kubernetes/**", "**/kustomization.yaml"]
---

**Prefer.** When an application's manifests use Kustomize, they live in `k8s/base/` with a `kustomization.yaml`, and each environment is an overlay in `k8s/overlays/<env>/` that patches just what differs from the base.
