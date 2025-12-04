# DockerExample

Simple Go webserver with a built-in HTML page, packaged for Docker and Docker Compose.

## Endpoints
- `GET /` renders the HTML landing page with a health check button.
- `GET /healthz` returns `{"status":"ok"}` for liveness checks.

## Run locally
```
PORT=8080 EXAMPLE1=olaamigo EXAMPLE2=hola-amigo HOST_PORT=8080 go run main.go
```

## Docker
```
docker build -t dockerexample .
HOST_PORT=8080 PORT=8080 EXAMPLE1=olaamigo EXAMPLE2=hola-amigo \
  docker run --rm \
  -e PORT -e HOST_PORT -e EXAMPLE1 -e EXAMPLE2 \
  -p ${HOST_PORT:-8080}:${PORT:-8080} \
  dockerexample
```

## Docker Compose
```
HOST_PORT=8080 PORT=8080 EXAMPLE1=olaamigo EXAMPLE2=hola-amigo docker compose up --build
```

## Environment variables
- `PORT` — container listening port (defaults to 8080).
- `HOST_PORT` — host-side port exposed by Docker/Compose (defaults to 8080).
- `EXAMPLE1`, `EXAMPLE2` — arbitrary values displayed on the landing page (e.g., set `EXAMPLE1=olaamigo` to see it in the UI).
