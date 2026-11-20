---
description: requires-python in pyproject.toml matches the Python the project runs on
load: on-demand
paths: ["**/pyproject.toml", "**/Dockerfile"]
---

**Prefer.** `pyproject.toml` declares `requires-python` with the minor version the project is developed and tested on as its lower bound (`requires-python = ">=3.x"`), and images use that same minor version.
