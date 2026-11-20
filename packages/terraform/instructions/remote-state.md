---
description: Remote Terraform state with locking
load: on-demand
paths: ["**/backend.tf", "**/*.tfstate", "**/*.tfstate.backup", ".gitignore"]
---

**Must.** Each root module stores its state in a remote backend with state locking, one state per environment. State files are never committed; `.gitignore` lists `*.tfstate*` and `.terraform/`. State holds secrets in plaintext, and locking stops two runs from corrupting it.
