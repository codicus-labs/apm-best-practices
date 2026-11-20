---
description: Common tags or labels on Terraform resources
load: on-demand
paths: ["**/*.tf"]
---

**Prefer.** Every resource that supports tags or labels carries the same default set (at least the environment and the managing repository), defined once per root module and applied through the provider's `default_tags` (or its equivalent) where it has them, so no resource repeats or forgets it.
