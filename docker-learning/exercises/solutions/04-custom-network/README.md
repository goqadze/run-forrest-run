# Exercise 4: Connect Two Containers on a Custom Network

## Goal
Create a custom network and connect two containers that can communicate via DNS.

## Solution

```bash
# Create custom bridge network
docker network create my-app-network

# Run first container (web server)
docker run -d --name web-server \
  --network my-app-network \
  nginx:alpine

# Run second container (client)
docker run -d --name client \
  --network my-app-network \
  alpine sleep 3600

# Test connectivity using container name (DNS)
docker exec client ping -c 3 web-server

# Test HTTP connectivity
docker exec client wget -qO- http://web-server

# Inspect the network
docker network inspect my-app-network

# Cleanup
docker rm -f web-server client
docker network rm my-app-network
```

## Key Concepts

- Custom networks provide automatic DNS resolution
- Containers can communicate using container names
- `--network` attaches container to a network
- Bridge networks isolate container groups
- Default bridge network does NOT provide DNS
