# echospa

Minimal Echo v5 server that serves a static SPA.

## Commands

```sh
go build -o echospa          # build binary
go run .                     # run dev (no hot-reload)
```

## Architecture

- **`main.go`** — single entrypoint, package `main`.
- Static SPA must live under **`/web`** (relative to binary root) with an `index.html`. Uses Echo's `StaticWithConfig` with `HTML5: true` (history API fallback).
- Server listens on **`:8080`** with Gzip middleware and 5s graceful shutdown on SIGINT/SIGTERM.

## SPA build

Before running, generate or copy an SPA build directory to `./web/`. This is not part of the Go project — it is produced by an external frontend build step (e.g., `npm run build` from a sibling frontend repo).

## Tests

No tests exist yet.
