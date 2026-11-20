---
description: "GitHub Actions workflow permissions: read-only GITHUB_TOKEN by default, write scopes per job, checkout without persisted credentials"
load: on-demand
paths: [".github/workflows/*.yml", ".github/workflows/*.yaml"]
---

**Must.** Every workflow sets `permissions: contents: read` at the top level, and a job that publishes (packages, images, releases) widens exactly the scopes it needs in its own `permissions`. Every `actions/checkout` step sets `persist-credentials: false`, so later steps cannot reuse the token from `.git/config`. A compromised step then holds no more rights than its job needs.
