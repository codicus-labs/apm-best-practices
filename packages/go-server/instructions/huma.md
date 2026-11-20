---
description: "Using Huma in Go HTTP APIs: registering operations on net/http's ServeMux, OpenAPI generated from handler types"
load: on-demand
paths: ["**/*handler*.go", "**/*server*.go", "**/*routes*.go", "**/http/**/*.go", "**/api/**/*.go", "**/openapi*.yaml", "**/openapi*.json"]
---

**Prefer.** When a Go API uses Huma, its operations are registered with `huma.Register` on a Huma adapter of the service's router (`humago` for `http.ServeMux`), with typed input and output structs, so the OpenAPI description is generated from the handlers. A checked-in OpenAPI file is the output of that generation (a command that prints it), never edited by hand.
