---
description: Code of a React app is grouped by feature
load: on-demand
paths: ["**/src/features/**", "**/src/components/**"]
---

**Prefer.** Code is grouped by feature in `src/features/<feature>/` (its components, hooks, API calls and types); a feature imports from another through that feature's `index.ts`, never from its inner files. Components shared by several features live in `src/components/`, without feature logic.
