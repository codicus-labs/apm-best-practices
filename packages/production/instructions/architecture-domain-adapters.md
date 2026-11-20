---
description: "Domain logic is separate from adapters: where business logic, persistence and integrations go when adding a feature, module or external system"
load: on-demand
paths: ["**/domain/**", "**/adapters/**", "**/ports/**"]
---

**Prefer.** Business logic lives in a domain layer that knows no framework, database, transport or third-party API; adapters translate between the domain and those systems (clean architecture, ports and adapters). This keeps adapters swappable and the domain testable without them. Where a framework brings its own structure, follow it and keep the domain inside it as plain modules that import nothing from the framework.
