# Use the official Golang image for building
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary, explicitly naming the binary and placing it in the /app directory
RUN go build -o /app/dev-api

# Use a minimal base image to run the binary
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Install bash for debugging, if necessary
RUN apk add --no-cache bash

# Copy the Go binary from the build stage
COPY --from=builder /app/dev-api .

# Expose the application port
EXPOSE 3002

# Make sure the binary has execute permissions (may not be needed but good for safety)
RUN chmod +x ./dev-api

# Run the Go binary
CMD ["./dev-api"]
