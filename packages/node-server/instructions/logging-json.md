---
description: "Logging in Node.js servers: structured JSON lines to stdout through one logger, no console.log"
load: on-demand
paths: ["**/logger.ts", "**/server/**/*.ts", "**/src/main.ts"]
---

**Must.** Server code logs structured JSON, one line per event, to stdout, through one logger module, never through `console`: log collectors parse the lines, and a stray `console.log` breaks that.

Exception: a command-line script prints its output for a person to stdout as it likes.
