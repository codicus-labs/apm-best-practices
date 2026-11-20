---
description: Nuxt configuration from the environment through runtimeConfig
load: on-demand
paths: ["**/nuxt.config.ts", "**/.env.example", "**/server/**/*.ts", "**/*.vue"]
---

**Must.** Configuration is declared in `runtimeConfig` of `nuxt.config.ts`, secrets at its top level (server only), values the browser needs under `runtimeConfig.public`. At runtime they come from `NUXT_*` and `NUXT_PUBLIC_*` variables; code reads them with `useRuntimeConfig()`, never `process.env`.
