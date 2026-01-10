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
