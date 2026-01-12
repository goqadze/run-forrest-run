# Exercise 12: Set Up a Reverse Proxy with Nginx

## Goal
Configure Nginx as a reverse proxy to route traffic to multiple backend services.

## Architecture

```
                    ┌──────────────┐
                    │    Nginx     │
                    │ Reverse Proxy│
                    │    :80       │
                    └──────┬───────┘
                           │
         ┌─────────────────┼─────────────────┐
         │                 │                 │
         ▼                 ▼                 ▼
    ┌─────────┐      ┌─────────┐      ┌─────────┐
    │  App 1  │      │  App 2  │      │  App 3  │
    │  :3001  │      │  :3002  │      │  :3003  │
    └─────────┘      └─────────┘      └─────────┘
    /app1/*          /app2/*          /app3/*
```

## Solution

```bash
cd exercises/solutions/12-reverse-proxy

# Start all services
docker compose up -d

# Test routing
curl http://localhost/           # Static welcome page
curl http://localhost/app1       # Routes to app1
curl http://localhost/app2       # Routes to app2
curl http://localhost/api/       # Routes to api service

# View nginx logs
docker compose logs nginx

# Cleanup
docker compose down
```

## Key Concepts

- `proxy_pass` forwards requests to upstream servers
- `upstream` blocks define backend server pools
- Load balancing across multiple instances
- Path-based routing with `location` blocks
- Header forwarding for client info
