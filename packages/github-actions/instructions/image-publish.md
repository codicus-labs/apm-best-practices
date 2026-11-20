---
description: "Publishing container images from GitHub Actions: after all checks pass, on the default branch and version tags, image tags that trace to branch, commit and version"
load: on-demand
paths: [".github/workflows/*.yml", ".github/workflows/*.yaml"]
---

**Prefer.** A workflow pushes a container image from a job that `needs:` every check job, on pushes to the default branch and on version tags (`v*`), never from pull requests. Each image is tagged with the branch name, the short commit SHA (`sha-<short>`) and, on version tags, the semantic version, so every image traces back to a commit.
