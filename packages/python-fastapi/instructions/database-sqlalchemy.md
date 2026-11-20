---
description: Async SQLAlchemy usage in FastAPI (Mapped columns, select, AsyncSession per request)
load: on-demand
paths: ["**/models.py", "**/models/**/*.py", "**/db/**/*.py", "**/repositories/**/*.py"]
---

**Prefer.** When a FastAPI app uses SQLAlchemy, it runs in async mode: models are declarative classes with typed `Mapped[...]` columns, queries use `select()`, and one `AsyncSession` per request comes from a dependency. No legacy `Query` API.
