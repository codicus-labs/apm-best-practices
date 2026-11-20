---
description: "Actions pinned by commit SHA in GitHub Actions workflows, with the version as a comment"
load: on-demand
paths: [".github/workflows/*.yml", ".github/workflows/*.yaml"]
---

**Prefer.** Every `uses:` of an action outside the repository names the full commit SHA, followed by the version it stands for as a comment: `uses: actions/checkout@<40-character SHA> # vX.Y.Z`. A tag can be moved to other code; a SHA cannot, and the comment lets update tools and readers see the version.
