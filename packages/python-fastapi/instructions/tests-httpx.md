---
description: Testing FastAPI routes with httpx and dependency overrides
load: on-demand
paths: ["**/tests/**/*.py", "**/conftest.py"]
---

**Prefer.** API tests call `create_app()` through `httpx.AsyncClient(transport=ASGITransport(app=app), base_url="http://test")` with `pytest-asyncio` or `anyio`, and replace dependencies with `app.dependency_overrides`. No `TestClient` in async test suites.
