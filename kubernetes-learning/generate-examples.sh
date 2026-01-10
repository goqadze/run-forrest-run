#!/bin/bash
# Script to generate all remaining example files for the Kubernetes learning project

set -e  # Exit on error

BASE_DIR="/Users/avtandilgokadze/Learning/kubernetes-learning"

echo "Generating Kubernetes learning project files..."

# 02-deployments
echo "Creating 02-deployments files..."

cat > "$BASE_DIR/02-deployments/README.md" << 'EOF'
# Deployments

Deployments manage the lifecycle of your applications in Kubernetes. They provide declarative updates, scaling, and self-healing capabilities.

## What is a Deployment?

A Deployment:
- Manages ReplicaSets
- Ensures desired number of Pods are running
- Handles rolling updates with zero downtime
- Enables rollback to previous versions
- Provides self-healing (restarts failed Pods)

## Key Concepts

**ReplicaSet:** Maintains a stable set of replica Pods
**Rolling Update:** Gradually replace old Pods with new ones
**Rollback:** Revert to a previous version
**Scaling:** Increase or decrease replicas

## Files

1. **deployment.yaml** - Basic deployment with 3 replicas
2. **deployment-with-resources.yaml** - Deployment with resource limits
3. **scaling.yaml** - Example showing scaling
4. **commands.sh** - kubectl commands to try

## Common Operations

```bash
# Create deployment
kubectl apply -f deployment.yaml

# View deployments
kubectl get deployments
kubectl get deploy

# View ReplicaSets (created by deployment)
kubectl get replicasets
kubectl get rs

# View Pods (created by ReplicaSet)
kubectl get pods

# Scale deployment
kubectl scale deployment nginx-deployment --replicas=5

# Update image (rolling update)
kubectl set image deployment/nginx-deployment nginx=nginx:1.21

# Check rollout status
kubectl rollout status deployment/nginx-deployment

# View rollout history
kubectl rollout history deployment/nginx-deployment

# Rollback
kubectl rollout undo deployment/nginx-deployment

# Delete deployment
kubectl delete deployment nginx-deployment
```

## Next: Services (03-services/)
EOF

cat > "$BASE_DIR/02-deployments/deployment.yaml" << 'EOF'
# Basic Deployment - manages multiple Pod replicas

apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3                    # Number of Pod replicas
  selector:                      # How Deployment finds Pods to manage
    matchLabels:
      app: nginx
  template:                      # Pod template
    metadata:
      labels:
        app: nginx               # Must match selector
    spec:
      containers:
      - name: nginx
        image: nginx:1.20
        ports:
        - containerPort: 80
EOF

cat > "$BASE_DIR/02-deployments/deployment-with-resources.yaml" << 'EOF'
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment-resources
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx-resources
  template:
    metadata:
      labels:
        app: nginx-resources
    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80
        resources:
          requests:              # Minimum guaranteed resources
            memory: "64Mi"
            cpu: "250m"
          limits:                # Maximum allowed resources
            memory: "128Mi"
            cpu: "500m"
        livenessProbe:           # Check if container is alive
          httpGet:
            path: /
            port: 80
          initialDelaySeconds: 3
          periodSeconds: 3
        readinessProbe:          # Check if container is ready for traffic
          httpGet:
            path: /
            port: 80
          initialDelaySeconds: 3
          periodSeconds: 3
EOF

cat > "$BASE_DIR/02-deployments/commands.sh" << 'EOF'
#!/bin/bash

echo "=== Deployments Practice ==="

echo "1. Creating deployment..."
kubectl apply -f deployment.yaml

echo "2. Viewing deployments..."
kubectl get deployments

echo "3. Viewing replica sets..."
kubectl get rs

echo "4. Viewing pods..."
kubectl get pods

echo "5. Scaling up to 5 replicas..."
kubectl scale deployment nginx-deployment --replicas=5
kubectl get pods

echo "6. Scaling down to 2 replicas..."
kubectl scale deployment nginx-deployment --replicas=2
kubectl get pods

echo "7. Updating image (rolling update)..."
kubectl set image deployment/nginx-deployment nginx=nginx:1.21
kubectl rollout status deployment/nginx-deployment

echo "8. Viewing rollout history..."
kubectl rollout history deployment/nginx-deployment

echo "=== Cleanup ==="
echo "kubectl delete -f deployment.yaml"
EOF

chmod +x "$BASE_DIR/02-deployments/commands.sh"

# 03-services
echo "Creating 03-services files..."

cat > "$BASE_DIR/03-services/README.md" << 'EOF'
# Services

Services provide stable networking for Pods. They enable load balancing and service discovery.

## Service Types

1. **ClusterIP** (default) - Internal access only
2. **NodePort** - Exposes service on each node's IP
3. **LoadBalancer** - Cloud provider's load balancer
4. **ExternalName** - Maps service to DNS name

## Files

