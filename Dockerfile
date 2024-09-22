# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code
COPY SMS-Server.go .

# Build the Go application
RUN go build -o PostApi-server .

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/server /app/server

# Copy the config file into the container
COPY config.json /app/config.json

# Expose the application port
EXPOSE 8080

# Command to run the binary
CMD ["./PostApi-server"]
