# Use a minimal base image for Go applications
FROM golang:1.21.4-alpine as builder

# Set necessary environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Install necessary build tools (including GCC) for sqllite (and CGO_ENABLED=1)
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /build

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY cmd/ ./cmd/
COPY internal/ ./internal/

# Build the Go API
RUN go build -o main ./cmd/api

# Use a minimal base image for the final container
FROM alpine:latest

# Set working directory for the application
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /build/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]