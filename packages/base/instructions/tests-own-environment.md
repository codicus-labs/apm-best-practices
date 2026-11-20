---
description: "Tests set their own environment: git identity, temporary HOME, cache and state directories, environment variables, so they behave alike on every machine"
load: on-demand
---

**Prefer.** A test sets every part of the environment it depends on itself: a git identity through `GIT_AUTHOR_NAME`, `GIT_AUTHOR_EMAIL`, `GIT_COMMITTER_NAME` and `GIT_COMMITTER_EMAIL`, and `HOME`, cache, config and state directories pointing into a temporary directory the test creates. Tests then pass alike on a developer machine and in CI, and never read or change the developer's real configuration.
