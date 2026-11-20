---
description: Described Terraform outputs
load: on-demand
paths: ["**/outputs.tf"]
---

**Must.** Every output has a `description`, and an output that carries a secret is marked `sensitive = true`. A module outputs what its callers use, not every attribute of its resources.
