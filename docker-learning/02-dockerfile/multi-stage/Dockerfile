# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o app main.go

# Production stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
