---
description: Spring Boot integration tests with Testcontainers and @ServiceConnection
load: on-demand
paths: ["**/src/test/java/**/*.java", "**/build.gradle.kts"]
---

**Prefer.** When Spring Boot integration tests use Testcontainers, they run with `@SpringBootTest`, and each container is declared as a bean or static field annotated with `@ServiceConnection`, so Spring Boot takes the connection details from it; no `@DynamicPropertySource` for what a service connection provides.
