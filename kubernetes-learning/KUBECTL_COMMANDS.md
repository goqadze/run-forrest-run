# kubectl Commands Reference

A comprehensive reference for kubectl, the Kubernetes command-line tool. This guide covers essential commands for managing Kubernetes clusters and resources.

## Table of Contents
- [Core Commands](#core-commands)
- [Pod Management](#pod-management)
- [Deployment Management](#deployment-management)
- [Service & Networking](#service--networking)
- [Configuration & Storage](#configuration--storage)
- [Cluster Management](#cluster-management)
- [Troubleshooting](#troubleshooting)
- [Advanced Commands](#advanced-commands)
- [Quick Tips](#quick-tips)

---

## Core Commands

### `kubectl get`
List resources in the cluster.

**Syntax:**
```bash
kubectl get <resource-type> [name] [flags]
```

**Examples:**
```bash
# List all pods in current namespace
kubectl get pods

# List all pods in all namespaces
kubectl get pods --all-namespaces
kubectl get pods -A

# List specific pod
kubectl get pod nginx-pod

# List multiple resource types
kubectl get pods,services,deployments

# List with more details
kubectl get pods -o wide

# List in YAML format
kubectl get pod nginx-pod -o yaml

# List in JSON format
kubectl get pod nginx-pod -o json

# Watch for changes
kubectl get pods --watch
kubectl get pods -w

# Show labels
kubectl get pods --show-labels

# Filter by label
kubectl get pods -l app=nginx
kubectl get pods --selector=app=nginx
```

**Common Flags:**
- `-o wide` - Show additional details
- `-o yaml` - Output in YAML format
- `-o json` - Output in JSON format
- `-A, --all-namespaces` - All namespaces
- `-n, --namespace` - Specific namespace
- `-w, --watch` - Watch for changes
- `--show-labels` - Show labels
- `-l, --selector` - Filter by labels

---

### `kubectl describe`
Show detailed information about a resource.

**Syntax:**
```bash
kubectl describe <resource-type> <name>
```

**Examples:**
```bash
# Describe a pod
kubectl describe pod nginx-pod

# Describe all pods
kubectl describe pods

# Describe a deployment
kubectl describe deployment my-app

# Describe a node
kubectl describe node minikube

# Describe with namespace
kubectl describe pod nginx-pod -n kube-system
```

**What you'll see:**
- Full resource configuration
- Status and conditions
- Events (very useful for debugging!)
- Related resources

---

### `kubectl create`
Create resources from command line.

**Syntax:**
```bash
kubectl create <resource-type> <name> [flags]
```

**Examples:**
```bash
# Create from file
kubectl create -f pod.yaml

# Create from directory
kubectl create -f ./manifests/

# Create from URL
kubectl create -f https://example.com/manifest.yaml

# Create namespace
kubectl create namespace dev

# Create deployment imperatively
kubectl create deployment nginx --image=nginx:latest

# Create service
kubectl create service clusterip my-service --tcp=80:80

# Create configmap from literal
kubectl create configmap app-config --from-literal=key1=value1

# Create secret
kubectl create secret generic my-secret --from-literal=password=secret123

# Dry run (don't actually create)
kubectl create -f pod.yaml --dry-run=client
```

**Common Flags:**
- `-f, --filename` - File or directory
- `--dry-run=client` - Test without creating
- `-o yaml` - Output YAML (useful with --dry-run)

---

### `kubectl apply`
Apply configuration to resources (declarative management).

**Syntax:**
```bash
kubectl apply -f <file/directory>
```

**Examples:**
```bash
# Apply single file
kubectl apply -f deployment.yaml

# Apply all files in directory
kubectl apply -f ./manifests/

# Apply recursively
kubectl apply -R -f ./

# Apply from URL
kubectl apply -f https://example.com/manifest.yaml

# Apply with server-side validation
kubectl apply -f deployment.yaml --server-side

# Apply and record the command (for rollback)
kubectl apply -f deployment.yaml --record
```

**Difference from create:**
- `create` - Fails if resource exists
- `apply` - Creates or updates resource (idempotent)

**Best Practice:** Use `apply` for declarative management

---

### `kubectl delete`
Delete resources from the cluster.

**Syntax:**
```bash
kubectl delete <resource-type> <name>
kubectl delete -f <file>
```

**Examples:**
```bash
# Delete specific pod
kubectl delete pod nginx-pod

# Delete using file
kubectl delete -f pod.yaml

# Delete all resources in file
kubectl delete -f ./manifests/

# Delete by label
kubectl delete pods -l app=nginx

# Delete all pods (dangerous!)
kubectl delete pods --all

# Delete namespace (and everything in it!)
kubectl delete namespace dev

# Force delete (skip graceful shutdown)
kubectl delete pod nginx-pod --force --grace-period=0

# Delete and wait for completion
kubectl delete pod nginx-pod --wait=true
```

**Common Flags:**
- `-f, --filename` - Delete from file
- `-l, --selector` - Delete by label
- `--all` - Delete all resources of type
- `--force` - Force deletion
- `--grace-period` - Seconds before force kill

---

## Pod Management

### `kubectl exec`
Execute command in a container.

**Syntax:**
```bash
kubectl exec <pod-name> -- <command>
```

**Examples:**
```bash
# Execute single command
kubectl exec nginx-pod -- ls /usr/share/nginx/html

# Interactive shell
kubectl exec -it nginx-pod -- /bin/sh
kubectl exec -it nginx-pod -- /bin/bash

# Specific container in multi-container pod
kubectl exec -it my-pod -c sidecar -- /bin/sh

# Execute in different namespace
kubectl exec -it nginx-pod -n dev -- /bin/sh

# Run as specific user
kubectl exec -it nginx-pod -- su - www-data
```

**Common Flags:**
- `-it` - Interactive terminal
- `-c, --container` - Specific container
- `-n, --namespace` - Specific namespace

---

### `kubectl logs`
View container logs.

**Syntax:**
```bash
kubectl logs <pod-name>
```

**Examples:**
```bash
# View logs
kubectl logs nginx-pod

# Follow logs (like tail -f)
kubectl logs -f nginx-pod
kubectl logs --follow nginx-pod

# Last 100 lines
kubectl logs --tail=100 nginx-pod

# Logs from last hour
kubectl logs --since=1h nginx-pod

# Logs from specific container
kubectl logs nginx-pod -c sidecar

# Previous container instance (if crashed)
kubectl logs -p nginx-pod
kubectl logs --previous nginx-pod

# All containers in pod
kubectl logs nginx-pod --all-containers=true

# Logs with timestamps
kubectl logs nginx-pod --timestamps
```

**Common Flags:**
- `-f, --follow` - Stream logs
- `--tail=N` - Show last N lines
- `--since=TIME` - Since time (1h, 5m, 2023-01-01T00:00:00Z)
- `-p, --previous` - Previous instance
- `-c, --container` - Specific container
- `--timestamps` - Include timestamps

---

### `kubectl port-forward`
Forward local port to pod port.

**Syntax:**
```bash
kubectl port-forward <pod-name> <local-port>:<pod-port>
```

**Examples:**
```bash
# Forward port 8080 to pod's port 80
kubectl port-forward pod/nginx-pod 8080:80

# Access via http://localhost:8080

# Forward to service
kubectl port-forward service/my-service 8080:80

# Forward to deployment
kubectl port-forward deployment/my-app 8080:80

# Listen on all addresses (not just localhost)
kubectl port-forward --address 0.0.0.0 pod/nginx-pod 8080:80

# Specific namespace
kubectl port-forward pod/nginx-pod 8080:80 -n dev
```

**Use cases:**
- Access services without exposing them
- Debugging applications locally
- Temporary access to databases

---

### `kubectl attach`
Attach to a running container.

**Syntax:**
```bash
kubectl attach <pod-name>
```

**Examples:**
```bash
# Attach to pod
kubectl attach nginx-pod

# Attach with stdin
kubectl attach -it nginx-pod

# Specific container
kubectl attach nginx-pod -c sidecar
```

**Difference from exec:**
- `attach` - Connect to main process
- `exec` - Start new process

---

### `kubectl cp`
Copy files to/from containers.

**Syntax:**
```bash
kubectl cp <source> <destination>
```

**Examples:**
```bash
# Copy from pod to local
kubectl cp nginx-pod:/etc/nginx/nginx.conf ./nginx.conf

# Copy from local to pod
kubectl cp ./index.html nginx-pod:/usr/share/nginx/html/

# Specific container
kubectl cp ./file nginx-pod:/tmp/file -c sidecar

# Specific namespace
kubectl cp nginx-pod:/app/log.txt ./log.txt -n dev
```

---

## Deployment Management

### `kubectl rollout`
Manage deployment rollouts.

**Syntax:**
```bash
kubectl rollout <subcommand> <resource>
```

**Examples:**
```bash
# Check rollout status
kubectl rollout status deployment/my-app

# View rollout history
kubectl rollout history deployment/my-app

# View specific revision
kubectl rollout history deployment/my-app --revision=2

# Undo rollout (rollback)
kubectl rollout undo deployment/my-app

# Rollback to specific revision
kubectl rollout undo deployment/my-app --to-revision=2

# Pause rollout
kubectl rollout pause deployment/my-app

# Resume rollout
kubectl rollout resume deployment/my-app

# Restart deployment (rolling restart)
kubectl rollout restart deployment/my-app
```

**Common subcommands:**
- `status` - Show rollout status
- `history` - View revision history
- `undo` - Rollback to previous revision
- `pause` - Pause rollout
- `resume` - Resume rollout
- `restart` - Trigger rolling restart

---

### `kubectl scale`
Scale deployments, replicasets, or statefulsets.

**Syntax:**
```bash
kubectl scale <resource-type>/<name> --replicas=<count>
```

**Examples:**
```bash
# Scale deployment
kubectl scale deployment/my-app --replicas=5

# Scale to zero (stop all pods)
kubectl scale deployment/my-app --replicas=0

# Scale based on current replicas
kubectl scale deployment/my-app --current-replicas=2 --replicas=3

# Scale multiple deployments
kubectl scale deployment/app1 deployment/app2 --replicas=3

# Scale replicaset
kubectl scale replicaset/my-rs --replicas=5
```

---

### `kubectl autoscale`
Create horizontal pod autoscaler.

**Syntax:**
```bash
kubectl autoscale <resource-type> <name> --min=<min> --max=<max> --cpu-percent=<percent>
```

**Examples:**
```bash
# Autoscale based on CPU
kubectl autoscale deployment my-app --min=2 --max=10 --cpu-percent=80

# View autoscalers
kubectl get hpa

# Describe autoscaler
kubectl describe hpa my-app
```

---

## Service & Networking

### `kubectl expose`
Create a service for a resource.

**Syntax:**
```bash
kubectl expose <resource-type> <name> --port=<port>
```

**Examples:**
```bash
# Expose deployment as ClusterIP service
kubectl expose deployment my-app --port=80 --target-port=8080

# Expose as NodePort
kubectl expose deployment my-app --type=NodePort --port=80

# Expose as LoadBalancer
kubectl expose deployment my-app --type=LoadBalancer --port=80

# Expose pod
kubectl expose pod nginx-pod --port=80 --name=nginx-service
```

---

### `kubectl proxy`
Create proxy to Kubernetes API server.

**Syntax:**
```bash
kubectl proxy
```

**Examples:**
```bash
# Start proxy on default port (8001)
kubectl proxy

# Access API at http://localhost:8001/api/v1

# Proxy on specific port
kubectl proxy --port=8080

# Accept connections from any address
kubectl proxy --address='0.0.0.0' --accept-hosts='^*$'
```

---

## Configuration & Storage

### `kubectl config`
Manage kubectl configuration (kubeconfig).

**Syntax:**
```bash
kubectl config <subcommand>
```

**Examples:**
```bash
# View current config
kubectl config view

# Get current context
kubectl config current-context

# List all contexts
kubectl config get-contexts

# Switch context
kubectl config use-context minikube

# Set namespace for context
kubectl config set-context --current --namespace=dev

# Add cluster
kubectl config set-cluster my-cluster --server=https://1.2.3.4

# Add user
kubectl config set-credentials my-user --token=bearer_token

# Create context
kubectl config set-context my-context --cluster=my-cluster --user=my-user

# Delete context
kubectl config delete-context old-context
```

---

### `kubectl label`
Add or update labels on resources.

**Syntax:**
```bash
kubectl label <resource-type> <name> <key>=<value>
```

**Examples:**
```bash
# Add label
kubectl label pod nginx-pod env=production

# Update label (requires --overwrite)
kubectl label pod nginx-pod env=staging --overwrite

# Remove label (use minus sign)
kubectl label pod nginx-pod env-

# Label multiple pods
kubectl label pods --all env=production

# Label by selector
kubectl label pods -l app=nginx tier=frontend
```

---

### `kubectl annotate`
Add or update annotations.

**Syntax:**
```bash
kubectl annotate <resource-type> <name> <key>=<value>
```

**Examples:**
```bash
# Add annotation
kubectl annotate pod nginx-pod description="Production nginx server"

# Update annotation
kubectl annotate pod nginx-pod description="Staging server" --overwrite

# Remove annotation
kubectl annotate pod nginx-pod description-

# Annotate all pods
kubectl annotate pods --all managed-by=kubectl
```

---

## Cluster Management

### `kubectl cluster-info`
Display cluster information.

**Syntax:**
```bash
kubectl cluster-info
```

**Examples:**
```bash
# Basic cluster info
kubectl cluster-info

# Detailed dump
kubectl cluster-info dump

# Dump to directory
kubectl cluster-info dump --output-directory=./cluster-state
```

---

### `kubectl top`
Display resource usage.

**Syntax:**
```bash
kubectl top <node|pod>
```

**Examples:**
```bash
# Node resource usage
kubectl top node

# Pod resource usage
kubectl top pod

# Specific namespace
kubectl top pod -n kube-system

# All namespaces
kubectl top pod --all-namespaces

# Sort by CPU or memory
kubectl top pod --sort-by=cpu
kubectl top pod --sort-by=memory

# Show containers
kubectl top pod --containers
```

**Note:** Requires metrics-server to be installed

---

### `kubectl api-resources`
List available API resources.

**Syntax:**
```bash
kubectl api-resources
```

**Examples:**
```bash
# List all API resources
kubectl api-resources

# Short names only
kubectl api-resources -o name

# Resources in specific API group
kubectl api-resources --api-group=apps

# Namespaced resources
kubectl api-resources --namespaced=true

# Non-namespaced resources
kubectl api-resources --namespaced=false

# Show resource verbs
kubectl api-resources -o wide
```

---

### `kubectl api-versions`
List available API versions.

**Syntax:**
```bash
kubectl api-versions
```

---

## Troubleshooting

### `kubectl debug`
Create debugging session.

**Syntax:**
```bash
kubectl debug <pod-name>
```

**Examples:**
```bash
# Debug pod with ephemeral container
kubectl debug nginx-pod -it --image=busybox

# Debug by creating copy of pod
kubectl debug nginx-pod --copy-to=nginx-debug --container=debug

# Debug node
kubectl debug node/minikube -it --image=ubuntu
```

---

### `kubectl events`
List events.

**Syntax:**
```bash
kubectl events
```

**Examples:**
```bash
# List events
kubectl get events

# Sort by timestamp
kubectl get events --sort-by='.lastTimestamp'

# Watch events
kubectl get events --watch

# Events for specific namespace
kubectl get events -n kube-system

# Events for specific object
kubectl events --for pod/nginx-pod
```

---

### `kubectl diff`
Show differences between live and applied configuration.

**Syntax:**
```bash
kubectl diff -f <file>
```

**Examples:**
```bash
# Show diff before applying
kubectl diff -f deployment.yaml

# Diff for directory
kubectl diff -f ./manifests/
```

---

## Advanced Commands

### `kubectl patch`
Update specific fields of a resource.

**Syntax:**
```bash
kubectl patch <resource-type> <name> -p '<patch>'
```

**Examples:**
```bash
# Patch with strategic merge
kubectl patch deployment my-app -p '{"spec":{"replicas":3}}'

# Patch with JSON
kubectl patch pod nginx-pod -p '{"metadata":{"labels":{"env":"prod"}}}'

# Patch from file
kubectl patch deployment my-app --patch-file=patch.yaml
```

---

### `kubectl edit`
Edit resource in default editor.

**Syntax:**
```bash
kubectl edit <resource-type> <name>
```

**Examples:**
```bash
# Edit deployment
kubectl edit deployment my-app

# Edit service
kubectl edit service my-service

# Use specific editor
KUBE_EDITOR="nano" kubectl edit deployment my-app
```

---

### `kubectl replace`
Replace resource with new definition.

**Syntax:**
```bash
kubectl replace -f <file>
```

**Examples:**
```bash
# Replace from file
kubectl replace -f deployment.yaml

# Force replacement
kubectl replace -f deployment.yaml --force
```

---

### `kubectl wait`
Wait for condition on resources.

**Syntax:**
```bash
kubectl wait --for=<condition> <resource-type>/<name>
```

**Examples:**
```bash
# Wait for pod to be ready
kubectl wait --for=condition=ready pod/nginx-pod

# Wait for deployment
kubectl wait --for=condition=available deployment/my-app

# With timeout
kubectl wait --for=condition=ready pod/nginx-pod --timeout=60s

# Wait for deletion
kubectl wait --for=delete pod/nginx-pod --timeout=30s
```

---

## Quick Tips

### Most Used Commands
```bash
kubectl get pods                  # List pods
kubectl get pods -w              # Watch pods
kubectl describe pod <name>       # Pod details
kubectl logs <pod>               # View logs
kubectl logs -f <pod>            # Stream logs
kubectl exec -it <pod> -- sh     # Shell into pod
kubectl apply -f <file>          # Apply config
kubectl delete -f <file>         # Delete resources
```

### Useful Aliases
Add to your `.bashrc` or `.zshrc`:
```bash
alias k=kubectl
alias kgp='kubectl get pods'
alias kgs='kubectl get services'
alias kgd='kubectl get deployments'
alias kga='kubectl get all'
alias kdp='kubectl describe pod'
alias kl='kubectl logs'
alias kex='kubectl exec -it'
alias kaf='kubectl apply -f'
alias kdf='kubectl delete -f'
```

### Context and Namespace
```bash
# Quick namespace switch
kubectl config set-context --current --namespace=dev

# Or use kubens (if installed)
kubens dev
```

### Output Formats
```bash
kubectl get pods -o yaml        # YAML output
kubectl get pods -o json        # JSON output
kubectl get pods -o wide        # Wide output
kubectl get pods -o name        # Just names
kubectl get pods -o custom-columns=NAME:.metadata.name,STATUS:.status.phase
```

### Debugging Workflow
```bash
# 1. Check pod status
kubectl get pods

# 2. Describe pod for events
kubectl describe pod <pod-name>

# 3. Check logs
kubectl logs <pod-name>

# 4. Exec into pod
kubectl exec -it <pod-name> -- sh

# 5. Check events
kubectl get events --sort-by='.lastTimestamp'
```

### Resource Cleanup
```bash
# Delete all pods (current namespace)
kubectl delete pods --all

# Delete all resources
kubectl delete all --all

# Delete specific label
kubectl delete pods -l app=nginx

# Force delete stuck pod
kubectl delete pod <pod-name> --force --grace-period=0
```

### Getting Help
```bash
# General help
kubectl --help
kubectl -h

# Command help
kubectl get --help

# Explain resource
kubectl explain pod
kubectl explain pod.spec
kubectl explain pod.spec.containers

# API reference
kubectl api-resources
kubectl api-versions
```

---

## See Also

- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [README.md](README.md) - Project learning guide
- [SETUP.md](SETUP.md) - Cluster setup guide
