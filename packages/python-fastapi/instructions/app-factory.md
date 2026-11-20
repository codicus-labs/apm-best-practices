---
description: The FastAPI app is built by a create_app factory
load: on-demand
paths: ["**/main.py", "**/app.py"]
---

**Prefer.** `src/<package>/main.py` defines `create_app() -> FastAPI`, which registers routers, exception handlers and the `lifespan` context manager (for pools and clients), and ends with `app = create_app()`. No `@app.on_event`, no work at import time besides building the app.
