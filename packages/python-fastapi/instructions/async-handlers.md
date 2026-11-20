---
description: Non-blocking async route handlers in FastAPI (no blocking I/O on the event loop)
load: on-demand
paths: ["**/api/**/*.py", "**/clients/**/*.py"]
---

**Must.** Route handlers are `async def` and never call blocking I/O on the event loop: outbound calls go through an async client shared for the app's lifetime, waiting uses `asyncio.sleep`, and a blocking library that has no async client runs in `asyncio.to_thread`.
