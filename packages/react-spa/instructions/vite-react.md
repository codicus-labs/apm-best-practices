---
description: React single-page apps are built with Vite
load: on-demand
paths: ["**/vite.config.ts", "**/package.json", "**/index.html", "**/src/main.tsx"]
---

**Prefer.** A React single-page app is a Vite project with `@vitejs/plugin-react`: `index.html` at the app's root loads `src/main.tsx`, which renders `<App />` into `#root` inside `<StrictMode>`.
