---
description: File-based routes and validated search params with TanStack Router
load: on-demand
paths: ["**/src/routes/**", "**/routeTree.gen.ts", "**/vite.config.ts"]
---

**Prefer.** When the app uses TanStack Router, routes are files in `src/routes/` (`__root.tsx`, `orders/$orderId.tsx`) through `@tanstack/router-plugin/vite`, and every route that reads search params validates them with `validateSearch`. `routeTree.gen.ts` is generated and never edited by hand.
