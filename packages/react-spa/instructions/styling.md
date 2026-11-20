---
description: Styling React components with Tailwind CSS
load: on-demand
paths: ["**/*.tsx", "**/*.css", "**/vite.config.ts"]
---

**Prefer.** When the app uses Tailwind CSS, it loads through `@tailwindcss/vite` with one `src/index.css` that starts with `@import "tailwindcss";`, and layout uses utility classes, no inline `style`.
