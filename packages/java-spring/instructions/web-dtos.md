---
description: Request and response DTOs of Spring controllers
load: on-demand
paths: ["**/*Controller.java", "**/web/**/*.java"]
---

**Must.** Controllers accept and return Java records from the `web` package, never entities. Request bodies are validated with `@Valid` and Jakarta constraints on the record's components.
