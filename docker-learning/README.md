# Docker Learning Project

A comprehensive, hands-on project to learn Docker from basics to production-ready applications.

## Prerequisites

- **Docker** installed (Docker Desktop recommended for beginners)
- **Basic command line** knowledge
- **Text editor** of your choice

Don't have Docker? See [SETUP.md](SETUP.md) for installation instructions.

## Project Structure

```
docker-learning/
├── README.md                    # You are here!
├── DOCKER_COMMANDS.md           # Complete Docker command reference
├── SETUP.md                     # Installation guide
├── 01-basics/                   # Containers and images fundamentals
├── 02-dockerfile/               # Building custom images
├── 03-volumes/                  # Data persistence
├── 04-networking/               # Container networking
├── 05-docker-compose/           # Multi-container applications
├── 06-complete-project/         # Full-stack production app
└── exercises/                   # Practice challenges
```

## Getting Started

### 1. Verify Docker Installation

```bash
# Check Docker is installed
docker --version

# Check Docker is running
docker run hello-world

# View Docker info
docker info
```

### 2. Start Learning

Work through the topics in order:

```bash
cd 01-basics
cat README.md
```

Each directory contains:
- **README.md** - Concept explanations
- **Example files** - Dockerfiles, scripts, applications
- **commands.sh** - Commands to try

## Topics Covered

### 1. Docker Basics (`01-basics/`)
**What you'll learn:**
- What is Docker and why use it
- Containers vs Images vs Registries
- Running your first containers
- Basic Docker commands
- Image management

**Key concepts:**
- Containers are isolated processes
- Images are templates for containers
- Docker Hub is the default registry
- Containers are ephemeral by default

### 2. Dockerfile (`02-dockerfile/`)
**What you'll learn:**
- Writing Dockerfiles
- Understanding layers and caching
- Multi-stage builds
- Best practices and optimization
- Security considerations

**Key concepts:**
- Each instruction creates a layer
- Layer caching speeds up builds
- Order matters for cache efficiency
- Multi-stage builds reduce image size
- Don't run as root

**Examples:** Python Flask app, Go app with multi-stage build, good vs bad practices

### 3. Volumes (`03-volumes/`)
**What you'll learn:**
- Data persistence in Docker
- Volume types (named, bind mounts, tmpfs)
- Volume management
- Data sharing between containers

**Key concepts:**
- Container filesystems are ephemeral
- Volumes persist data
- Bind mounts link host directories
- Named volumes are managed by Docker

### 4. Networking (`04-networking/`)
**What you'll learn:**
- Docker network types
- Container communication
- Port mapping
- Custom networks
- DNS and service discovery

**Key concepts:**
- Default bridge network
- User-defined networks provide DNS
- Containers can talk via network
- Port mapping exposes services

### 5. Docker Compose (`05-docker-compose/`)
**What you'll learn:**
- Multi-container applications
- docker-compose.yml syntax
- Service orchestration
- Environment variables
- Depends_on and health checks

**Key concepts:**
- Compose manages multiple containers
- Services are containers
- Networks are created automatically
- Volumes can be shared

**Examples:** Simple web+database, full MERN/PERN stack

### 6. Complete Project (`06-complete-project/`)
**What you'll learn:**
- Real-world application architecture
- Production best practices
- Nginx reverse proxy
- Environment configuration
- Build automation with Make

**Demonstrates:** Full-stack app with frontend, backend, database, and reverse proxy

## Learning Path

**Recommended progression:**

1. **Setup** - Install Docker (SETUP.md)
2. **Basics** - Understand containers and images (01-basics/)
3. **Dockerfiles** - Build custom images (02-dockerfile/)
4. **Volumes** - Persist data (03-volumes/)
5. **Networking** - Connect containers (04-networking/)
6. **Compose** - Multi-container apps (05-docker-compose/)
7. **Complete Project** - Put it all together (06-complete-project/)
8. **Practice** - Exercises (exercises/)

## Common Docker Patterns

