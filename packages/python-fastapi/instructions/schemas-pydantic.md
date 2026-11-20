---
description: Pydantic models for FastAPI request and response bodies
load: on-demand
paths: ["**/api/**/*.py", "**/schemas.py", "**/schemas/**/*.py"]
---

**Must.** Request and response bodies are Pydantic models in `schemas.py` (or `schemas/`), and every route declares its `response_model` or return type. Routes never return ORM objects or bare dicts, so the OpenAPI description matches the responses.
