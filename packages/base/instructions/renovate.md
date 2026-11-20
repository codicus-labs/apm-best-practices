---
description: "Renovate configuration: minimumReleaseAge, vulnerability alerts and pinned digests for GitHub Actions and Docker images"
load: on-demand
paths: ["renovate.json*", ".github/renovate.json*", ".gitlab/renovate.json*", ".renovaterc*"]
---

**Prefer.** When the repository uses Renovate, its configuration sets `minimumReleaseAge` to the release age the package manager waits for, also inside `vulnerabilityAlerts`, so Renovate never proposes a release the install would refuse. It pins digests for the `github-actions`, `dockerfile` and `docker-compose` managers (`pinDigests: true` in their `packageRules`), so actions and images are referenced immutably and updated through pull requests.
