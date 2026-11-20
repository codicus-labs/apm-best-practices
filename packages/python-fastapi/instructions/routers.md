---
description: One APIRouter per resource in FastAPI
load: on-demand
paths: ["**/api/**/*.py", "**/routers/**/*.py"]
---

**Prefer.** Routes live in `src/<package>/api/<resource>.py`, each module with `router = APIRouter(prefix="/<resource>", tags=["<resource>"])`, included by `create_app()`. A route function only translates HTTP to a call into the domain and back.
