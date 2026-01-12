# Exercise 9: Build a Full-Stack Application with Docker Compose

## Goal
Create a complete full-stack application with frontend, backend API, and database.

## Architecture

```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   Frontend  │───▶│   Backend   │───▶│  Database   │
│   (React)   │    │  (Node.js)  │    │ (PostgreSQL)│
│   :3000     │    │   :4000     │    │   :5432     │
└─────────────┘    └─────────────┘    └─────────────┘
```

## Solution

```bash
cd exercises/solutions/09-full-stack

# Build and start all services
docker compose up -d --build

# Access the application
# Frontend: http://localhost:3000
# API: http://localhost:4000/api/todos

# View logs
docker compose logs -f

# Cleanup
docker compose down -v
```

## Key Concepts

- Build custom images with `build:` directive
- Service dependencies with `depends_on`
- Persistent data with named volumes
- Network isolation
- Environment-based configuration
