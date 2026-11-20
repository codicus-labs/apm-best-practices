---
description: "axum HTTP services: router function, extractors, error type implementing IntoResponse"
load: on-demand
paths: ["**/main.rs", "**/*route*.rs", "**/routes/**/*.rs", "**/*handler*.rs", "**/handlers/**/*.rs", "**/api/**/*.rs", "**/*error*.rs"]
---

**Prefer.** When an HTTP service uses axum, the router is built in `fn router(state: AppState) -> Router`, so tests can call it without a socket; handlers take extractors (`State`, `Path`, `Json`) and return `Result<impl IntoResponse, AppError>`, where `AppError` implements `IntoResponse`.
