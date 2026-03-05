# Gunakan Golang 1.22 Alpine yang ukurannya sangat ringan
FROM golang:1.22-alpine

# Set direktori kerja di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum terlebih dahulu untuk caching dependency
COPY go.mod go.sum ./

# Unduh package Go Fiber
RUN go mod download

# Salin sisa file kode (main.go)
COPY . .

# Build aplikasi menjadi binary bernama 'server'
RUN go build -o server main.go

# Beritahu platform bahwa container menggunakan port ini
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["/app/server"]