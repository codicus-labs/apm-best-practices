---
description: Minimal production images
load: on-demand
paths: ["**/Dockerfile", "**/*.Dockerfile", "**/.dockerignore"]
---

**Must.** Images that go to production contain only what the program needs at runtime: multi-stage builds, a small base image, no build tools, no dev dependencies. Fewer parts mean fewer vulnerabilities, less memory and smaller images.
