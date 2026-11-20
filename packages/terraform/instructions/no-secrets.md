---
description: No secrets in Terraform files
load: on-demand
paths: ["**/*.tf", "**/*.tfvars"]
---

**Must.** No secret is written in a `.tf` or committed `.tfvars` file. Secrets reach Terraform through a data source that reads them at plan time or through `TF_VAR_<name>` environment variables in CI; `*.secret.tfvars` is git-ignored.
