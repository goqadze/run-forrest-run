# Exercise 6: Create a Docker Compose with 3 Services

## Goal
Create a docker-compose.yml with 3 interconnected services: web, api, and database.

## Solution

```bash
cd exercises/solutions/06-three-services

# Start all services
docker compose up -d

# View logs
docker compose logs -f

# Test the services
curl http://localhost:8080  # Web (nginx)
curl http://localhost:3000  # API

# Check service status
docker compose ps

# Stop and cleanup
docker compose down -v
```

## Key Concepts

- `services:` defines multiple containers
- `depends_on:` sets startup order
- Services communicate via service names
- Shared networks enable DNS resolution
- `docker compose up -d` starts all services
- `docker compose down` stops and removes all
