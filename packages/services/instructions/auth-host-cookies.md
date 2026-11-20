---
description: Sign-in cookies carry the `__Host-` prefix
load: on-demand
paths: ["**/auth/**", "**/*auth*.*", "**/*session*.*", "**/*cookie*.*"]
---

**Must.** Every cookie of the sign-in flow (session and pending login) is set with the `__Host-` prefix on https: Secure, `Path=/`, no `Domain`. Only that prefix stops a sibling host from planting a cookie (login CSRF); signing the cookie does not, because an attacker can obtain a validly signed one from their own sign-in.

Exception: plain-http local development uses the same name without the prefix.
