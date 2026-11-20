---
description: "Database migrations are fixed forward: an applied migration is never edited, schema changes, SQL migrations"
load: on-demand
paths: ["**/migrations/**", "**/db/migrations/**", "**/drizzle/**", "**/alembic/versions/**", "**/db/migration/**"]
---

**Must.** A migration that has been applied anywhere outside a developer's machine is never edited, renamed or deleted; a fix is a new migration. Databases that ran the old file would otherwise differ from the ones that run the edited one, and migration tools that checksum applied files refuse to start.