1. **clusterip-service.yaml** - Internal service
2. **nodeport-service.yaml** - External access via node port
3. **loadbalancer-service.yaml** - Load balancer service
4. **commands.sh** - kubectl commands

## Common Commands

```bash
kubectl apply -f clusterip-service.yaml
kubectl get services
kubectl get svc
kubectl describe service my-service
```

## Next: ConfigMaps & Secrets (04-config-secrets/)
EOF

cat > "$BASE_DIR/03-services/clusterip-service.yaml" << 'EOF'
apiVersion: v1
kind: Service
metadata:
  name: nginx-clusterip
spec:
  type: ClusterIP          # Default type - internal access only
  selector:
    app: nginx             # Selects Pods with this label
  ports:
  - port: 80               # Service port
    targetPort: 80         # Container port
    protocol: TCP
EOF

cat > "$BASE_DIR/03-services/nodeport-service.yaml" << 'EOF'
apiVersion: v1
kind: Service
metadata:
  name: nginx-nodeport
spec:
  type: NodePort
  selector:
    app: nginx
  ports:
  - port: 80
    targetPort: 80
    nodePort: 30080        # External port (30000-32767)
EOF

cat > "$BASE_DIR/03-services/commands.sh" << 'EOF'
#!/bin/bash

echo "=== Services Practice ==="

echo "First, create a deployment..."
kubectl create deployment nginx --image=nginx --replicas=3

echo "1. Creating ClusterIP service..."
kubectl apply -f clusterip-service.yaml
kubectl get svc

echo "2. Creating NodePort service..."
kubectl apply -f nodeport-service.yaml
kubectl get svc

echo "3. Accessing service (minikube)..."
echo "Run: minikube service nginx-nodeport"

echo "=== Cleanup ==="
echo "kubectl delete deployment nginx"
echo "kubectl delete -f ."
EOF

chmod +x "$BASE_DIR/03-services/commands.sh"

# 04-config-secrets
echo "Creating 04-config-secrets files..."

cat > "$BASE_DIR/04-config-secrets/README.md" << 'EOF'
# ConfigMaps and Secrets

Separate configuration from application code using ConfigMaps and Secrets.

## ConfigMaps
- Store non-sensitive configuration
- Key-value pairs or entire files
- Can be used as environment variables or volumes

## Secrets
- Store sensitive data (passwords, tokens, keys)
- Base64 encoded (NOT encrypted!)
- Should use encryption at rest in production

## Files

1. **configmap.yaml** - Basic ConfigMap
2. **secret.yaml** - Basic Secret
3. **pod-with-config.yaml** - Pod using ConfigMap and Secret
4. **commands.sh** - kubectl commands

## Next: Storage (05-storage/)
EOF

cat > "$BASE_DIR/04-config-secrets/configmap.yaml" << 'EOF'
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  app.properties: |
    environment=development
    debug=true
    max_connections=100
  log_level: "info"
  database_url: "postgres://localhost:5432/mydb"
EOF

cat > "$BASE_DIR/04-config-secrets/secret.yaml" << 'EOF'
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
type: Opaque
stringData:                 # Automatically base64 encoded
  username: admin
  password: supersecret123
# Or use data with base64 encoded values:
# data:
#   username: YWRtaW4=     # echo -n 'admin' | base64
#   password: c3VwZXJzZWNyZXQxMjM=
EOF

cat > "$BASE_DIR/04-config-secrets/pod-with-config.yaml" << 'EOF'
apiVersion: v1
kind: Pod
metadata:
  name: pod-with-config
spec:
  containers:
  - name: app
    image: busybox
    command: ["sh", "-c", "echo Username: $USERNAME && sleep 3600"]
    env:
    - name: USERNAME
      valueFrom:
        secretKeyRef:
          name: app-secret
          key: username
    - name: LOG_LEVEL
      valueFrom:
        configMapKeyRef:
          name: app-config
          key: log_level
    volumeMounts:
    - name: config-volume
      mountPath: /etc/config
  volumes:
  - name: config-volume
    configMap:
      name: app-config
EOF

cat > "$BASE_DIR/04-config-secrets/commands.sh" << 'EOF'
#!/bin/bash

echo "=== ConfigMaps and Secrets Practice ==="

echo "1. Creating ConfigMap..."
kubectl apply -f configmap.yaml

echo "2. Creating Secret..."
kubectl apply -f secret.yaml

echo "3. Viewing ConfigMap..."
kubectl get configmap app-config -o yaml

echo "4. Creating Pod with config..."
kubectl apply -f pod-with-config.yaml

echo "5. Checking Pod logs..."
sleep 3
kubectl logs pod-with-config

echo "6. Exec into Pod to see mounted config..."
kubectl exec pod-with-config -- ls -la /etc/config
kubectl exec pod-with-config -- cat /etc/config/app.properties

echo "=== Cleanup ==="
echo "kubectl delete -f ."
EOF

