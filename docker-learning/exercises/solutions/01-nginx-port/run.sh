#!/bin/bash
# Exercise 1: Run nginx on port 8080

echo "Starting nginx on port 8080..."
docker run -d --name ex1-nginx -p 8080:80 nginx:alpine

echo "Nginx is running! Access it at http://localhost:8080"
echo ""
echo "Press Enter to stop and cleanup..."
read

docker stop ex1-nginx
docker rm ex1-nginx
echo "Cleanup complete!"
