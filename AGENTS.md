# echospa

Minimal Echo v5 server that serves a static SPA.

## Commands

```sh
go build -o echospa          # build binary
go run .                     # run dev (no hot-reload)
```

## Architecture

- **`main.go`** — single entrypoint, package `main`.
- Static SPA must live under **`./web`** (configurable) with an `index.html`. Uses Echo's `StaticWithConfig` with `HTML5: true` (history API fallback).
- Server listens on **`:8080`** (configurable) with Gzip middleware and 5s graceful shutdown on SIGINT/SIGTERM.

## Configuration

Both command-line flags and environment variables are supported (flags take priority).

| Setting          | Flag               | Env             | Default  |
|------------------|--------------------|-----------------|----------|
| Listen port      | `--port`           | `PORT`          | `8080`   |
| Static files root| `--static-root`    | `STATIC_ROOT`   | `./web`  |

```sh
# Via flags
echospa --port 9090 --static-root /app/spa

# Via env vars
PORT=9090 STATIC_ROOT=/app/spa echospa
```

## SPA build

Before running, generate or copy an SPA build directory to `./web/`. This is not part of the Go project — it is produced by an external frontend build step (e.g., `npm run build` from a sibling frontend repo).

## Docker

```sh
docker build -t echospa .
```

Multi-stage build (golang:1.26-alpine → alpine:3.21), outputs a ~9 MB image. Child SPA projects `FROM` this image and copy their build output to `/web/`:

```dockerfile
FROM ghcr.io/beeleelee/echospa:latest
COPY build/ /web/
```

## Tests

No tests exist yet.
