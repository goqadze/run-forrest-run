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
