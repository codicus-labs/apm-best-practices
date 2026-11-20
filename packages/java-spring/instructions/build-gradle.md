---
description: Gradle wrapper and version catalog in Spring Boot builds
load: on-demand
paths: ["**/build.gradle.kts", "**/build.gradle", "**/settings.gradle.kts", "**/gradle/libs.versions.toml", "**/gradle/wrapper/**"]
---

**Prefer.** When the build uses Gradle, it runs through the committed wrapper (`./gradlew`, with `gradle/wrapper/`), and dependency and plugin versions live in `gradle/libs.versions.toml`, referenced from build scripts as `libs.…`. The Gradle toolchain names the JDK version the project runs on.
