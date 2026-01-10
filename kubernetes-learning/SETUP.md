# Kubernetes Local Setup Guide

This guide will help you set up a local Kubernetes cluster for learning. We'll cover the most popular options for running Kubernetes on your local machine.

## Prerequisites

Before installing Kubernetes, ensure you have:
- **4GB+ RAM** available
- **20GB+ disk space**
- **Virtualization** enabled in BIOS (for minikube/kind)
- **Docker** installed (for Docker Desktop/kind)

## Option 1: Minikube (Recommended for Beginners)

Minikube runs a single-node Kubernetes cluster in a VM or container.

### Installation

**macOS:**
```bash
# Using Homebrew
brew install minikube

# Verify installation
minikube version
```

**Linux:**
```bash
# Download and install
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

# Verify
minikube version
```

**Windows:**
```powershell
# Using Chocolatey
choco install minikube

# Or download installer from:
# https://minikube.sigs.k8s.io/docs/start/
```

### Starting Minikube

```bash
# Start cluster (default: 2 CPUs, 4GB RAM)
minikube start

# Start with custom resources
minikube start --cpus=4 --memory=8192

# Start with specific Kubernetes version
minikube start --kubernetes-version=v1.28.0

# Use Docker driver (instead of VM)
minikube start --driver=docker
```

### Useful Minikube Commands

```bash
# Check status
minikube status

# Stop cluster
minikube stop

# Delete cluster
minikube delete

# Access dashboard
minikube dashboard

# Get cluster IP
minikube ip

# SSH into node
minikube ssh

# Access LoadBalancer services
minikube tunnel

# Access NodePort services
minikube service <service-name>

# View logs
minikube logs
```

### Verify Installation

```bash
# Check kubectl is configured
kubectl cluster-info

# View nodes
kubectl get nodes

# Should see minikube node in Ready state
```

---

## Option 2: kind (Kubernetes in Docker)

kind runs Kubernetes clusters using Docker containers as nodes. Great for testing multi-node setups.

### Installation

**macOS:**
```bash
# Using Homebrew
brew install kind

# Verify
kind version
```

**Linux:**
```bash
# Download binary
curl -Lo ./kind https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind

# Verify
kind version
```

**Windows:**
```powershell
# Using Chocolatey
choco install kind

# Or download from:
# https://kind.sigs.k8s.io/
```

### Starting kind

```bash
# Create cluster
kind create cluster

# Create with custom name
kind create cluster --name my-cluster

# Create multi-node cluster
kind create cluster --name multi-node --config kind-config.yaml
```

**kind-config.yaml** (for multi-node):
```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
  - role: worker
  - role: worker
```

### Useful kind Commands

```bash
# List clusters
kind get clusters

# Get kubeconfig
kind get kubeconfig --name my-cluster

# Delete cluster
kind delete cluster --name my-cluster

# Delete all clusters
kind delete clusters --all

# Load Docker image into cluster
kind load docker-image my-image:tag --name my-cluster
```

---

## Option 3: Docker Desktop

Docker Desktop includes Kubernetes support (easiest if you already use Docker).

### Enabling Kubernetes

1. Open **Docker Desktop** settings
2. Go to **Kubernetes** tab
3. Check **Enable Kubernetes**
4. Click **Apply & Restart**
5. Wait for Kubernetes to start (green indicator)

### Verify

```bash
kubectl config get-contexts

# Should see docker-desktop context
# Switch if needed
kubectl config use-context docker-desktop

kubectl get nodes
# Should see docker-desktop node
```

### Managing

```bash
# Reset Kubernetes cluster
# Docker Desktop → Settings → Kubernetes → Reset Kubernetes Cluster

# Disable/Enable via Docker Desktop settings
```

---

## Option 4: k3d (Lightweight k3s in Docker)

k3d runs k3s (lightweight Kubernetes) in Docker.

### Installation

**macOS/Linux:**
```bash
# Install script
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash

# Or with brew
brew install k3d
```

### Starting k3d

```bash
# Create cluster
k3d cluster create mycluster

# Create with multiple nodes
k3d cluster create mycluster --agents 2

# With port mapping
k3d cluster create mycluster --api-port 6550 -p "8080:80@loadbalancer"
```

### Useful k3d Commands

```bash
# List clusters
k3d cluster list

# Stop cluster
k3d cluster stop mycluster

# Start cluster
k3d cluster start mycluster

# Delete cluster
k3d cluster delete mycluster
```

---

## Installing kubectl

kubectl is the command-line tool for interacting with Kubernetes.

### Installation

**macOS:**
```bash
# Using Homebrew
brew install kubectl

# Verify
kubectl version --client
```

**Linux:**
```bash
# Download latest
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

# Install
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Verify
kubectl version --client
```

