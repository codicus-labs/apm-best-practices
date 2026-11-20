---
description: Names of Terraform resources and variables
load: on-demand
paths: ["**/*.tf"]
---

**Prefer.** Resource, data, variable and output names are `snake_case` and do not repeat the type: `aws_s3_bucket.assets`, not `aws_s3_bucket.assets_bucket`. A module's only resource of a type is named `this`.
