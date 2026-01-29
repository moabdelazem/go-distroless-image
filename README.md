# Go Distroless API

A simple Go API demonstrating how to build and deploy applications using distroless Docker images.

## Why Distroless?

Distroless images contain only your application and its runtime dependencies - no shell, package manager, or other OS utilities. This results in:
- Smaller image sizes
- Reduced attack surface
- Fewer CVEs to patch

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Returns welcome message |
| `/health` | GET | Returns service health status |

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | `:8080` | Port to listen on |
| `APP_NAME` | `Distroless API` | Service name in health response |

## Make Commands

```bash
make build          # Build Go binary
make docker-build   # Build Docker image
make docker-run     # Run container (default port 8080)
make docker-run PORT=3000   # Run on custom port
```

## Testing

```bash
# Build and run
make docker-build
make docker-run

# Test endpoints
curl http://localhost:8080/
curl http://localhost:8080/health
```
