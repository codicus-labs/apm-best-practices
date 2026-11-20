---
description: Configuration of FastAPI apps from the environment, validated once at startup (settings, env vars)
load: on-demand
paths: ["**/settings.py", "**/config.py", "**/.env.example"]
---

**Must.** Configuration comes from the environment into one typed settings object, validated once at startup so a missing or malformed variable stops the app before it serves, and passed to routes with `Depends`. Nothing else reads `os.environ`; `.env.example` lists every variable.