### Running Containers
```bash
# Run container in foreground
docker run nginx

# Run container in background (detached)
docker run -d nginx

# Run with port mapping
docker run -d -p 8080:80 nginx

# Run with name
docker run -d --name my-nginx nginx

# Run with environment variables
docker run -d -e MY_VAR=value nginx

# Run with volume
docker run -d -v my-data:/data nginx
```

### Building Images
```bash
# Build from Dockerfile
docker build -t my-app .

# Build with tag
docker build -t my-app:v1.0 .

# Build with build args
docker build --build-arg VERSION=1.0 -t my-app .

# Build without cache
docker build --no-cache -t my-app .
```

### Managing Containers
```bash
# List running containers
docker ps

# List all containers
docker ps -a

# Stop container
docker stop my-nginx

# Start container
docker start my-nginx

# Remove container
docker rm my-nginx

# Remove running container
docker rm -f my-nginx

# View logs
docker logs my-nginx

# Follow logs
docker logs -f my-nginx

# Execute command in container
docker exec my-nginx ls /app

# Interactive shell
docker exec -it my-nginx /bin/bash
```

### Cleanup
```bash
# Remove stopped containers
docker container prune

# Remove unused images
docker image prune

# Remove unused volumes
docker volume prune

# Remove everything unused
docker system prune -a

# View disk usage
docker system df
```

## Resources

- **[Docker Commands Reference](DOCKER_COMMANDS.md)** - Complete command guide
- **[Setup Guide](SETUP.md)** - Installation instructions
- Official Docker Documentation: https://docs.docker.com/
- Docker Hub: https://hub.docker.com/
- Docker Get Started: https://docs.docker.com/get-started/
- Best Practices: https://docs.docker.com/develop/dev-best-practices/
- Dockerfile Reference: https://docs.docker.com/engine/reference/builder/

## Tips for Learning

- **Start simple** - Master basics before advanced topics
- **Read Dockerfiles** - Understand how images are built
- **Use docker inspect** - See detailed container/image info
- **Check logs often** - Debugging starts with logs
- **Clean up regularly** - Use prune commands
- **Use .dockerignore** - Keep images small
- **Don't run as root** - Security best practice
- **Layer caching** - Order Dockerfile commands wisely
- **Multi-stage builds** - Keep production images small

## Common Gotchas

1. **Port already in use** - Stop conflicting services or use different port
2. **Image not found** - Check image name and tag, pull if needed
3. **Permission denied** - Add user to docker group or use sudo
4. **No space left** - Run `docker system prune`
5. **Container exits immediately** - Check logs with `docker logs`
6. **Changes not reflected** - Rebuild image after Dockerfile changes
7. **Volumes not mounting** - Check paths and permissions
8. **Network issues** - Ensure containers are on same network

## Exercises

1. Run nginx and access it from your browser
2. Build a custom image with your own application
3. Create a container that persists data using volumes
4. Connect two containers via a custom network
5. Write a docker-compose.yml for a multi-tier app
6. Optimize a Dockerfile to reduce image size
7. Implement a multi-stage build
8. Deploy the complete project

See [exercises/README.md](exercises/README.md) for detailed challenges.

## Next Steps

After completing this project:

- **Production deployment** - AWS ECS, Azure Container Instances, Google Cloud Run
- **Orchestration** - Kubernetes, Docker Swarm
- **CI/CD** - GitHub Actions, GitLab CI, Jenkins with Docker
- **Security** - Image scanning, secrets management, rootless Docker
- **Monitoring** - Prometheus, Grafana, ELK stack
- **Advanced networking** - Overlay networks, service mesh
- **Registry** - Private Docker registry, Harbor
- **Certification** - Docker Certified Associate (DCA)

## Troubleshooting

### Docker daemon not running
```bash
# macOS/Windows: Start Docker Desktop
# Linux: sudo systemctl start docker
```

### Cannot connect to Docker daemon
```bash
# Check Docker is running
docker info

# On Linux, add user to docker group
sudo usermod -aG docker $USER
# Then logout and login again
```

### Container keeps restarting
```bash
# Check logs
docker logs <container-name>

# Inspect container
docker inspect <container-name>
```

### Build fails with "no space left on device"
```bash
# Clean up
docker system prune -a
docker volume prune
```

Happy learning with Docker!
