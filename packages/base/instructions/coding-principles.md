---
description: Coding principles that keep code easy to change, from an obvious main path to pure logic
load: always
---

Optimize for the next change: leave the code easier to understand and modify than before.

1. **Keep the main path obvious.** Prefer early returns and small functions over nested control flow.
    ```ts
    if (!order) return notFound();
    if (!order.canShip) return rejected('order_not_shippable');
    return ship(order);
    ```
2. **Name by meaning.** Names must reveal domain intent and role: `pendingOrder`, not `data`; `isPaymentOverdue`, not `check`.
3. **Keep external systems behind a boundary.** Isolate third-party APIs, storage, clocks, and frameworks behind adapters owned by the application. Translate their data and errors at the boundary; do not spread their models or quirks through the codebase.
    ```ts
    const payment = await paymentGateway.charge(request);
    // Not: stripe.paymentIntents.create(...) throughout the application.
    ```
4. **Make invalid states hard to represent.** Prefer explicit variants and required fields over nullable, partially valid objects.
    ```ts
    type Order = { status: 'pending'; items: Item[] } | { status: 'paid'; items: Item[]; paymentId: PaymentId };
    ```
5. **Separate decisions from actions.** Keep policy pure and independently testable; perform side effects only after the decision.
    ```ts
    const decision = canUseFeature(account, feature);
    if (!decision.allowed) return decision;
    return enableFeature(account, feature);
    ```
6. **Make errors actionable.** Human messages explain what failed and what to do. Stable codes let systems react. Include relevant context, never secrets or sensitive data.
    ```ts
    throw new AppError('payment_declined', 'Payment was declined', { orderId });
    ```
7. **Keep changes focused.** Every change has one clear purpose. Avoid unrelated refactors, formatting, renames, dependencies, or abstractions.

8. **Write pure logic.** Logic is a pure function of its inputs; state, I/O and other side effects live at the edges. Extract a function when it gives a piece of logic a name, one responsibility or a test of its own. Do not extract for its own sake: a function that only wraps one line, or is called once and hides no decision, adds indirection, not clarity.
9. **One responsibility per unit, concerns kept apart.** Every module, class and function has one reason to change. Follow the language's standard conventions and the framework's structure before inventing your own.
