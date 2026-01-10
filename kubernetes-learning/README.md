# Kubernetes Learning Project

A hands-on project to learn Kubernetes through practical examples, from basic concepts to deploying complete applications.

## Prerequisites

Before starting, you should have:
- **Docker** installed and running
- **kubectl** command-line tool installed
- **A local Kubernetes cluster** (minikube, kind, or Docker Desktop)

Don't have these? See [SETUP.md](SETUP.md) for installation instructions.

## Project Structure

```
kubernetes-learning/
├── README.md                       # You are here!
├── KUBECTL_COMMANDS.md             # kubectl command reference
├── SETUP.md                        # Local cluster setup guide
├── 01-basics/                      # Pods, containers, namespaces
│   ├── README.md
│   ├── simple-pod.yaml
│   ├── multi-container-pod.yaml
│   ├── namespace.yaml
│   └── commands.sh
├── 02-deployments/                 # Deployments, scaling, updates
│   ├── README.md
│   ├── deployment.yaml
│   ├── deployment-with-resources.yaml
│   ├── scaling.yaml
│   └── commands.sh
├── 03-services/                    # Services and networking
│   ├── README.md
│   ├── clusterip-service.yaml
│   ├── nodeport-service.yaml
│   ├── loadbalancer-service.yaml
│   └── commands.sh
├── 04-config-secrets/              # Configuration management
│   ├── README.md
│   ├── configmap.yaml
│   ├── configmap-env.yaml
│   ├── secret.yaml
│   ├── pod-with-config.yaml
│   └── commands.sh
├── 05-storage/                     # Persistent storage
│   ├── README.md
│   ├── emptydir-volume.yaml
│   ├── persistent-volume.yaml
│   ├── persistent-volume-claim.yaml
│   ├── pod-with-pvc.yaml
│   └── commands.sh
├── 06-complete-app/                # Full application deployment
│   ├── README.md
│   ├── All manifests for 2-tier app
│   └── deploy-all.sh
└── exercises/                      # Practice challenges
    └── README.md
```

## Getting Started

### 1. Set Up a Local Cluster

Follow the [SETUP.md](SETUP.md) guide to install and start a local Kubernetes cluster.

### 2. Verify Your Setup

```bash
# Check kubectl is installed
kubectl version --client

# Check cluster is running
kubectl cluster-info

# View nodes
kubectl get nodes
```

You should see your local cluster node in a "Ready" state.

### 3. Start Learning

Work through the topics in order:

```bash
# Start with basics
cd 01-basics
cat README.md
```

Each topic directory contains:
- **README.md** - Explains the concepts
- **YAML files** - Kubernetes manifests with detailed comments
- **commands.sh** - kubectl commands to try

## Topics Covered

### 1. Basics (`01-basics/`)
**What you'll learn:**
- What is a Pod and why it matters
- How containers run inside Pods
- Creating and managing Pods
- Namespaces for organizing resources
- Basic kubectl commands

**Key concepts:**
- Pods are the smallest deployable units in Kubernetes
- Pods can contain one or more containers
- Namespaces provide logical isolation
- kubectl is your primary tool for interacting with clusters

**Files:** `simple-pod.yaml`, `multi-container-pod.yaml`, `namespace.yaml`

### 2. Deployments (`02-deployments/`)
**What you'll learn:**
- Managing application lifecycle with Deployments
- Scaling applications up and down
- Rolling updates and rollbacks
- Resource limits and requests
- ReplicaSets (managed by Deployments)

**Key concepts:**
- Deployments manage ReplicaSets
- ReplicaSets ensure desired number of Pods are running
- Scaling is declarative (you declare desired state)
- Rolling updates happen automatically with zero downtime
- Resource limits prevent resource starvation

**Files:** `deployment.yaml`, `deployment-with-resources.yaml`, `scaling.yaml`

