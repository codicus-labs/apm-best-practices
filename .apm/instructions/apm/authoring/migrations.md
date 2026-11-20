---
description: How a migration of a package release is written
load: on-demand
paths: ["**/migrations/*.md"]
---

**Must.**

- A migration is `migrations/<version>.md`, named after the exact minor or major version that needs it (`1.2.0.md`); a patch never brings one.
- Front matter `description` says in one line what changes; `verify` is a shell command that succeeds once a repository is migrated, when such a command exists.
- The body is written for an agent in another repository that cannot ask back: what changed, how to find the code it affects, and how to change it, with an example.
- A migration stays in the package after its release; later releases keep every earlier migration unchanged.
