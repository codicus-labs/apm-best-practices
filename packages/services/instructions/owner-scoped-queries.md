---
description: "Data access is scoped to its owner: every query filters by user, tenant or organisation id, multi-tenancy, authorization, IDOR"
load: on-demand
---

**Must.** Every read or write of data that belongs to a user or tenant filters by that owner in the same query (`WHERE id = $1 AND owner_id = $2`), with the owner taken from the authenticated session, never from the request body. Checking ownership in a separate step leaves a gap between check and use, and a forgotten check exposes other owners' data.

Exception: background jobs working on records whose ids they claimed from a queue need no owner filter.
