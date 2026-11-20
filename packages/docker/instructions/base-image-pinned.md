---
description: Base images pinned by tag and digest
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile"]
---

**Must.** Every `FROM` names a version tag and its digest (`FROM golang:<version>-alpine@sha256:…`), so a build is reproducible; the repository's dependency update tool keeps both current. No `latest`, no tag without digest.
