# Exercise 11: Use Secrets Management

## Goal
Learn to securely manage sensitive data like passwords, API keys, and certificates.

## Methods

1. **Docker Secrets** (Swarm mode)
2. **Environment files** (Docker Compose)
3. **External secrets** (HashiCorp Vault, AWS Secrets Manager)

## Solution

### Method 1: Docker Compose with Secret Files

```bash
cd exercises/solutions/11-secrets

# Create secret files (in production, these would be managed securely)
echo "super_secret_password" > secrets/db_password.txt
echo "my_api_key_12345" > secrets/api_key.txt

# Start services
docker compose up -d

# Verify secrets are mounted
docker exec secrets-app cat /run/secrets/db_password

# Cleanup
docker compose down
rm secrets/*.txt
```

### Method 2: Docker Swarm Secrets

```bash
# Initialize swarm (if not already)
docker swarm init

# Create secrets
echo "my_password" | docker secret create db_password -

# List secrets
docker secret ls

# Deploy stack
docker stack deploy -c docker-compose.swarm.yml myapp
```

## Key Concepts

- Secrets are mounted as files in `/run/secrets/`
- Never hardcode secrets in images or compose files
- Use secret files or external secret managers
- Secrets are encrypted at rest in Swarm mode
