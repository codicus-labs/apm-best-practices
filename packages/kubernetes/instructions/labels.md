---
description: Recommended labels on Kubernetes objects
load: on-demand
paths: ["**/kustomization.yaml", "**/k8s/**/*.yaml", "**/kubernetes/**/*.yaml"]
---

**Must.** Every object carries the labels `app.kubernetes.io/name` and `app.kubernetes.io/part-of`. When manifests use Kustomize, they are set once through `labels:` with `includeSelectors: false` in the base's `kustomization.yaml`, plus `app.kubernetes.io/managed-by: kustomize`; selectors name `app.kubernetes.io/name` in the manifests themselves. Selectors of Deployments are immutable, so a label added to them later forces the workload to be recreated.
