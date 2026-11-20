# apm-best-practices

Best practices for coding agents as [APM](https://github.com/codicus-labs/agent-package-manager) packages: how to use a language, framework or tool well once a repository uses it. They never choose a tool for you and hold nothing specific to one organisation; an organisation layers its own choices on top in a source of its own that depends on these packages ([codicus-labs/apm-packages](https://github.com/codicus-labs/apm-packages) is one). Releases are git tags `<package>@<version>`; the format is APM's `docs/format.md`, terms are defined in its `GLOSSARY.md`.

## Using them

```jsonc
// apm.jsonc of a consuming repository
{
    "sources": { "bp": "https://github.com/codicus-labs/apm-best-practices.git" },
    "packages": {
        "bp/base": "^1.0.0",
        "bp/go-server": { "version": "^1.0.0", "scope": "services/api" },
        "bp/nuxt": { "version": "^1.0.0", "scope": "apps/web" },
        "bp/docker": "^1.0.0",
    },
}
```

Subscribe to `base` without a scope, to the package of each kind of project (`go-server`, `go-cli`, `node-server`, `nuxt` …) with its directory as scope, and to the tool packages a repository uses (`docker`, `kubernetes`, `terraform`, `github-actions`, `pnpm`). Instructions about a tool that has no package of its own apply only where the tool's files exist. An organisation's source that depends on these packages must write this repository's URL exactly as above, and repositories subscribing to it name the source `bp` with the same URL.

**Must.** and **Prefer.** are relative to the subscribing repository's owner: **Must.** needs the owner's consent to depart from, **Prefer.** a written reason.

## Packages

Counts are always-loaded / on-demand instructions.

| Package | Description | Content | Dependencies |
| --- | --- | --- | --- |
| `base` | Every repository: how agents work, the coding principles, tests, static analysis, CI, latest versions, Renovate, Testcontainers and EditorConfig | 2 / 9 | – |
| `docker` | Dockerfiles and Docker Compose: pinned base images, non-root, .dockerignore, caching, exec form, build secrets, hadolint, labels and Compose files | 0 / 10 | `base` |
| `github-actions` | GitHub Actions workflows: least-privilege permissions, actions pinned by SHA and image publishing | 0 / 3 | `base` |
| `go` | Go: context, error wrapping, module layout, valid states, table-driven tests, slog usage, go tool and checks | 0 / 8 | `base` |
| `go-cli` | Go command-line tools: the version set at build time and reproducible static release builds | 0 / 2 | `go` |
| `go-server` | Go HTTP services: a store package over sqlc and pgx, Huma usage, goose migrations, graceful shutdown and database tests | 0 / 5 | `go`, `services` |
| `java-spring` | Spring Boot services in Java: injection, typed configuration, feature packages, DTOs, problem details, Actuator, Gradle, JPA, Flyway and Testcontainers usage | 0 / 11 | `services` |
| `kubernetes` | Kubernetes manifests: Kustomize, resources, probes, security context, labels, images, secrets, availability, network policies, config and checks | 0 / 11 | `base` |
| `node` | Node.js and TypeScript repositories: running TypeScript directly | 0 / 1 | `base` |
| `node-library` | Node.js libraries published to npm: release pull requests and Trusted Publishing | 0 / 2 | `node` |
| `node-server` | Node.js servers: JSON logging and the V8 heap limit in containers | 0 / 2 | `node`, `services` |
| `nuxt` | Nuxt apps: directory structure, components, composables, runtimeConfig, server routes and utils, Nuxt UI, i18n and tests | 0 / 9 | `node` |
| `pnpm` | pnpm used well: supply-chain hardening and installing pnpm in images and CI | 0 / 2 | `base` |
| `production` | Applications and services meant for production: domain separate from adapters, minimal images, observability and metrics | 0 / 4 | `base` |
| `python` | Python: the src layout, requires-python, uv, strict mypy and pytest usage | 0 / 5 | `base` |
| `python-fastapi` | Python HTTP APIs with FastAPI: app factory, routers, Pydantic schemas, settings, async handlers, SQLAlchemy, Alembic and tests | 0 / 10 | `python`, `services` |
| `react-spa` | React single-page apps with Vite: components, feature folders, configuration, TanStack Query and Router, forms, i18n, Tailwind CSS and tests | 0 / 11 | `node` |
| `rust` | Rust: workspace and lints, errors, Tokio, tests, cargo-deny configuration and checks | 0 / 7 | `base` |
| `rust-server` | Rust HTTP services: axum, sqlx, tracing and database tests | 0 / 4 | `rust`, `services` |
| `services` | Applications and services that run as a server: sign-in cookies, asynchronous work, migrations fixed forward and owner-scoped queries | 0 / 4 | `base` |
| `terraform` | Terraform repositories: layout, pinned versions, remote state, variables and outputs, for_each, secrets, tflint, tags, naming and module sources | 0 / 12 | `base` |

## Layout

```
packages/<package>/
    manifest.jsonc              description and dependencies
    instructions/<id>.md        front matter description, load (always | on-demand), paths
```

## Writing packages

See `AGENTS.md`. The instructions and skills for writing packages come from the package `authoring` of [codicus-labs/agent-package-manager](https://github.com/codicus-labs/agent-package-manager/tree/main/packages/authoring), installed here with `apm install`.

## License

MIT, see `LICENSE`.
