# Ramadan Tracker API

A simple REST API built with [Go Fiber](https://gofiber.io/) to track Ramadan targets.

## Requirements

- Go 1.22+
- Docker with Buildx

## Run Locally

```bash
go run main.go
```

## Docker

Build and run with Docker:

```bash
make run               # build + run (default tag: latest)
make run TAG=v1.0.0    # build + run with specific tag
make stop              # stop running container
```

Build only:

```bash
make build             # local platform, loads into Docker
make build TAG=v1.0.0  # with specific tag
```

Build & push multiplatform (`linux/amd64`, `linux/arm64`):

```bash
make build-multi REGISTRY=docker.io/myuser TAG=v1.0.0
```

## API Endpoints

Base URL: `http://localhost:8080`

| Method | Endpoint           | Description        |
|--------|--------------------|--------------------|
| GET    | `/`                | Health check       |
| GET    | `/api/targets`     | Get all targets    |
| GET    | `/api/targets/:id` | Get target by ID   |
| POST   | `/api/targets`     | Create target      |
| PUT    | `/api/targets/:id` | Update target      |
| DELETE | `/api/targets/:id` | Delete target      |

## Environment Variables

| Variable | Default | Description        |
|----------|---------|--------------------|
| `PORT`   | `8080`  | Port to listen on  |
