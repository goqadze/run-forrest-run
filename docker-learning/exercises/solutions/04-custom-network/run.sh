#!/bin/bash
# Exercise 4: Custom Network Demo

echo "=== Custom Network Demonstration ==="

echo "1. Creating custom network..."
docker network create demo-network

echo "2. Starting web server..."
docker run -d --name webserver --network demo-network nginx:alpine

echo "3. Starting client container..."
docker run -d --name client --network demo-network alpine sleep 3600

echo "4. Testing DNS resolution (ping webserver by name)..."
docker exec client ping -c 3 webserver

echo "5. Testing HTTP connection..."
docker exec client wget -qO- http://webserver | head -5

echo ""
echo "Press Enter to cleanup..."
read

docker rm -f webserver client
docker network rm demo-network
echo "Cleanup complete!"
