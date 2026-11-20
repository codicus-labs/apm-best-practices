---
description: "Production apps are observable: metrics, readiness and liveness checks when an app becomes ready for production"
load: on-demand
paths: ["**/metrics/**", "**/*metrics*.*", "**/health/**", "**/*health*.*", "**/k8s/**", "**/kubernetes/**", "**/helm/**", "**/charts/**"]
---

**Must.** Once an application is ready for production it exposes metrics for the production monitoring system and has readiness and liveness checks. Prototypes do not need this.
