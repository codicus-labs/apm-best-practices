---
description: "Linting Dockerfiles with hadolint: CI, ignored rules with a reason"
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile", ".hadolint.yaml"]
---

**Must.** When the repository lints Dockerfiles with hadolint, every Dockerfile passes it in CI, and a rule is ignored solely in `.hadolint.yaml` or with an inline `# hadolint ignore=DL3018` next to a comment that gives the reason.
