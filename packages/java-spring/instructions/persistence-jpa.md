---
description: Spring Data JPA settings and repositories (open-in-view, ddl-auto validate, JpaRepository)
load: on-demand
paths: ["**/*Repository.java", "**/persistence/**/*.java", "**/src/main/resources/application*.yml"]
---

**Prefer.** When an app persists with Spring Data JPA, it sets `spring.jpa.open-in-view: false` and `spring.jpa.hibernate.ddl-auto: validate`. Repositories extend `JpaRepository`; queries beyond derived names are written with `@Query` in JPQL. Entities never leave the `persistence` package.
