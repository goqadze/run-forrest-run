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
