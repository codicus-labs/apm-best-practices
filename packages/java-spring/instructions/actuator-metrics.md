---
description: Actuator health probes and metrics endpoints on a separate management port
load: on-demand
paths: ["**/src/main/resources/application*.yml", "**/src/main/resources/application*.yaml", "**/build.gradle.kts"]
---

**Must.** Spring Boot Actuator is on the classpath, its endpoints listen on their own port (`management.server.port`), so they are not reachable through the public one, the exposed endpoints are `health` and the metrics endpoint the monitoring scrapes, and `management.endpoint.health.probes.enabled: true` serves `/actuator/health/readiness` and `/liveness` for Kubernetes.