**Windows:**
```powershell
# Using Chocolatey
choco install kubernetes-cli

# Verify
kubectl version --client
```

### Configure kubectl

kubectl configuration is usually automatic for the tools above.

**Check configuration:**
```bash
# View config
kubectl config view

# Current context
kubectl config current-context

# Available contexts
kubectl config get-contexts

# Switch context
kubectl config use-context <context-name>
```

---

## Verification Steps

After setting up your cluster, verify everything works:

### 1. Check Cluster Info

```bash
# Should show cluster endpoints
kubectl cluster-info

# View cluster information
kubectl cluster-info dump
```

### 2. Check Nodes

```bash
# List nodes
kubectl get nodes

# Nodes should be in "Ready" state
# NAME       STATUS   ROLES           AGE   VERSION
# minikube   Ready    control-plane   5m    v1.28.0
```

### 3. Check System Pods

```bash
# View system components
kubectl get pods -n kube-system

# All pods should be Running or Completed
```

### 4. Create Test Pod

```bash
# Create test pod
kubectl run nginx --image=nginx

# Check it's running
kubectl get pods

# Clean up
kubectl delete pod nginx
```

### 5. Test kubectl Commands

```bash
# These should all work without errors
kubectl get namespaces
kubectl get all --all-namespaces
kubectl api-resources
```

---

## Common Issues & Solutions

### Minikube won't start

```bash
# Delete and recreate
minikube delete
minikube start

# Try different driver
minikube start --driver=docker

# Check logs
minikube logs
```

### kubectl not connecting

```bash
# Check context
kubectl config current-context

# Set context
kubectl config use-context minikube  # or docker-desktop, kind-<name>

# Verify config
kubectl config view
```

### kind cluster issues

```bash
# Delete and recreate
kind delete cluster --name <name>
kind create cluster --name <name>

# Check Docker is running
docker ps
```

### Not enough resources

```bash
# For minikube, allocate more
minikube delete
minikube start --cpus=4 --memory=8192

# Check system resources
# Close other applications
# Restart Docker/minikube
```

### Pods stuck in Pending

```bash
# Describe pod to see events
kubectl describe pod <pod-name>

# Common causes:
# - Insufficient resources
# - Image pull errors
# - Volume mount issues

# Check node resources
kubectl top nodes  # Requires metrics-server
```

---

## Enabling Add-ons (Minikube)

Minikube provides useful add-ons:

```bash
# List available add-ons
minikube addons list

# Enable metrics-server (for kubectl top)
minikube addons enable metrics-server

# Enable dashboard
minikube addons enable dashboard
minikube dashboard

# Enable ingress
minikube addons enable ingress

# Disable add-on
minikube addons disable <addon-name>
```

---

## Metrics Server (for kubectl top)

### Minikube
```bash
minikube addons enable metrics-server
```

### kind/others
```bash
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

# For local dev, may need to disable TLS:
kubectl patch deployment metrics-server -n kube-system --type='json' -p='[{"op": "add", "path": "/spec/template/spec/containers/0/args/-", "value": "--kubelet-insecure-tls"}]'
```

### Verify
```bash
# Wait a minute, then:
kubectl top nodes
kubectl top pods
```

---

## Useful Tools

### kubectx & kubens
Switch contexts and namespaces easily.

```bash
# macOS
brew install kubectx

# Switch context
kubectx minikube

# Switch namespace
kubens dev
```

### k9s
Terminal UI for Kubernetes.

```bash
# macOS
brew install k9s

# Run
k9s
```

### Helm
Kubernetes package manager.

```bash
# macOS
brew install helm

# Verify
helm version
```

---

## Cluster Cleanup

### Stop but keep cluster

```bash
# Minikube
minikube stop

# kind
kind cluster stop <name>

# Docker Desktop
# Use Docker Desktop UI to disable Kubernetes
```

### Delete cluster completely

```bash
# Minikube
minikube delete

# kind
kind delete cluster --name <name>

# k3d
k3d cluster delete <name>

# Docker Desktop
# Docker Desktop → Settings → Kubernetes → Reset Cluster
```

---

## Next Steps

Once your cluster is running:

1. **Verify setup** using the verification steps above
2. **Read the main README** - [README.md](README.md)
3. **Start with basics** - `cd 01-basics`
4. **Reference kubectl commands** - [KUBECTL_COMMANDS.md](KUBECTL_COMMANDS.md)

---

## Recommendation for This Learning Project

**Best choice: Minikube**
- Easy to set up
- Well-documented
- Includes useful add-ons
- Simulates real cluster behavior
- Easy to reset and restart

**Start command:**
```bash
minikube start --cpus=2 --memory=4096
kubectl get nodes
# You're ready to learn!
```

Happy learning!
