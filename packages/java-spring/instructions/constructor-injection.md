---
description: Constructor injection in Spring beans
load: on-demand
paths: ["**/*.java"]
---

**Must.** Beans receive their dependencies through a single constructor into `private final` fields. No `@Autowired` on fields or setters; a class with too many constructor parameters gets split, not field-injected.
