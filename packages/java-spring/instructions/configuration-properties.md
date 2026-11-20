---
description: Typed configuration with @ConfigurationProperties records
load: on-demand
paths: ["**/*.java", "**/src/main/resources/application*.yml"]
---

**Must.** Application settings are bound to `@ConfigurationProperties("<prefix>")` records annotated with `@Validated` and Jakarta constraints, enabled with `@ConfigurationPropertiesScan`. No `@Value` and no `Environment.getProperty` in application code.
