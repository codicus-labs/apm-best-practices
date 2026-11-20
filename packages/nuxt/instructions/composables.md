---
description: Nuxt composables and state shared during server rendering
load: on-demand
paths: ["**/composables/**", "**/*.vue"]
---

**Must.** Composables live in `app/composables/use<Name>.ts`. State shared between components uses `useState('<key>', () => initial)` or a Pinia store, never a `ref` at module level: on the server a module-level ref is shared by every request and leaks data between users.
