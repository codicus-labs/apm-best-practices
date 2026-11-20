---
description: Commands a Kubernetes manifest change passes before it counts as done
load: on-demand
paths: ["**/k8s/**", "**/kubernetes/**", "**/kustomization.yaml"]
---

**Must.** When manifests use Kustomize, a change to them counts as done once `kubectl kustomize` renders every overlay without errors, in CI as well as locally.
