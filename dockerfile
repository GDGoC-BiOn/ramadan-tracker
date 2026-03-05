# Build stage — runs natively on the build machine for speed
FROM --platform=$BUILDPLATFORM golang:1.25 AS builder

# BuildKit automatically provides these ARGs for cross-compilation
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Cross-compile for the target platform with CGO disabled
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -a -installsuffix cgo -trimpath -o server main.go

# Runtime stage — distroless nonroot is available for amd64 & arm64
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]