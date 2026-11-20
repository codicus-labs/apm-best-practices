---
description: How the body and front matter of an instruction file are written so an agent can follow it
load: on-demand
paths: ["**/instructions/**/*.md"]
---

**Must.**

- An instruction has one concern. A file that answers two questions becomes two instructions; one that repeats another is merged into it.
- Front matter is `description`, `load` and, for on-demand instructions, `paths`; how to choose them is in the instruction on how instructions load.
- The body starts with its strength, **Must.** or **Prefer.**, and states what the code does, positively and checkably: an agent or reviewer can tell from a file whether it follows the instruction. Give the reason where the instruction would otherwise look arbitrary, and a short example where words leave room.
- The instruction itself is one to three sentences; exceptions follow in a separate paragraph starting with "Exception:".
- An instruction that names a library, tool or file says what to do in a repository that lacks it: add it as part of following the instruction, or leave the repository as it is. An agent that cannot tell stops to ask, and an unattended one does not finish its task.
- Instructions name no versions: a tool an instruction names means its current release.
- An instruction that code enforces (a shared config, a lint rule, a check) points to it in a last line "Enforced by …" instead of repeating its details.
- The body holds none of APM's block markers (HTML comments starting with `apm:begin` or `apm:end`); they would break the block in `AGENTS.md`.

Enforced by `apm package check` (front matter, markers, strength, length).
