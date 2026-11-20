---
description: Testing Nuxt components with @nuxt/test-utils
load: on-demand
paths: ["**/*.spec.ts", "**/*.test.ts", "**/vitest.config.ts", "**/vite.config.ts"]
---

**Prefer.** Components and composables that need Nuxt's context are tested with `@nuxt/test-utils`: the `nuxt` Vitest environment and `mountSuspended`. Plain logic without Nuxt imports is tested in the default environment, which is faster.
