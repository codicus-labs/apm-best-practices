---
description: Where state lives in a React app with TanStack Query
load: on-demand
paths: ["**/*.tsx", "**/src/stores/**"]
---

**Prefer.** When the app uses TanStack Query, data from the API lives in its cache and is never copied into other state. Client state stays in the component (`useState`), in the URL (search params), or, when shared across distant components, in a store in `src/stores/`.
