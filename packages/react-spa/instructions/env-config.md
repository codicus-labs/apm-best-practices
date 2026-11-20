---
description: Build-time configuration of a Vite React app
load: on-demand
paths: ["**/src/config.ts", "**/.env", "**/.env.*", "**/vite-env.d.ts"]
---

**Must.** `src/config.ts` is the one module that reads `import.meta.env`; it validates the `VITE_*` variables against a schema and exports a typed `config`. `VITE_*` variables end up in the bundle every visitor downloads, so they never hold a secret.