### 3. Services (`03-services/`)
**What you'll learn:**
- Exposing applications with Services
- Service types (ClusterIP, NodePort, LoadBalancer)
- Service discovery and DNS
- Load balancing across Pods

**Key concepts:**
- Services provide stable networking for Pods
- ClusterIP: Internal access only (default)
- NodePort: External access via node's IP:port
- LoadBalancer: Cloud provider's load balancer
- Services use selectors to find Pods

**Files:** `clusterip-service.yaml`, `nodeport-service.yaml`, `loadbalancer-service.yaml`

### 4. ConfigMaps & Secrets (`04-config-secrets/`)
**What you'll learn:**
- Separating configuration from application code
- Using ConfigMaps for non-sensitive data
- Using Secrets for sensitive data
- Injecting configuration as environment variables
- Mounting configuration as files

**Key concepts:**
- Don't hardcode configuration in container images
- ConfigMaps store key-value pairs or files
- Secrets are base64 encoded (not encrypted!)
- Configuration can be env vars or volume mounts
- Changes require Pod restart to take effect

**Files:** `configmap.yaml`, `secret.yaml`, `pod-with-config.yaml`

### 5. Storage (`05-storage/`)
**What you'll learn:**
- Ephemeral vs persistent storage
- Volume types (emptyDir, hostPath, PV/PVC)
- Persistent Volumes and Claims
- Storage classes
- StatefulSets basics

**Key concepts:**
- Container filesystems are ephemeral by default
- Volumes persist data beyond container restarts
- PersistentVolumes are cluster resources
- PersistentVolumeClaims are user requests for storage
- Dynamic provisioning with StorageClasses

**Files:** `emptydir-volume.yaml`, `persistent-volume.yaml`, `persistent-volume-claim.yaml`

### 6. Complete Application (`06-complete-app/`)
**What you'll learn:**
- Deploying a real 2-tier application
- Database with persistent storage
- Web frontend with multiple replicas
- Combining all concepts learned
- Best practices for production

**Demonstrates:**
- PostgreSQL database deployment
- Web application deployment
- Service networking between tiers
- Configuration management
- Persistent storage
- Complete deployment automation

## Learning Path

**Recommended progression:**

1. **Start with Setup** - Get your local cluster running (SETUP.md)
2. **Learn Basics** - Understand Pods and namespaces (01-basics/)
3. **Master Deployments** - Learn application management (02-deployments/)
4. **Understand Networking** - Expose applications with Services (03-services/)
5. **Manage Configuration** - Separate config from code (04-config-secrets/)
6. **Handle Storage** - Persist data across restarts (05-storage/)
7. **Deploy Complete App** - Put it all together (06-complete-app/)
8. **Practice** - Try the exercises (exercises/)

**Each topic builds on previous ones**, so follow the order for best results.

## How to Use This Project

### For Each Topic:

1. **Read the README** in the topic directory
2. **Examine the YAML files** - they're heavily commented
3. **Run the commands** in `commands.sh` (or run them manually)
4. **Experiment** - modify values and see what happens
5. **Clean up** - delete resources before moving to next topic

### Example Workflow:

```bash
cd 01-basics

# Read the guide
cat README.md

# Look at the YAML
cat simple-pod.yaml

# Apply the manifest
kubectl apply -f simple-pod.yaml

# Check the result
kubectl get pods
kubectl describe pod nginx-pod

# Clean up
kubectl delete -f simple-pod.yaml
```

## Common Kubernetes Patterns

### Declarative Configuration
```bash
# Apply configuration from file
kubectl apply -f deployment.yaml

# Apply all files in directory
kubectl apply -f 01-basics/

# Delete using the same file
kubectl delete -f deployment.yaml
```

### Checking Resources
```bash
# List resources
kubectl get pods
kubectl get deployments
kubectl get services
kubectl get all

# Detailed information
kubectl describe pod <pod-name>

# Watch resources update in real-time
kubectl get pods --watch
```

