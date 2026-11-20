---
description: Commands a Terraform change passes before it counts as done
load: on-demand
paths: ["**/*.tf", "**/*.tfvars", "**/.terraform.lock.hcl", "**/.tflint.hcl"]
---

**Must.** Before a Terraform change counts as done, `terraform fmt -check -recursive` and `terraform validate` pass, `tflint --recursive` passes when the repository uses tflint, and the `terraform plan` of every affected environment shows nothing but the intended changes.