chmod +x "$BASE_DIR/04-config-secrets/commands.sh"

# 05-storage
echo "Creating 05-storage files..."

cat > "$BASE_DIR/05-storage/README.md" << 'EOF'
# Storage

Persistent storage in Kubernetes using Volumes, PersistentVolumes, and PersistentVolumeClaims.

## Volume Types

**emptyDir** - Temporary storage (deleted with Pod)
**hostPath** - Mount from node's filesystem
**PersistentVolume (PV)** - Cluster-wide storage resource
**PersistentVolumeClaim (PVC)** - User request for storage

## Files

1. **emptydir-volume.yaml** - Temporary storage
2. **persistent-volume.yaml** - PV definition
3. **persistent-volume-claim.yaml** - PVC definition
4. **pod-with-pvc.yaml** - Pod using PVC
5. **commands.sh** - kubectl commands

## Next: Complete App (06-complete-app/)
EOF

cat > "$BASE_DIR/05-storage/emptydir-volume.yaml" << 'EOF'
apiVersion: v1
kind: Pod
metadata:
  name: pod-with-emptydir
spec:
  containers:
  - name: writer
    image: busybox
    command: ["sh", "-c", "echo Hello > /data/hello.txt && sleep 3600"]
    volumeMounts:
    - name: shared-data
      mountPath: /data
  - name: reader
    image: busybox
    command: ["sh", "-c", "sleep 5 && cat /data/hello.txt && sleep 3600"]
    volumeMounts:
    - name: shared-data
      mountPath: /data
  volumes:
  - name: shared-data
    emptyDir: {}
EOF

cat > "$BASE_DIR/05-storage/persistent-volume.yaml" << 'EOF'
apiVersion: v1
kind: PersistentVolume
metadata:
  name: my-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /tmp/data
EOF

cat > "$BASE_DIR/05-storage/persistent-volume-claim.yaml" << 'EOF'
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 500Mi
EOF

cat > "$BASE_DIR/05-storage/commands.sh" << 'EOF'
#!/bin/bash

echo "=== Storage Practice ==="

echo "1. Creating Pod with emptyDir..."
kubectl apply -f emptydir-volume.yaml
sleep 5
kubectl logs pod-with-emptydir -c reader

echo "2. Creating PersistentVolume..."
kubectl apply -f persistent-volume.yaml
kubectl get pv

echo "3. Creating PersistentVolumeClaim..."
kubectl apply -f persistent-volume-claim.yaml
kubectl get pvc

echo "=== Cleanup ==="
echo "kubectl delete -f ."
EOF

chmod +x "$BASE_DIR/05-storage/commands.sh"

# exercises
echo "Creating exercises..."

cat > "$BASE_DIR/exercises/README.md" << 'EOF'
# Practice Exercises

Test your Kubernetes skills with these challenges!

## Beginner Exercises

1. **Create a Pod**
   - Deploy a Pod running `httpd:2.4`
   - Name it `apache-pod`
   - Verify it's running

2. **Use a Namespace**
   - Create a namespace called `practice`
   - Deploy nginx in that namespace
   - List all Pods in the namespace

3. **Create a Deployment**
   - Create a deployment with 3 replicas of nginx
   - Scale it to 5 replicas
   - Scale it back to 2

4. **Expose with Service**
   - Create a deployment
   - Expose it with a ClusterIP service
   - Expose it with a NodePort service

## Intermediate Exercises

5. **ConfigMap Usage**
   - Create a ConfigMap with database settings
   - Create a Pod that reads those settings as env vars

6. **Secrets**
   - Create a Secret with username/password
   - Mount it as a volume in a Pod

7. **Rolling Update**
   - Create a deployment with nginx:1.20
   - Update to nginx:1.21
   - Check rollout status
   - Rollback to previous version

8. **Resource Limits**
   - Create a deployment with CPU and memory limits
   - Set requests and limits appropriately

## Advanced Exercises

9. **Multi-Tier Application**
   - Deploy a database (postgres)
   - Deploy a web application
   - Connect them with services
   - Use ConfigMaps for database connection info

10. **Troubleshooting**
    - Deploy an intentionally broken Pod
    - Use describe and logs to find the issue
    - Fix the problem

## Project: Deploy a Real Application

Deploy a complete 2-tier application:
- PostgreSQL database with persistent storage
- Node.js/Python web app connecting to database
- Service to expose the web app
- ConfigMaps for configuration
- Secrets for database credentials

Good luck!
EOF

echo ""
echo "✅ All files generated successfully!"
echo ""
echo "Project structure:"
tree -L 2 "$BASE_DIR" 2>/dev/null || ls -R "$BASE_DIR"

echo ""
echo "Next steps:"
echo "1. cd $BASE_DIR"
echo "2. Read README.md"
echo "3. Follow SETUP.md to start a cluster"
echo "4. Work through 01-basics/ to 06-complete-app/"
echo ""
echo "Happy learning!"
EOF
