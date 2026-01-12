# Stage 1: Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy source code
COPY main.go .

# Build the binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go

# Stage 2: Production stage (minimal image)
FROM alpine:3.19

WORKDIR /app

# Copy only the binary from builder stage
COPY --from=builder /app/server .

# Run as non-root user for security
RUN adduser -D appuser
USER appuser

# Run the application
CMD ["./server"]
