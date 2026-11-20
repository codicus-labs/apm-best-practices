---
description: Flyway migration files for the Spring Boot database schema
load: on-demand
paths: ["**/db/migration/**", "**/build.gradle.kts"]
---

**Must.** When an app migrates its schema with Flyway, the schema changes through migrations `src/main/resources/db/migration/V<yyyyMMddHHmm>__<description>.sql`; a timestamp as the version keeps migrations from parallel branches from colliding.
