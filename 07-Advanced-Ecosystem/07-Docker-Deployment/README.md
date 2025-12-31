# Docker & Deployment

Learn about containerizing and deploying Go applications with Docker.

## Learning Objectives

- Understand Docker basics for Go
- Learn to create efficient Dockerfiles
- Master multi-stage builds
- Understand Docker best practices
- Learn Docker Compose
- Understand deployment strategies

## Concepts Covered

### Basic Dockerfile

Simple Dockerfile for Go applications:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o myapp
FROM alpine:latest
COPY --from=builder /app/myapp .
CMD ["./myapp"]
```

### Multi-Stage Builds

Minimize image size with multi-stage builds:

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
# ... build steps ...

# Runtime stage
FROM scratch
COPY --from=builder /build/myapp /myapp
ENTRYPOINT ["/myapp"]
```

### Docker Compose

Orchestrate services with Docker Compose:

```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
```

## Running the Example

```bash
# Navigate to this directory
cd 07-Docker-Deployment

# Initialize the module (if not already done)
go mod init github.com/codinsec/go-learning-lab/07-advanced-ecosystem/docker-deployment

# Build Docker image
docker build -t myapp .

# Run container
docker run -p 8080:8080 myapp

# Run with environment variables
docker run -p 8080:8080 -e PORT=8080 -e ENV=production myapp

# Use Docker Compose
docker-compose up
```

## Running Tests

```bash
# Run tests locally
go test

# Run tests in Docker
docker build -t myapp-test .
docker run myapp-test go test
```

## Key Takeaways

1. **Multi-stage builds** - Minimize image size
2. **Alpine base** - Use lightweight base images
3. **.dockerignore** - Exclude unnecessary files
4. **Layer caching** - Optimize build layers
5. **Health checks** - Add health endpoints
6. **Environment variables** - Use for configuration

## Common Patterns

**Minimal image:**
```dockerfile
FROM scratch
COPY myapp /myapp
ENTRYPOINT ["/myapp"]
```

**With certificates:**
```dockerfile
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY myapp /myapp
```

**Build optimizations:**
```dockerfile
# Copy go.mod first for better caching
COPY go.mod go.sum ./
RUN go mod download
COPY . .
```

## Best Practices

- **Multi-stage builds** - Separate build and runtime
- **Minimal base images** - Use alpine or scratch
- **Layer caching** - Order layers for cache efficiency
- **.dockerignore** - Exclude unnecessary files
- **Specific tags** - Use specific Go version tags
- **Non-root user** - Run as non-root when possible
- **Health checks** - Add health endpoints
- **Resource limits** - Set memory/CPU limits

## Important Notes

- **Image size** - Multi-stage builds reduce size significantly
- **Build context** - Only include necessary files
- **CGO** - Disable CGO for smaller images (CGO_ENABLED=0)
- **Static linking** - Use static binaries
- **Security** - Keep base images updated
- **Ports** - Document exposed ports

## Dockerfile Examples

### Example 1: Simple Build

```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o myapp
CMD ["./myapp"]
```

### Example 2: Multi-Stage (Recommended)

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o myapp

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/myapp /myapp
CMD ["./myapp"]
```

### Example 3: Minimal (Scratch)

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o myapp

FROM scratch
COPY --from=builder /build/myapp /myapp
ENTRYPOINT ["/myapp"]
```

## Docker Commands

```bash
# Build image
docker build -t myapp .

# Run container
docker run -p 8080:8080 myapp

# Run with environment
docker run -p 8080:8080 -e PORT=8080 myapp

# Run in background
docker run -d -p 8080:8080 myapp

# View logs
docker logs <container-id>

# Stop container
docker stop <container-id>

# Remove container
docker rm <container-id>

# Remove image
docker rmi myapp
```

## Docker Compose

```bash
# Start services
docker-compose up

# Start in background
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs

# Rebuild
docker-compose build
```

## Deployment Strategies

1. **Docker Hub** - Public/private registry
2. **AWS ECR** - Amazon container registry
3. **Google GCR** - Google container registry
4. **Azure ACR** - Azure container registry
5. **Self-hosted** - Private registry

## Image Size Optimization

- **Multi-stage builds** - Remove build tools
- **Alpine base** - Smaller than Debian
- **Scratch base** - Minimal (no OS)
- **Strip binaries** - Use `-ldflags="-w -s"`
- **Disable CGO** - Smaller binaries
- **.dockerignore** - Exclude files

## Next Steps

After understanding Docker, proceed to:
- [08-Go-Tools](../08-Go-Tools/README.md) - Learn about Go tools

---
**Created by:** [Codinsec](https://codinsec.com) | [info@codinsec.com](mailto:info@codinsec.com)  
**Author:** Barbaros Kaymak

