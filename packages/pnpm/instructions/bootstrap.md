---
description: Installing pnpm itself in a Docker image or CI job without Corepack
load: on-demand
paths: ["**/Dockerfile", ".github/workflows/**", ".gitlab-ci.yml"]
---

**Prefer.** Where no pnpm is available, such as a Node image without Corepack, `npm install --global pnpm@<version>` installs the version the root `package.json` pins in `packageManager`, and every step after that uses pnpm. A hand-rolled download with checksums is more code to maintain than the risk it removes.
