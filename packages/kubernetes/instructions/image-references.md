---
description: Immutable image references in Kubernetes manifests
load: on-demand
paths: ["**/k8s/**/*.yaml", "**/kubernetes/**/*.yaml", "**/kustomization.yaml"]
---

**Must.** Images are referenced by an immutable version tag or a digest, never `:latest` or a branch name. When manifests use Kustomize, the overlay sets them through the `images:` field of `kustomization.yaml`, which CI updates on release.
