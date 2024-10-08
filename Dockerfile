# Use the official Golang image to build the app
FROM golang:1.23-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o kamal-proxy-sidecar .

# Use a minimal base image to run the application
FROM alpine:latest

RUN apk add --no-cache docker-cli

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/kamal-proxy-sidecar .

CMD ["./kamal-proxy-sidecar"]
