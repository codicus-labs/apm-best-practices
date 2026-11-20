---
description: How React components are written
load: on-demand
paths: ["**/*.tsx"]
---

**Must.** Components are function components with a named export (`export function OrderList(props: OrderListProps)`), one per file, the file named like the component. Props are a `type <Component>Props`. No class components, no `React.FC`, no default exports except where a tool requires them.
