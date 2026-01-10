#!/bin/bash
set -e
BASE="/Users/avtandilgokadze/Learning/docker-learning"

echo "Creating all Docker learning examples..."

# 01-basics
cat > "$BASE/01-basics/README.md" << 'README01'
# Docker Basics

Learn Docker fundamentals: containers, images, and essential commands.

## Key Concepts

- **Container**: Running instance of an image
- **Image**: Template for creating containers  
- **Registry**: Storage for images (Docker Hub)
- **Dockerfile**: Instructions to build an image

## Files

- `commands.sh` - Essential Docker commands to practice

## Practice

```bash
# Run commands interactively
bash commands.sh
```
README01

cat > "$BASE/01-basics/commands.sh" << 'CMD01'
#!/bin/bash

echo "=== Docker Basics Practice ==="

echo "1. Pull an image..."
docker pull nginx:alpine

echo "2. List images..."
docker images

echo "3. Run container..."
docker run -d --name my-nginx -p 8080:80 nginx:alpine

echo "4. List running containers..."
docker ps

echo "5. View logs..."
docker logs my-nginx

echo "6. Exec into container..."
docker exec my-nginx ls /usr/share/nginx/html

echo "7. Stop container..."
docker stop my-nginx

echo "8. Remove container..."
docker rm my-nginx

echo "9. Remove image..."
docker rmi nginx:alpine

echo "Done! Try these commands individually to learn Docker basics."
CMD01

chmod +x "$BASE/01-basics/commands.sh"

# 02-dockerfile
cat > "$BASE/02-dockerfile/README.md" << 'README02'
# Dockerfiles

Learn to build custom Docker images.

## Topics

- Dockerfile syntax
- Layer caching
- Multi-stage builds
- Best practices

## Examples

- `simple-app/` - Python Flask application
- `multi-stage/` - Go app with multi-stage build
- `best-practices/` - Good vs bad examples
README02

cat > "$BASE/02-dockerfile/simple-app/Dockerfile" << 'DF1'
FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY app.py .

EXPOSE 5000

CMD ["python", "app.py"]
DF1

cat > "$BASE/02-dockerfile/simple-app/app.py" << 'APP1'
from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
    return "Hello from Docker!"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
APP1

cat > "$BASE/02-dockerfile/simple-app/requirements.txt" << 'REQ1'
Flask==3.0.0
REQ1

cat > "$BASE/02-dockerfile/multi-stage/Dockerfile" << 'DF2'
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
DF2

cat > "$BASE/02-dockerfile/multi-stage/main.go" << 'GO1'
package main

import "fmt"

func main() {
    fmt.Println("Hello from multi-stage Docker build!")
}
GO1

# 03-volumes
cat > "$BASE/03-volumes/README.md" << 'README03'
# Docker Volumes

Learn data persistence in Docker.

## Volume Types

- Named volumes
- Bind mounts
- tmpfs mounts

## Examples

Practice with the commands.sh script.
README03

cat > "$BASE/03-volumes/commands.sh" << 'CMD03'
#!/bin/bash

echo "=== Docker Volumes Practice ==="

echo "1. Create named volume..."
docker volume create my-data

echo "2. List volumes..."
docker volume ls

echo "3. Run container with volume..."
docker run -d --name vol-test -v my-data:/data alpine sleep 1000

echo "4. Write data to volume..."
docker exec vol-test sh -c "echo 'Persistent data!' > /data/file.txt"

echo "5. Read data..."
docker exec vol-test cat /data/file.txt

echo "6. Remove container..."
docker rm -f vol-test

echo "7. Create new container with same volume..."
docker run --rm -v my-data:/data alpine cat /data/file.txt

echo "8. Cleanup..."
docker volume rm my-data

echo "Data persisted across container lifecycles!"
CMD03

chmod +x "$BASE/03-volumes/commands.sh"

# 04-networking
cat > "$BASE/04-networking/README.md" << 'README04'
# Docker Networking

Learn container networking and communication.

## Network Types

- Bridge (default)
- Host
- None
- Custom networks

## Practice

Run commands.sh to see networking in action.
README04

cat > "$BASE/04-networking/commands.sh" << 'CMD04'
#!/bin/bash

echo "=== Docker Networking Practice ==="

echo "1. List networks..."
docker network ls

echo "2. Create custom network..."
docker network create my-network

echo "3. Run containers on network..."
docker run -d --name web --network my-network nginx:alpine
docker run -d --name app --network my-network alpine sleep 1000

echo "4. Test connectivity (DNS works!)..."
docker exec app ping -c 3 web

echo "5. Inspect network..."
docker network inspect my-network

echo "6. Cleanup..."
docker rm -f web app
docker network rm my-network

echo "Containers can communicate via DNS on custom networks!"
CMD04

chmod +x "$BASE/04-networking/commands.sh"

# 05-docker-compose
cat > "$BASE/05-docker-compose/README.md" << 'README05'
# Docker Compose

Learn multi-container applications with Docker Compose.

## Topics

- docker-compose.yml syntax
- Service definitions
- Networks and volumes
- Environment variables

## Examples

- `simple-app/` - Web + Database
README05

cat > "$BASE/05-docker-compose/simple-app/docker-compose.yml" << 'DC1'
version: '3.8'

services:
  web:
    image: nginx:alpine
    ports:
      - "8080:80"
    networks:
      - app-network

  db:
    image: postgres:15-alpine
    environment:
      POSTGRES_PASSWORD: example
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - app-network

networks:
  app-network:

volumes:
  db-data:
DC1

# exercises
cat > "$BASE/exercises/README.md" << 'EXER'
# Docker Exercises

Practice your Docker skills!

## Beginner

1. Run nginx and access it on port 8080
2. Build a custom image with a simple web app
3. Create and use a named volume
4. Connect two containers on a custom network

## Intermediate

5. Write a multi-stage Dockerfile
6. Create a docker-compose.yml with 3 services
7. Optimize a Dockerfile to reduce image size
8. Use environment variables in Docker Compose

## Advanced

9. Build a full-stack application with Docker Compose
10. Implement health checks
11. Use secrets management
12. Set up a reverse proxy with Nginx

Good luck!
EXER

echo "✅ All examples created successfully!"
