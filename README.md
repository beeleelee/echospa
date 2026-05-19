# echospa

Minimal [Echo v5](https://echo.labstack.com/) server that serves a static SPA.

```
echospa
 └── web/          ← SPA build output (not included)
     └── index.html
```

## Usage

```sh
# Build the binary
make build

# Run (requires ./web/index.html)
make run
```

## Docker

This image acts as a base for SPA projects. Build and push:

```sh
docker build -t ghcr.io/beeleelee/echospa:latest .
docker push ghcr.io/beeleelee/echospa:latest
```

Child SPA projects extend it:

```dockerfile
FROM ghcr.io/beeleelee/echospa:latest
COPY build/ /web/
```

## License

MIT
