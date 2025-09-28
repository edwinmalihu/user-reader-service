# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod and go sum files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the Go application
RUN go build -o /api ./main.go

# Final stage
FROM alpine:latest
RUN apk update && apk add tzdata

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /api .

# Expose port (Istio akan mengelola traffic)

# Command to run the executable
CMD ["./api"]