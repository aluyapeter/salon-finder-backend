# Step 1: Build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod ./
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the binary
RUN go build -o api ./cmd/api

# Step 2: Run in a lightweight container
FROM debian:bullseye-slim

WORKDIR /app

# Copy only the binary from builder
COPY --from=builder /app/api .

# Expose API port (adjust if your app uses another)
EXPOSE 8080

# Run the binary
CMD ["./api"]
