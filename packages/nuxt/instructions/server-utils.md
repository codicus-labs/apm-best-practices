---
description: Shared Nuxt server code in server/utils
load: on-demand
paths: ["**/server/utils/**", "**/server/plugins/**"]
---

**Prefer.** Code shared by server routes (the database client, auth helpers, adapters) lives in `server/utils/`, where Nitro auto-imports it. Connections are created once at module level or in a `server/plugins/` plugin, not per request.
