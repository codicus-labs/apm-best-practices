---
description: Mapping domain errors to HTTP responses in FastAPI
load: on-demand
paths: ["**/main.py", "**/errors.py", "**/exceptions.py"]
---

**Prefer.** The domain raises its own exceptions (`OrderNotFound`); exception handlers registered in `create_app()` map them to responses with a stable body `{"code": "order_not_found", "message": "…"}`. `HTTPException` is raised only inside `api/`.
