---
description: tflint configuration for Terraform
load: on-demand
paths: ["**/.tflint.hcl"]
---

**Must.** When the repository uses tflint, `.tflint.hcl` at the root enables the `terraform` plugin with `preset = "recommended"` and the ruleset plugin of every cloud provider the repository uses, each with a pinned `version`.
