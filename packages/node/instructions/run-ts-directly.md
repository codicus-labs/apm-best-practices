---
description: "Running TypeScript directly with Node (type stripping): no build step, erasable syntax, .ts import extensions"
load: on-demand
paths: ["**/tsconfig.json", "**/package.json", "**/Dockerfile"]
---

**Prefer.** When the repository's Node version runs TypeScript natively, a service starts with `node src/main.ts`, in development (`node --watch src/main.ts`) and in its image, without a build step. Its code keeps to erasable syntax (no enums, namespaces or parameter properties) and relative imports carry the `.ts` extension; `tsconfig.json` sets `erasableSyntaxOnly` and `allowImportingTsExtensions` so the type check catches both.

Exception: a package published to npm is compiled to JavaScript for its consumers.
