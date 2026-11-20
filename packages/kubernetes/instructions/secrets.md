---
description: "Kubernetes Secrets in git: never plaintext, SOPS rules that encrypt data and stringData"
load: on-demand
paths: ["**/k8s/**", "**/kubernetes/**", ".sops.yaml"]
---

**Must.** A `Secret` with values is never committed in plaintext. When the repository encrypts secrets with SOPS, `.sops.yaml` holds a creation rule whose `encrypted_regex` covers `data` and `stringData` (`^(data|stringData)$`), so keys and metadata stay readable in diffs while every value is encrypted.
