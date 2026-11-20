---
description: Nuxt server API routes with validated input
load: on-demand
paths: ["**/server/api/**", "**/server/routes/**"]
---

**Must.** Server endpoints are files `server/api/<resource>.<method>.ts` (`orders.post.ts`, `orders/[id].get.ts`) exporting `defineEventHandler`. Input is validated against a schema through `readValidatedBody(event, schema.parse)` and `getValidatedQuery(event, schema.parse)`; failures are thrown with `createError({ statusCode, statusMessage })`.
