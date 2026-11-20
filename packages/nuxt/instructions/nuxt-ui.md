---
description: Using Nuxt UI components and its theme colors in app.config.ts
load: on-demand
paths: ["**/*.vue", "**/app.config.ts", "**/assets/css/**"]
---

**Prefer.** When the app uses Nuxt UI, interfaces are built from its components (`UButton`, `UForm`, `UTable`, `UModal`) before writing custom ones, and colors are set once in `app.config.ts` under `ui.colors`, not per component.
