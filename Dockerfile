#Build stage
FROM golang:1.24-alpine AS builder
# Set working directory inside the container
WORKDIR /src
# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# Copy the entire source code into the container
COPY . .
# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/cmd/api/api ./cmd/api/

# Runtime stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/cmd/api/api /usr/local/bin/greenlight
EXPOSE 4000
ENTRYPOINT ["/usr/local/bin/greenlight"]