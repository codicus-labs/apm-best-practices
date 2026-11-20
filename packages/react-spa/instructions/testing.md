---
description: React component tests with Testing Library, queries by role, API mocked at the network
load: on-demand
paths: ["**/*.test.tsx", "**/src/mocks/**"]
---

**Prefer.** When the app uses Testing Library, component tests find elements by role or label (`getByRole('button', { name: 'Save' })`), act through `@testing-library/user-event`, and mock the API at the network, for example with request handlers in `src/mocks/handlers.ts`, never by mocking hooks.
