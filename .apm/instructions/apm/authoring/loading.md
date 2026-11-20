---
description: How an instruction chooses load, paths and description so that it reaches agents exactly when it matters, also under a scope
load: on-demand
paths: ["**/instructions/**/*.md"]
---

**Must.**

- `load: always` is for what every session needs, short (under 150 words): it is in every subscribing repository's `AGENTS.md`. Everything else is `load: on-demand`. Under a scope, an always-loaded instruction becomes on-demand for every file of the scope, so prefer on-demand with paths for anything about certain files.
- `paths` names the files where the concern lives: the file kinds it is about (`**/*.vue`, `**/Dockerfile`) or the places a topic sits (`**/storage/**`, `**/*cache*.*`, `compose.yaml`). A topic that can sit in any source file (asynchronous work, observability, architecture) gets no paths rather than every `**/*.ts`: each matching file injects the instruction, and fifteen of those on one file bury the one that matters.
- Paths are relative to the root of a subscribing repository, or of its scope directory when the package is subscribed with a scope: `**/x` matches at any depth under either, `x` only at that root.
- `description` is one line that says what the instruction is about and when it matters, in the words an agent would search for, including the common other names of the thing (`background jobs, queues, workers`). It is all an agent sees of an instruction without paths, through `apm context --search`.
- An on-demand instruction stays under 400 words: it is injected whenever a matching file is touched.

Exception: an instruction that applies to every file of the repository is `load: always`, never `paths: ["**"]`.

Enforced by `apm package check`, which refuses paths for every file and warns about long instructions, missing strengths and descriptions too short to find.
