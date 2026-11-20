---
description: Strict mypy and complete type annotations in Python
load: on-demand
paths: ["**/*.py", "**/pyproject.toml"]
---

**Must.** When a Python project type-checks with mypy, `pyproject.toml` sets `[tool.mypy] strict = true`, and every function and method has annotations for its parameters and return type. `Any` and `# type: ignore` appear with a comment giving the reason (`# type: ignore[import-untyped]  # no stubs for legacy-sdk`).
