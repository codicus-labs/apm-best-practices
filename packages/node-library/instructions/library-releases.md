---
description: "Releasing a library: release pull request, changelog, publishing from CI"
load: on-demand
paths: [".changeset/**", "**/CHANGELOG.md", ".github/workflows/**", ".gitlab-ci.yml"]
---

**Prefer.** A library is released by merging a release pull or merge request that a CI job opens from the pending changes: it bumps the version and the changelog, and its merge publishes. Nobody publishes from a laptop, so every release is reproducible from the default branch.
