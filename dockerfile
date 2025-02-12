# Stage 1: Build the application
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application (without CGO dependencies)
RUN CGO_ENABLED=0 GOOS=linux go build -o getcountcart .

# Stage 2: Create the final image
FROM alpine:latest

# Install necessary dependencies
RUN apk add --no-cache bash

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/getcountcart .

# Expose the port the app runs on
EXPOSE 8080
# Command to run the application
CMD ["./getcountcart"]