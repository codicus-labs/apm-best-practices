---
description: "Testcontainers in tests: images matching production, one container per test binary, a database per test, integration tests against real backing services"
load: on-demand
paths: ["**/*_test.go", "**/*.test.ts", "**/*.spec.ts", "**/tests/**", "**/test_*.py", "**/conftest.py", "**/src/test/**", "**/testutil/**", "**/testdata/**"]
---

**Prefer.** When tests use Testcontainers, each container runs the same image as production or `compose.yaml` and is started by the test itself, never expected to be running already. One container per test binary or session serves all its tests, and each test gets its own database (or bucket, or key prefix) in it, so tests run in parallel without sharing state and stay fast.
