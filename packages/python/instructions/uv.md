---
description: Managing Python dependencies and environments with uv (uv.lock, uv add, uv run, uv sync)
load: on-demand
paths: ["**/pyproject.toml", "**/uv.lock", "**/requirements*.txt", "**/Dockerfile", ".github/workflows/**", ".gitlab-ci.yml"]
---

**Must.** When a Python project uses uv, dependencies are declared in `pyproject.toml` and locked in a committed `uv.lock`, added with `uv add` (`uv add --dev` for tools). Commands run through `uv run`; CI and images install with `uv sync --frozen`, so a stale lockfile fails rather than resolves anew. There is no `requirements.txt` and no `pip install` beside it.
