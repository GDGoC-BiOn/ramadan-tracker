# Ramadan Tracker API

A simple REST API built with [Go Fiber](https://gofiber.io/) to track Ramadan targets.

## Requirements

- Go 1.25
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

## Build

Build image for local platform and load into Docker:

```bash
make build             # local platform, loads into Docker
make build TAG=v1.0.0  # with specific tag
```

Build & push multiplatform (`linux/amd64`, `linux/arm64`):

```bash
make build-multi REGISTRY=docker.io/myuser TAG=v1.0.0
```

## Clean

Remove the buildx builder and clean up:

```bash
make clean
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

## Deployment to Google Cloud Run

This project includes a `cloudbuild.yaml` file to automate building the Docker image, pushing it to Google Cloud Artifact Registry, and deploying it to Google Cloud Run.

### Prerequisites

1.  **Google Cloud Project**: You need an active Google Cloud project.
2.  **Enable APIs**: Ensure the **Cloud Build API**, **Artifact Registry API**, and **Cloud Run API** are enabled in your project.
3.  **Artifact Registry Repository**: Create a Docker repository in Artifact Registry.
    ```bash
    gcloud artifacts repositories create demo-bion \
      --repository-format=docker \
      --location=asia-southeast1 \
      --description="Docker repository for demo-bion"
    ```
4.  **gcloud CLI**: Install and configure the Google Cloud SDK (`gcloud`).

### Deploying

You can submit a build using the attached `cloudbuild.yaml` and specify the version tag:

```bash
gcloud builds submit --config=cloudbuild.yaml --substitutions=_TAG="1.0.0" .
```

If you don't provide a tag, it defaults to `latest`.

This command will automatically:
1. Build the Docker container image.
2. Push the image to the Artifact Registry repository (`asia-southeast1-docker.pkg.dev/lab-islam/demo-bion/ramadan-tracker:1.0.0`).
3. Deploy the container to a Cloud Run service named `demo-bion`.

## Clean Google Cloud Resources

To avoid incurring unexpected charges, you can delete the deployment infrastructure:

```bash
# Delete the Artifact Registry repository
gcloud artifacts repositories delete demo-bion --location=asia-southeast1 --quiet

# Delete the Cloud Run service
gcloud run services delete demo-bion --region=asia-southeast1 --quiet
```
