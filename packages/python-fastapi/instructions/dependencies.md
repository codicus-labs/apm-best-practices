---
description: Dependency injection with Depends and Annotated in FastAPI
load: on-demand
paths: ["**/api/**/*.py", "**/dependencies.py", "**/deps.py"]
---

**Prefer.** Sessions, the current user and settings reach routes through `Depends`, declared once as `Annotated` aliases in `dependencies.py`:

```python
SessionDep = Annotated[AsyncSession, Depends(get_session)]
CurrentUser = Annotated[User, Depends(get_current_user)]
```
