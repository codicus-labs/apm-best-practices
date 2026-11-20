---
description: The src layout of a Python project
load: on-demand
paths: ["**/pyproject.toml", "**/src/**/*.py"]
---

**Prefer.** Importable code lives in `src/<package>/`, tests in `tests/` beside `src/`. `pyproject.toml` declares the package with a build backend (`uv_build` or `hatchling`), so tests import the installed package, not the working directory.
