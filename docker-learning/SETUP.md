# Docker Setup Guide

This guide will help you install Docker on your system.

## Option 1: Docker Desktop (Recommended for Beginners)

Docker Desktop includes Docker Engine, Docker CLI, Docker Compose, and a GUI.

### macOS

**Requirements:**
- macOS 11 or newer
- Apple Silicon (M1/M2) or Intel processor

**Installation:**
1. Download Docker Desktop from https://www.docker.com/products/docker-desktop/
2. Open the `.dmg` file
3. Drag Docker to Applications folder
4. Launch Docker from Applications
5. Follow the setup wizard

**Verify:**
```bash
docker --version
docker run hello-world
```

### Windows

**Requirements:**
- Windows 10 64-bit: Pro, Enterprise, or Education (Build 19041 or higher)
- OR Windows 11
- WSL 2 backend (recommended)

**Installation:**
1. Download Docker Desktop from https://www.docker.com/products/docker-desktop/
2. Run the installer
3. Enable WSL 2 when prompted
4. Restart computer
5. Launch Docker Desktop

**Verify:**
```bash
docker --version
docker run hello-world
```

### Linux (Docker Desktop)

Docker Desktop is available for Linux:
1. Download from https://docs.docker.com/desktop/install/linux-install/
2. Install using package manager
3. Launch Docker Desktop

---

## Option 2: Docker Engine (Linux)

For Linux servers or lightweight installations.

### Ubuntu/Debian

```bash
# Update package index
sudo apt-get update

# Install prerequisites
sudo apt-get install ca-certificates curl gnupg lsb-release

# Add Docker's official GPG key
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

# Set up repository
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker Engine
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin

# Verify installation
sudo docker run hello-world
```

### CentOS/RHEL/Fedora

```bash
# Install required packages
sudo yum install -y yum-utils

# Set up repository
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

# Install Docker Engine
sudo yum install docker-ce docker-ce-cli containerd.io docker-compose-plugin

# Start Docker
sudo systemctl start docker
sudo systemctl enable docker

# Verify
sudo docker run hello-world
```

---

## Post-Installation Steps

### Run Docker without sudo (Linux)

```bash
# Create docker group
sudo groupadd docker

# Add your user to docker group
sudo usermod -aG docker $USER

# Log out and log back in, then verify
docker run hello-world
```

### Configure Docker to start on boot (Linux)

```bash
sudo systemctl enable docker.service
sudo systemctl enable containerd.service
```

---

## Install Docker Compose

### Docker Desktop
Docker Compose is included with Docker Desktop.

### Linux (Plugin)
```bash
# Install Compose plugin
sudo apt-get update
sudo apt-get install docker-compose-plugin

# Verify
docker compose version
```

### Linux (Standalone)
```bash
# Download
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# Make executable
sudo chmod +x /usr/local/bin/docker-compose

# Verify
docker-compose --version
```

---

## Verification

Run these commands to verify your installation:

```bash
# Check Docker version
docker --version

# Check Docker Compose version
docker compose version  # Plugin
docker-compose --version  # Standalone

# Check Docker is running
docker info

# Run test container
docker run hello-world

# Check system info
docker system info
```

---

## Common Issues

### Docker daemon not running

**macOS/Windows:**
- Start Docker Desktop application
- Check system tray/menu bar for Docker icon

**Linux:**
```bash
sudo systemctl start docker
sudo systemctl status docker
```

### Permission denied

**Linux:**
```bash
# Add user to docker group
sudo usermod -aG docker $USER

# Log out and back in, or run:
newgrp docker
```

**macOS/Windows:**
- Ensure Docker Desktop is running

### WSL 2 issues (Windows)

```bash
# Install WSL 2
wsl --install

# Update WSL
wsl --update

# Set WSL 2 as default
wsl --set-default-version 2
```

### Cannot connect to Docker daemon

```bash
# Check Docker is running
docker info

# Restart Docker
# macOS/Windows: Restart Docker Desktop
# Linux: sudo systemctl restart docker
```

---

## Configuration

### Increase resources (Docker Desktop)

1. Open Docker Desktop
2. Go to Settings/Preferences
3. Resources tab
4. Adjust CPUs, Memory, Disk space
5. Apply & Restart

### Configure daemon (Linux)

Edit `/etc/docker/daemon.json`:
```json
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "10m",
    "max-file": "3"
  }
}
```

Restart Docker:
```bash
sudo systemctl restart docker
```

---

## Uninstallation

### macOS
1. Quit Docker Desktop
2. Delete Docker.app from Applications
3. Clean up (optional):
```bash
rm -rf ~/Library/Group\ Containers/group.com.docker
rm -rf ~/Library/Containers/com.docker.docker
rm -rf ~/.docker
```

### Windows
1. Uninstall from Apps & Features
2. Clean up (optional):
```powershell
Remove-Item -Recurse -Force "$env:APPDATA\Docker"
Remove-Item -Recurse -Force "$env:LOCALAPPDATA\Docker"
```

### Linux
```bash
# Ubuntu/Debian
sudo apt-get purge docker-ce docker-ce-cli containerd.io docker-compose-plugin
sudo rm -rf /var/lib/docker
sudo rm -rf /var/lib/containerd

# CentOS/RHEL
sudo yum remove docker-ce docker-ce-cli containerd.io docker-compose-plugin
sudo rm -rf /var/lib/docker
sudo rm -rf /var/lib/containerd
```

---

## Next Steps

Once Docker is installed and verified:

1. Read the main [README.md](README.md)
2. Check [DOCKER_COMMANDS.md](DOCKER_COMMANDS.md) for command reference
3. Start with [01-basics/](01-basics/)

Happy Dockering!
