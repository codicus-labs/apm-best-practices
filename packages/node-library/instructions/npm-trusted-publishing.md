---
description: Publishing npm packages with Trusted Publishing (OIDC from CI), no npm tokens
load: on-demand
paths: ["**/package.json", ".npmrc", ".github/workflows/**", ".gitlab-ci.yml"]
---

**Must.** When a package is published to npm, it is published through Trusted Publishing from a CI job, never with a token. The first setup needs the owner, because it is configured on npm.
