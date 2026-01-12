# Exercise 10: Implement Health Checks

## Goal
Add health checks to containers to monitor application health and enable automatic recovery.

## Health Check Types

1. **Dockerfile HEALTHCHECK** - Built into image
2. **Docker Compose healthcheck** - Defined in compose file
3. **Orchestrator health checks** - Swarm/Kubernetes probes

## Solution

```bash
cd exercises/solutions/10-health-checks

# Start services with health checks
docker compose up -d

# Watch health status
watch docker compose ps

# View health check logs
docker inspect --format='{{json .State.Health}}' health-api | jq

# Simulate unhealthy state
docker exec health-api touch /tmp/unhealthy
# Watch the container become unhealthy and restart

# Cleanup
docker compose down
```

## Key Concepts

- `test:` command to check health
- `interval:` how often to check
- `timeout:` max time for check to complete
- `retries:` failures before unhealthy
- `start_period:` grace period at startup
- `depends_on.condition: service_healthy`
