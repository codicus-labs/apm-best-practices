---
description: Form validation with react-hook-form and a zod schema (zodResolver)
load: on-demand
paths: ["**/*Form.tsx", "**/*form*.ts", "**/*schema*.ts"]
---

**Prefer.** When the app uses react-hook-form, a form validates with a zod schema through `zodResolver` from `@hookform/resolvers/zod`, and the form's type is `z.infer<typeof schema>`. The same schema validates on submit, with no second hand-written check.
