# Build stage
FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -trimpath -o server main.go

# Runtime stage — distroless nonroot is available for amd64 & arm64
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]