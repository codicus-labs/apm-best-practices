---
description: Alembic migrations for the FastAPI database schema (autogenerate, review, no create_all)
load: on-demand
paths: ["**/alembic/**", "**/alembic.ini"]
---

**Must.** When the schema is migrated with Alembic, it changes through revisions in `alembic/versions/`, generated with `alembic revision --autogenerate` and read and corrected before committing, since autogenerate misses renames and some type changes. `metadata.create_all()` runs in tests, never against a deployed database.
