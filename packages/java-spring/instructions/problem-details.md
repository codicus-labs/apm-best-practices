---
description: Errors as RFC 9457 problem details in Spring
load: on-demand
paths: ["**/*Advice.java", "**/*ExceptionHandler.java", "**/web/**/*.java", "**/src/main/resources/application*.yml"]
---

**Prefer.** Error responses are `ProblemDetail` (RFC 9457): `spring.mvc.problemdetails.enabled: true`, and one `@RestControllerAdvice` maps domain exceptions to a status and a stable `type` URI. Stack traces never reach the response.
