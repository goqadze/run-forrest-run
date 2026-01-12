#!/bin/bash
# Exercise 3: Named Volume Demo

echo "=== Named Volume Demonstration ==="

echo "1. Creating named volume 'app-data'..."
docker volume create app-data

echo "2. Running container and writing data..."
docker run --rm -v app-data:/data alpine sh -c "echo 'Data persists!' > /data/test.txt"

echo "3. Data written. Verifying with new container..."
docker run --rm -v app-data:/data alpine cat /data/test.txt

echo "4. Volume info:"
docker volume inspect app-data

echo ""
echo "Press Enter to cleanup..."
read

docker volume rm app-data
echo "Cleanup complete!"
