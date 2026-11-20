---
description: Validating constructors and unexported fields keep Go values valid
load: on-demand
paths: ["**/*.go"]
---

**Prefer.** Go has no sum types and cannot enforce required fields, so invalid states (principle 4 of the coding principles) are only partly prevented. A type whose fields must hold together keeps them unexported and is built through a validating constructor (`func NewOrder(items []Item) (Order, error)`).
