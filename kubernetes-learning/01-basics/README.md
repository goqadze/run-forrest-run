# Kubernetes Basics

Welcome to Kubernetes! This section covers the fundamental building blocks: Pods, Containers, and Namespaces.

## What is Kubernetes?

Kubernetes (K8s) is an open-source container orchestration platform that automates:
- **Deployment** - Rolling out applications
- **Scaling** - Adding/removing replicas
- **Management** - Health checks, restarts, load balancing

## Key Concepts

### Pods
- **Smallest deployable unit** in Kubernetes
- Contains one or more containers
- Containers in a Pod share:
  - Network (same IP address)
  - Storage (can mount same volumes)
  - Lifecycle (created and destroyed together)

### Containers
- Lightweight, standalone packages of software
- Include code, runtime, libraries, dependencies
- Run inside Pods

### Namespaces
- Virtual clusters within a physical cluster
- Organize and isolate resources
- Useful for teams, environments (dev/staging/prod)

## Files in This Directory

1. **simple-pod.yaml** - Basic single-container Pod
2. **multi-container-pod.yaml** - Pod with multiple containers (sidecar pattern)
3. **namespace.yaml** - Create a custom namespace
4. **commands.sh** - kubectl commands to try

## Learning Objectives

By the end of this section, you'll understand:
- What Pods are and why they exist
- How to create and manage Pods
- Multi-container Pod patterns
- Namespace usage and benefits
- Basic kubectl commands

## Common Pod Patterns

### Single Container Pod
Most common pattern - one application per Pod.

### Multi-Container Pod Patterns:

**Sidecar:**
Helper container supporting main container (logging, monitoring)

**Ambassador:**
Proxy container connecting main container to outside world

**Adapter:**
Standardizes output from main container

## Hands-On Practice

### 1. Create a Simple Pod

```bash
# View the YAML first
cat simple-pod.yaml

# Create the Pod
kubectl apply -f simple-pod.yaml

# Check Pod status
kubectl get pods

# Get detailed information
kubectl describe pod nginx-pod

# View logs
kubectl logs nginx-pod

# Delete the Pod
kubectl delete -f simple-pod.yaml
```

### 2. Create Multi-Container Pod

```bash
# Apply multi-container Pod
kubectl apply -f multi-container-pod.yaml

# View both containers
kubectl get pod multi-container-pod

# Check logs from specific container
kubectl logs multi-container-pod -c nginx
kubectl logs multi-container-pod -c sidecar

# Exec into specific container
kubectl exec -it multi-container-pod -c nginx -- /bin/sh

# Clean up
kubectl delete -f multi-container-pod.yaml
```

### 3. Work with Namespaces

```bash
# View default namespaces
kubectl get namespaces

# Create custom namespace
kubectl apply -f namespace.yaml

# List namespaces
kubectl get ns

# Create Pod in specific namespace
kubectl apply -f simple-pod.yaml -n learning

# View Pods in namespace
kubectl get pods -n learning

# Delete namespace (deletes all resources in it!)
kubectl delete namespace learning
```

## Common Commands

```bash
# Get Pods
kubectl get pods
kubectl get pods -o wide
kubectl get pods --all-namespaces

# Describe Pod
kubectl describe pod <pod-name>

# Logs
kubectl logs <pod-name>
kubectl logs <pod-name> -f  # Follow logs

# Exec into Pod
kubectl exec -it <pod-name> -- /bin/sh

# Delete Pod
kubectl delete pod <pod-name>
kubectl delete -f <file>.yaml
```

## Pod Lifecycle

1. **Pending** - Pod accepted, waiting to schedule
2. **Running** - Pod bound to node, containers created
3. **Succeeded** - All containers terminated successfully
4. **Failed** - All containers terminated, at least one failed
5. **Unknown** - Pod state cannot be determined

## Troubleshooting

### Pod stuck in Pending
```bash
# Check events
kubectl describe pod <pod-name>

# Common causes:
# - Insufficient resources
# - Image pull errors
# - Volume issues
```

### Pod in CrashLoopBackOff
```bash
# Check logs
kubectl logs <pod-name>
kubectl logs <pod-name> --previous

# Container is crashing and restarting
# Check application errors in logs
```

### Image Pull Errors
```bash
# ImagePullBackOff or ErrImagePull
# Check image name and tag
# Verify image exists in registry
# Check image pull secrets (for private registries)
```

## Key Takeaways

- Pods are the fundamental unit in Kubernetes
- Containers run inside Pods
- Pods are ephemeral (temporary)
- Use kubectl to interact with Pods
- Namespaces organize resources
- Always check `describe` and `logs` when debugging

## Next Steps

Once comfortable with Pods, move on to:
- **Deployments** (02-deployments/) - Managing multiple Pods
- Learn how to scale and update applications
- Understand self-healing and rolling updates

## Exercises

1. Create a Pod running `httpd:2.4` (Apache web server)
2. Create a namespace called `dev` and deploy a Pod in it
3. Create a multi-container Pod with nginx and busybox
4. Use `kubectl exec` to create a file inside a running Pod
5. Delete a Pod and observe how long it takes to terminate

Happy learning!
