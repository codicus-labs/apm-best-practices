---
description: Commands a Spring Boot change passes before it counts as done
load: on-demand
paths: ["**/*.java", "**/*.gradle.kts", "**/build.gradle", "**/gradle/libs.versions.toml", "**/src/main/resources/**"]
---

**Must.** When the build uses Gradle, a Java change counts as done once `./gradlew check` passes, and `check` depends on every formatting, static-analysis and test task the build defines, so that one command is the whole gate.
