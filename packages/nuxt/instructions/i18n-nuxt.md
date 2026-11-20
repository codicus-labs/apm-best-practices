---
description: Translations in Nuxt with @nuxtjs/i18n
load: on-demand
paths: ["**/nuxt.config.ts", "**/i18n/**", "**/*.vue"]
---

**Prefer.** When the app uses the `@nuxtjs/i18n` module, it keeps one file per language in `i18n/locales/<code>.json`. Components use `const { t } = useI18n()` and links `useLocalePath()`; visible text is never a literal in a template.
