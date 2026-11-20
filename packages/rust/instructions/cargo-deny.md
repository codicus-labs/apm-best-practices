---
description: "cargo-deny configuration in deny.toml: advisories, license allow-list, duplicate versions"
load: on-demand
paths: ["deny.toml"]
---

**Prefer.** When a repository checks its Rust dependencies with cargo-deny, the root `deny.toml` enables `[advisories]`, lists the accepted licenses under `[licenses] allow`, and sets `[bans] multiple-versions = "warn"`. An advisory or license it ignores is listed with a comment giving the reason.
