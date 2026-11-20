---
description: Packages of a Spring Boot app are organised by feature
load: on-demand
paths: ["**/src/main/java/**/*.java"]
---

**Prefer.** Packages are cut by feature (`com.<company>.<app>.order`), each with `domain` (plain Java, no Spring imports), `web` (controllers, DTOs) and `persistence` (repositories, entities) subpackages. Controllers call services, services call repositories; nothing calls upward.