### Troubleshooting
```bash
# View logs
kubectl logs <pod-name>

# Exec into a container
kubectl exec -it <pod-name> -- /bin/sh

# Port forward to local machine
kubectl port-forward pod/<pod-name> 8080:80

# Check events
kubectl get events --sort-by='.lastTimestamp'
```

## Resources

- **[kubectl Commands Reference](KUBECTL_COMMANDS.md)** - Complete kubectl command guide
- **[Setup Guide](SETUP.md)** - Local cluster installation
- Official Kubernetes Documentation: https://kubernetes.io/docs/
- Kubernetes Basics Tutorial: https://kubernetes.io/docs/tutorials/kubernetes-basics/
- kubectl Cheat Sheet: https://kubernetes.io/docs/reference/kubectl/cheatsheet/
- Interactive Tutorial: https://kubernetes.io/docs/tutorials/
- YAML Tutorial: https://learnxinyminutes.com/docs/yaml/

## Tips for Learning

- **Start simple** - Don't skip the basics
- **Read the YAML** - Understanding manifest structure is crucial
- **Use describe and logs** - These are your debugging friends
- **Experiment safely** - Local clusters are for breaking things
- **Clean up resources** - Avoid cluttering your cluster
- **Check official docs** - They're comprehensive and well-written
- **Practice regularly** - Kubernetes has many concepts; repetition helps
- **Use labels** - They help organize and select resources
- **Understand the declarative model** - You describe desired state, Kubernetes makes it happen

## Common Gotchas

1. **YAML indentation** - Must use spaces, not tabs
2. **Image pull errors** - Check image name and availability
3. **Resource not found** - Check namespace with `-n <namespace>`
4. **Pending Pods** - Usually resource constraints or scheduling issues
5. **Service can't connect** - Check selectors match Pod labels
6. **Changes not applying** - Remember to `kubectl apply` after editing
7. **Permission errors** - Check RBAC and service account permissions

## Exercises to Try

1. Create a Pod with your favorite application (not nginx!)
2. Deploy an application with 5 replicas and scale it to 10
3. Set up a service to access your deployment
4. Add configuration using ConfigMaps
5. Deploy a stateful application with persistent storage
6. Create a multi-tier application from scratch
7. Implement a rolling update with zero downtime
8. Debug a failing Pod using logs and describe
9. Set up resource limits to prevent resource hogging
10. Deploy the complete app and customize it

See [exercises/README.md](exercises/README.md) for detailed challenges.

## Next Steps

After completing this project, you'll be ready to:

- **Deploy to production clusters** - AWS EKS, Google GKE, Azure AKS
- **Learn advanced topics** - Ingress controllers, network policies, RBAC
- **Explore observability** - Prometheus, Grafana, logging solutions
- **Study GitOps** - ArgoCD, Flux for declarative deployments
- **Dive into Helm** - Package manager for Kubernetes
- **Understand operators** - Extend Kubernetes with custom resources
- **Master security** - Pod security policies, network policies, secrets management
- **Learn service mesh** - Istio, Linkerd for advanced networking
- **Practice CKA/CKAD** - Get certified!

## Troubleshooting This Project

### Cluster not starting
- Check Docker is running
- Try `minikube delete` and `minikube start` again
- Check system resources (CPU, memory)

### kubectl commands not working
- Verify cluster is running: `kubectl cluster-info`
- Check context: `kubectl config current-context`
- Verify kubectl is installed: `kubectl version`

### Pods staying in Pending
- Check node resources: `kubectl describe nodes`
- Look at Pod events: `kubectl describe pod <pod-name>`
- For PVCs, check if storage class exists

### Can't access services
- For LoadBalancer on minikube: use `minikube tunnel`
- For NodePort: use `minikube service <service-name>`
- Check firewall rules

## Contributing & Feedback

This is a learning project. If you find errors or have suggestions:
- Practice fixing it yourself first (great learning!)
- Refer to official Kubernetes documentation
- Check kubectl command reference in this repo

Happy learning with Kubernetes!
