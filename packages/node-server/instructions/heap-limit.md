---
description: "V8 heap limit in a Node.js container: --max-old-space-size in the Dockerfile CMD"
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile"]
---

**Prefer.** The production `CMD` of a Node.js image caps the V8 heap with `--max-old-space-size=<MiB>`, sized to about two thirds of the container's memory limit to leave room for memory outside the heap, for example `CMD ["node", "--max-old-space-size=256", "src/main.ts"]` for a 384 MiB limit. Without it Node sizes its heap from the host's memory, and the container limit kills the process before it collects garbage.
