---
description: Writing Python tests with pytest (fixtures, conftest.py, parametrize)
load: on-demand
paths: ["**/tests/**/*.py", "**/conftest.py", "**/pyproject.toml"]
---

**Prefer.** When a Python project tests with pytest, tests are plain functions in `tests/**/test_*.py`, with fixtures in `conftest.py` and `@pytest.mark.parametrize` for cases, and no `unittest.TestCase` classes. Configuration sits in `[tool.pytest.ini_options]` of `pyproject.toml`.
