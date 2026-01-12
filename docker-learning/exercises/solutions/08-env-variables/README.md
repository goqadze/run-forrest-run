# Exercise 8: Use Environment Variables in Docker Compose

## Goal
Learn different ways to configure services using environment variables.

## Methods

1. **Inline in docker-compose.yml** - `environment:` key
2. **From .env file** - Automatic loading
3. **From env_file** - Explicit file reference
4. **Variable substitution** - `${VAR}` syntax

## Solution

```bash
cd exercises/solutions/08-env-variables

# Create .env file
cp .env.example .env

# Start services
docker compose up -d

# Test - see environment variables in action
curl http://localhost:3000

# View resolved config
docker compose config

# Cleanup
docker compose down
```

## Key Concepts

- `.env` files are auto-loaded by Docker Compose
- `env_file:` loads from specified files
- `${VAR:-default}` provides defaults
- Never commit secrets to `.env` files in git
