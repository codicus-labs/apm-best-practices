---
description: Restricted security context for Kubernetes pods
load: on-demand
paths: ["**/k8s/**/*.yaml", "**/kubernetes/**/*.yaml"]
---

**Must.** Pods meet the `restricted` Pod Security Standard: `runAsNonRoot: true`, `seccompProfile.type: RuntimeDefault`, and every container with `allowPrivilegeEscalation: false`, `readOnlyRootFilesystem: true` and `capabilities.drop: ["ALL"]`. Writable paths are `emptyDir` volumes.
