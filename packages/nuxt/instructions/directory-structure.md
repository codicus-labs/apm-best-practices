---
description: Directory structure of a Nuxt app
load: on-demand
paths: ["**/nuxt.config.ts", "**/app/**", "**/server/**", "**/shared/**"]
---

**Prefer.** A Nuxt app follows Nuxt's current directory structure: client code in `app/` (`app/pages/`, `app/components/`, `app/composables/`, `app/layouts/`), server code in `server/`, code and types both use in `shared/`. Nothing in `app/` imports from `server/`.
