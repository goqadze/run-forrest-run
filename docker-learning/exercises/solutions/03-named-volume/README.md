# Exercise 3: Create and Use a Named Volume

## Goal
Create a named volume and demonstrate data persistence across container lifecycles.

## Solution

```bash
# Create a named volume
docker volume create app-data

# Run container with volume mounted
docker run -d --name db-container \
  -v app-data:/var/lib/data \
  alpine sh -c "echo 'Important data!' > /var/lib/data/myfile.txt && sleep 3600"

# Verify data was written
docker exec db-container cat /var/lib/data/myfile.txt

# Remove the container
docker rm -f db-container

# Data persists! Create new container with same volume
docker run --rm -v app-data:/var/lib/data alpine cat /var/lib/data/myfile.txt

# Inspect the volume
docker volume inspect app-data

# Cleanup
docker volume rm app-data
```

## Key Concepts

- Named volumes persist beyond container lifecycle
- `-v volume-name:/path` mounts a named volume
- `docker volume create` explicitly creates volumes
- `docker volume inspect` shows volume details
- Data is stored in Docker's managed directory
