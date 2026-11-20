---
description: Pinned Terraform and provider versions
load: on-demand
paths: ["**/versions.tf", "**/.terraform.lock.hcl"]
---

**Must.** `versions.tf` sets `required_version` and lists every provider in `required_providers` with `source` and a `~>` constraint. The `.terraform.lock.hcl` of each root module is committed, so every run uses the same provider builds.
