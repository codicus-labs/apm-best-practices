---
description: TanStack Query hooks and query keys in React
load: on-demand
paths: ["**/src/features/**/api.ts", "**/src/features/**/queries.ts", "**/src/**/use*.ts"]
---

**Prefer.** When the app uses TanStack Query, every API call is wrapped in a hook in its feature's `api.ts` (`useOrders`, `useCreateOrder`) built on `useQuery` or `useMutation`; components never call `fetch` directly. Query keys come from one factory per feature:

```ts
export const orderKeys = { all: ['orders'] as const, detail: (id: string) => ['orders', id] as const };
```
