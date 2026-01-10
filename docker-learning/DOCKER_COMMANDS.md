# Docker Commands Reference

A comprehensive reference for Docker CLI commands. Master Docker through practical examples.

## Table of Contents
- [Container Commands](#container-commands)
- [Image Commands](#image-commands)
- [Volume Commands](#volume-commands)
- [Network Commands](#network-commands)
- [Docker Compose Commands](#docker-compose-commands)
- [System Commands](#system-commands)
- [Registry Commands](#registry-commands)
- [Quick Tips](#quick-tips)

---

## Container Commands

### `docker run`
Create and start a new container from an image.

**Syntax:**
```bash
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
```

**Examples:**
```bash
# Run container (foreground)
docker run nginx

# Run in background (detached)
docker run -d nginx

# Run with name
docker run -d --name my-nginx nginx

# Run with port mapping
docker run -d -p 8080:80 nginx
docker run -d -p 127.0.0.1:8080:80 nginx  # Bind to localhost only

# Run with environment variables
docker run -d -e MY_VAR=value -e DEBUG=true nginx

# Run with volume
docker run -d -v my-data:/data nginx
docker run -d -v /host/path:/container/path nginx

# Run interactive terminal
docker run -it ubuntu /bin/bash

# Run and remove when stopped
docker run --rm nginx

# Run with resource limits
docker run -d --memory="512m" --cpus="1.5" nginx

# Run with restart policy
docker run -d --restart=always nginx
docker run -d --restart=on-failure:5 nginx

# Run with custom network
docker run -d --network=my-network nginx

# Override entrypoint
docker run --entrypoint=/bin/sh nginx
```

**Common Flags:**
- `-d, --detach` - Run in background
- `-it` - Interactive terminal
- `-p, --publish` - Port mapping (host:container)
- `-v, --volume` - Mount volume
- `-e, --env` - Set environment variable
- `--name` - Container name
- `--rm` - Remove container when stopped
- `--network` - Connect to network
- `--restart` - Restart policy
- `--memory` - Memory limit
- `--cpus` - CPU limit

---

### `docker ps`
List containers.

**Syntax:**
```bash
docker ps [OPTIONS]
```

**Examples:**
```bash
# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# List with size
docker ps -s

# List only IDs
docker ps -q

# Custom format
docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Status}}"

# Filter by status
docker ps --filter "status=running"
docker ps --filter "status=exited"

# Filter by name
docker ps --filter "name=nginx"

# Latest container
docker ps -l
```

**Common Flags:**
- `-a, --all` - Show all containers
- `-q, --quiet` - Only IDs
- `-s, --size` - Show sizes
- `-l, --latest` - Latest container
- `--filter` - Filter output
- `--format` - Custom format

---

### `docker start / stop / restart`
Control container state.

**Examples:**
```bash
# Start stopped container
docker start my-nginx

# Start multiple containers
docker start container1 container2

# Stop running container
docker stop my-nginx

# Stop with timeout (default 10s)
docker stop -t 30 my-nginx

# Restart container
docker restart my-nginx

# Pause container
docker pause my-nginx

# Unpause container
docker unpause my-nginx

# Kill container (force stop)
docker kill my-nginx
```

---

### `docker exec`
Execute command in running container.

**Syntax:**
```bash
docker exec [OPTIONS] CONTAINER COMMAND [ARG...]
```

**Examples:**
```bash
# Execute command
docker exec my-nginx ls /usr/share/nginx/html

# Interactive shell
docker exec -it my-nginx /bin/bash
docker exec -it my-nginx sh

# Execute as different user
docker exec -u root my-nginx whoami

# Execute with environment variable
docker exec -e MY_VAR=value my-nginx env

# Execute in specific directory
docker exec -w /app my-nginx pwd
```

**Common Flags:**
- `-it` - Interactive terminal
- `-u, --user` - User/UID
- `-e, --env` - Environment variable
- `-w, --workdir` - Working directory
- `-d, --detach` - Detached mode

---

### `docker logs`
View container logs.

**Syntax:**
```bash
docker logs [OPTIONS] CONTAINER
```

**Examples:**
```bash
# View logs
docker logs my-nginx

# Follow logs (like tail -f)
docker logs -f my-nginx

# Last 100 lines
docker logs --tail 100 my-nginx

# Logs since timestamp
docker logs --since 2024-01-01T00:00:00 my-nginx
docker logs --since 1h my-nginx

# Logs until timestamp
docker logs --until 2024-01-02 my-nginx

# Show timestamps
docker logs -t my-nginx

# Follow with tail
docker logs -f --tail 50 my-nginx
```

**Common Flags:**
- `-f, --follow` - Stream logs
- `--tail N` - Show last N lines
- `--since` - Show since timestamp
- `--until` - Show until timestamp
- `-t, --timestamps` - Show timestamps

---

### `docker rm`
Remove containers.

**Examples:**
```bash
# Remove stopped container
docker rm my-nginx

# Force remove running container
docker rm -f my-nginx

# Remove multiple containers
docker rm container1 container2 container3

# Remove all stopped containers
docker container prune

# Remove all containers (stopped)
docker rm $(docker ps -aq)

# Remove with filter
docker rm $(docker ps -aq --filter "status=exited")
```

**Common Flags:**
- `-f, --force` - Force removal
- `-v, --volumes` - Remove volumes too

---

### `docker inspect`
View detailed container/image information.

**Examples:**
```bash
# Inspect container
docker inspect my-nginx

# Inspect image
docker inspect nginx:latest

# Get specific field
docker inspect --format='{{.NetworkSettings.IPAddress}}' my-nginx
docker inspect --format='{{.State.Status}}' my-nginx

# Get multiple fields
docker inspect --format='{{.Name}} {{.State.Status}}' my-nginx
```

---

### `docker stats`
Display live resource usage statistics.

**Examples:**
```bash
# View all running containers
docker stats

# View specific containers
docker stats my-nginx my-app

# No stream (one-time)
docker stats --no-stream

# Custom format
docker stats --format "table {{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}"
```

---

## Image Commands

### `docker build`
Build image from Dockerfile.

**Syntax:**
```bash
docker build [OPTIONS] PATH
```

**Examples:**
```bash
# Build from current directory
docker build .

# Build with tag
docker build -t my-app:latest .
docker build -t my-app:v1.0 .

# Build with multiple tags
docker build -t my-app:latest -t my-app:v1.0 .

# Build with build arguments
docker build --build-arg VERSION=1.0 -t my-app .

# Build without cache
docker build --no-cache -t my-app .

# Build specific Dockerfile
docker build -f Dockerfile.prod -t my-app .

# Build with target stage (multi-stage)
docker build --target production -t my-app .

# Build and show output
docker build --progress=plain -t my-app .

# Build with build context from URL
docker build -t my-app github.com/user/repo

# Build with labels
docker build --label version=1.0 -t my-app .
```

**Common Flags:**
- `-t, --tag` - Image name:tag
- `-f, --file` - Dockerfile name
- `--build-arg` - Build arguments
- `--no-cache` - No cache
- `--target` - Multi-stage target
- `--progress` - Progress output type
- `--platform` - Target platform

---

### `docker pull / push`
Pull/push images from/to registry.

**Examples:**
```bash
# Pull image
docker pull nginx
docker pull nginx:1.21
docker pull nginx:1.21-alpine

# Pull from specific registry
docker pull gcr.io/my-project/my-image

# Pull all tags
docker pull -a nginx

# Push image
docker push my-username/my-app:latest

# Push all tags
docker push -a my-username/my-app
```

---

### `docker images`
List images.

**Examples:**
```bash
# List images
docker images

# List all images (including intermediate)
docker images -a

# List image IDs only
docker images -q

# Filter by name
docker images nginx

# Filter by reference
docker images --filter=reference='nginx:*'

# Show digests
docker images --digests

# Custom format
docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}"
```

---

### `docker tag`
Tag an image.

**Examples:**
```bash
# Tag image
docker tag my-app my-app:v1.0

# Tag for registry
docker tag my-app my-username/my-app:latest

# Tag for private registry
docker tag my-app registry.example.com/my-app:v1.0
```

---

### `docker rmi`
Remove images.

**Examples:**
```bash
# Remove image
docker rmi nginx:latest

# Force remove
docker rmi -f nginx:latest

# Remove multiple images
docker rmi nginx:latest redis:alpine

# Remove all images
docker rmi $(docker images -q)

# Remove dangling images
docker image prune

# Remove all unused images
docker image prune -a
```

---

### `docker history`
Show image layer history.

**Examples:**
```bash
# View image history
docker history nginx

# Show full commands
docker history --no-trunc nginx

# Human-readable sizes
docker history -H nginx
```

---

## Volume Commands

### `docker volume create`
Create a volume.

**Examples:**
```bash
# Create volume
docker volume create my-data

# Create with driver options
docker volume create --driver local my-data

# Create with labels
docker volume create --label env=prod my-data
```

---

### `docker volume ls`
List volumes.

**Examples:**
```bash
# List volumes
docker volume ls

# Filter by name
docker volume ls --filter name=my-

# Filter by dangling
docker volume ls --filter dangling=true

# Custom format
docker volume ls --format "table {{.Name}}\t{{.Driver}}"
```

---

### `docker volume inspect`
View volume details.

**Examples:**
```bash
# Inspect volume
docker volume inspect my-data

# Get specific field
docker volume inspect --format='{{.Mountpoint}}' my-data
```

---

### `docker volume rm`
Remove volumes.

**Examples:**
```bash
# Remove volume
docker volume rm my-data

# Remove multiple volumes
docker volume rm vol1 vol2 vol3

# Remove all unused volumes
docker volume prune

# Force prune (no prompt)
docker volume prune -f
```

---

## Network Commands

### `docker network create`
Create a network.

**Examples:**
```bash
# Create network
docker network create my-network

# Create with driver
docker network create --driver bridge my-network

# Create with subnet
docker network create --subnet=172.18.0.0/16 my-network

# Create with gateway
docker network create --subnet=172.18.0.0/16 --gateway=172.18.0.1 my-network
```

---

### `docker network ls`
List networks.

**Examples:**
```bash
# List networks
docker network ls

# Filter by name
docker network ls --filter name=my-

# Filter by driver
docker network ls --filter driver=bridge
```

---

### `docker network connect / disconnect`
Connect/disconnect container to/from network.

**Examples:**
```bash
# Connect container to network
docker network connect my-network my-container

# Connect with alias
docker network connect --alias web my-network my-container

# Disconnect
docker network disconnect my-network my-container
```

---

### `docker network inspect`
View network details.

**Examples:**
```bash
# Inspect network
docker network inspect my-network

# View containers on network
docker network inspect --format='{{range .Containers}}{{.Name}} {{end}}' my-network
```

---

### `docker network rm`
Remove networks.

**Examples:**
```bash
# Remove network
docker network rm my-network

# Remove all unused networks
docker network prune
```

---

## Docker Compose Commands

### `docker-compose up`
Create and start services.

**Examples:**
```bash
# Start services
docker-compose up

# Start in background
docker-compose up -d

# Build before starting
docker-compose up --build

# Force recreate
docker-compose up --force-recreate

# Scale service
docker-compose up --scale web=3

# Specific services only
docker-compose up web db
```

---

### `docker-compose down`
Stop and remove services.

**Examples:**
```bash
# Stop and remove
docker-compose down

# Remove volumes too
docker-compose down -v

# Remove images too
docker-compose down --rmi all
```

---

### `docker-compose ps`
List services.

**Examples:**
```bash
# List services
docker-compose ps

# List all (including stopped)
docker-compose ps -a
```

---

### `docker-compose logs`
View service logs.

**Examples:**
```bash
# View logs
docker-compose logs

# Follow logs
docker-compose logs -f

# Specific service
docker-compose logs web

# Tail last 100 lines
docker-compose logs --tail=100 web
```

---

### `docker-compose exec`
Execute command in service.

**Examples:**
```bash
# Execute command
docker-compose exec web ls /app

# Interactive shell
docker-compose exec web /bin/bash
```

---

### `docker-compose build`
Build services.

**Examples:**
```bash
# Build all services
docker-compose build

# Build specific service
docker-compose build web

# Build without cache
docker-compose build --no-cache

# Build with build args
docker-compose build --build-arg VERSION=1.0
```

---

## System Commands

### `docker system df`
Show Docker disk usage.

**Examples:**
```bash
# Show disk usage
docker system df

# Verbose output
docker system df -v
```

---

### `docker system prune`
Remove unused data.

**Examples:**
```bash
# Remove stopped containers, unused networks, dangling images
docker system prune

# Remove all unused images (not just dangling)
docker system prune -a

# Force (no prompt)
docker system prune -f

# Remove volumes too
docker system prune --volumes
```

---

### `docker info`
Display system information.

**Examples:**
```bash
# System info
docker info

# Specific format
docker info --format '{{.DriverStatus}}'
```

---

### `docker version`
Show Docker version.

**Examples:**
```bash
# Version info
docker version

# Short format
docker version --format '{{.Server.Version}}'
```

---

## Registry Commands

### `docker login / logout`
Login/logout from registry.

**Examples:**
```bash
# Login to Docker Hub
docker login

# Login with username
docker login -u myusername

# Login to private registry
docker login registry.example.com

# Logout
docker logout

# Logout from private registry
docker logout registry.example.com
```

---

## Quick Tips

### Most Used Commands
```bash
docker ps                    # List containers
docker images                # List images
docker run -d nginx          # Run container
docker stop <container>      # Stop container
docker rm <container>        # Remove container
docker rmi <image>           # Remove image
docker logs -f <container>   # Follow logs
docker exec -it <container> sh  # Shell into container
```

### Cleanup Commands
```bash
# Remove stopped containers
docker container prune

# Remove unused images
docker image prune -a

# Remove unused volumes
docker volume prune

# Remove unused networks
docker network prune

# Remove everything
docker system prune -a --volumes
```

### Useful Aliases
Add to `.bashrc` or `.zshrc`:
```bash
alias d='docker'
alias dc='docker-compose'
alias dps='docker ps'
alias dpa='docker ps -a'
alias di='docker images'
alias dex='docker exec -it'
alias dlogs='docker logs -f'
alias dprune='docker system prune -af --volumes'
```

### Quick Container Shell
```bash
# Alpine/BusyBox
docker exec -it <container> sh

# Ubuntu/Debian
docker exec -it <container> bash

# As root
docker exec -it -u root <container> bash
```

### View Container IP
```bash
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <container>
```

### Copy files to/from container
```bash
# To container
docker cp file.txt my-container:/app/

# From container
docker cp my-container:/app/file.txt ./
```

### Debug failed build
```bash
# Build and stop at failing step
docker build --target <stage> -t debug .

# Run intermediate image
docker run -it <image-id> sh
```

---

## See Also

- [Docker Documentation](https://docs.docker.com/)
- [Docker Hub](https://hub.docker.com/)
- [Dockerfile Reference](https://docs.docker.com/engine/reference/builder/)
- [Compose File Reference](https://docs.docker.com/compose/compose-file/)
- [README.md](README.md) - Project learning guide
- [SETUP.md](SETUP.md) - Installation guide
