---
description: Pinned sources of Terraform modules
load: on-demand
paths: ["**/*.tf"]
---

**Must.** A module from outside the repository is pinned: a registry module with an exact `version`, a git module with `?ref=<tag>`. Modules of the repository use relative paths (`../../modules/network`).
