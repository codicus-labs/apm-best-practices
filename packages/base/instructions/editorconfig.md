---
description: EditorConfig at the repository root for charset, line endings, final newline and indentation
load: on-demand
paths: [".editorconfig"]
---

**Prefer.** The repository root holds an `.editorconfig` with `root = true` and, under `[*]`, `charset = utf-8`, `end_of_line = lf` and `insert_final_newline = true`, plus `indent_style` and `indent_size` per language as its formatter writes them (for example `[*.go]` with `indent_style = tab`). Editors then write files the way the formatter and the diff expect. A repository without one adds it.
