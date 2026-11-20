---
description: Typed and described Terraform variables
load: on-demand
paths: ["**/variables.tf"]
---

**Must.** Every variable has a `type` and a `description`. Constrained values carry a `validation` block with a message, and values that are secrets are marked `sensitive = true`.

```hcl
variable "environment" {
  type        = string
  description = "Deployment environment"
  validation {
    condition     = contains(["staging", "production"], var.environment)
    error_message = "environment must be staging or production."
  }
}
```
