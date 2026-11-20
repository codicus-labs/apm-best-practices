---
description: Layout of a Terraform repository with modules and environments
load: on-demand
paths: ["**/*.tf"]
---

**Prefer.** Reusable code lives in modules under `modules/<name>/`, and each environment is its own root module that calls them. Every module splits into `main.tf`, `variables.tf`, `outputs.tf` and `versions.tf`; a root module adds `backend.tf`. Resources are not declared in a root module when a module of the repository covers them.
