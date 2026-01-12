# Exercise 1: Run Nginx on Port 8080

## Goal
Run an nginx container and access it on port 8080.

## Solution

```bash
# Pull and run nginx, mapping container port 80 to host port 8080
docker run -d --name my-nginx -p 8080:80 nginx:alpine

# Test it
curl http://localhost:8080

# Or open in browser: http://localhost:8080

# View container logs
docker logs my-nginx

# Cleanup
docker stop my-nginx
docker rm my-nginx
```

## Key Concepts

- `-d` runs container in detached mode (background)
- `-p 8080:80` maps host port 8080 to container port 80
- `--name` gives the container a friendly name
- `nginx:alpine` is a smaller nginx image (~23MB vs ~187MB)
