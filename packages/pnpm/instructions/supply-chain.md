---
description: "pnpm supply-chain hardening: lockfile, frozen install, packageManager, allowed build scripts, minimumReleaseAge"
load: on-demand
paths: ["pnpm-workspace.yaml", "pnpm-lock.yaml", "package.json", ".github/workflows/**", ".gitlab-ci.yml"]
---

**Must.** `pnpm-lock.yaml` is committed, the root `package.json` pins pnpm in `packageManager`, and CI installs with `pnpm install --frozen-lockfile`. The workspace settings set `minimumReleaseAge` and allow the build scripts of dependencies by name (`allowBuilds` or `onlyBuiltDependencies`), none by default: a fresh release and an install script are where a compromised package strikes.

The settings live in `pnpm-workspace.yaml` or come from a config dependency listed under `configDependencies` there.
