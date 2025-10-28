# MinusOne Demo Application

A simple Go web application for demonstrating Kargo + Argo Workflows progressive delivery.

## Endpoints

- `/` - Root endpoint with version information
- `/health` - Health check endpoint (JSON)
- `/info` - Application info endpoint (JSON)

## Environment Variables

- `VERSION` - Application version (default: v1.0.0)
- `PORT` - HTTP port (default: 8080)

## Building

```bash
docker build -t minusone-demo:latest .
```

## Running Locally

```bash
docker run -p 8080:8080 -e VERSION=v1.0.0 minusone-demo:latest
```

## Testing

```bash
curl http://localhost:8080/health
```

