# Start from the official Go image
FROM golang:1.24 AS builder

WORKDIR /app

# Copy go mod and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go binary (main.go is at project root)
RUN go build -o main main.go

# Small final image
FROM debian:bookworm-slim

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]