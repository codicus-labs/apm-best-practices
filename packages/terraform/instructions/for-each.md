---
description: for_each over count for repeated Terraform resources
load: on-demand
paths: ["**/*.tf"]
---

**Prefer.** Several instances of a resource or module are created with `for_each` over a map or set with stable keys; `count` is only a switch (`count = var.enabled ? 1 : 0`). With `count`, removing one item re-indexes and replaces the ones after it.
