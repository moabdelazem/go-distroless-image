# Stage 1: Build Stage
FROM golang:1.25 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first for better layer caching
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 creates a static binary (no C dependencies)
# ! Required for distroless images which have no libc
RUN CGO_ENABLED=0 go build -o main ./cmd/

# Stage 2: Production Stage
# Use distroless image for minimal attack surface
FROM gcr.io/distroless/static-debian12

# Set working directory to root
WORKDIR /

# Document the port the application listens on
EXPOSE 8080

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /

# Run as non-root user for security
# ! distroless images include a nonroot user by default
USER nonroot:nonroot

# Set the entrypoint to run the application
ENTRYPOINT [ "/main" ]