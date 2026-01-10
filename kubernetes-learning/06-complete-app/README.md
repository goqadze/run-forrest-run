# Complete Application Deployment

This example demonstrates deploying a complete 2-tier application combining all concepts learned:
- PostgreSQL database with persistent storage
- Web application frontend
- Services for networking
- ConfigMaps and Secrets for configuration
- Namespaces for isolation

## Application Architecture

```
┌─────────────────────────────────────┐
│         Namespace: webapp           │
│                                     │
│  ┌──────────────┐  ┌─────────────┐ │
│  │   Web App    │  │  PostgreSQL │ │
│  │  Deployment  │──│  Deployment │ │
│  │  (3 replicas)│  │  (1 replica)│ │
│  └──────┬───────┘  └──────┬──────┘ │
│         │                 │         │
│  ┌──────▼───────┐  ┌──────▼──────┐ │
│  │  Web Service │  │ DB Service  │ │
│  │  (NodePort)  │  │ (ClusterIP) │ │
│  └──────────────┘  └──────┬──────┘ │
│                           │         │
│                    ┌──────▼──────┐  │
│                    │ PersistentVol│  │
│                    │    Claim     │  │
│                    └──────────────┘  │
└─────────────────────────────────────┘
```

## What Gets Deployed

1. **Namespace** - `webapp` namespace for isolation
2. **PostgreSQL Database**
   - Deployment with 1 replica
   - Persistent storage for data
   - ConfigMap for database configuration
   - Secret for database credentials
   - Internal ClusterIP service
3. **Web Application**
   - Deployment with 3 replicas
   - Connects to PostgreSQL
   - NodePort service for external access
4. **All supporting resources**

## Files

- `namespace.yaml` - Creates webapp namespace
- `postgres-configmap.yaml` - Database configuration
- `postgres-secret.yaml` - Database credentials
- `postgres-pv.yaml` - PersistentVolume for database
- `postgres-pvc.yaml` - PersistentVolumeClaim
- `postgres-deployment.yaml` - Database deployment
- `postgres-service.yaml` - Database service
- `webapp-deployment.yaml` - Web app deployment
- `webapp-service.yaml` - Web app service
- `deploy-all.sh` - Script to deploy everything

## Deployment Order

1. Namespace
2. ConfigMap and Secret
3. PersistentVolume and PersistentVolumeClaim
4. Database deployment and service
5. Web app deployment and service

## Quick Start

```bash
# Deploy everything
./deploy-all.sh

# Check deployment status
kubectl get all -n webapp

# Access the application
minikube service webapp-service -n webapp
# Or get the URL
minikube service webapp-service -n webapp --url
```

## Manual Deployment

```bash
# 1. Create namespace
kubectl apply -f namespace.yaml

# 2. Create configuration
kubectl apply -f postgres-configmap.yaml
kubectl apply -f postgres-secret.yaml

# 3. Create storage
kubectl apply -f postgres-pv.yaml
kubectl apply -f postgres-pvc.yaml

# 4. Deploy database
kubectl apply -f postgres-deployment.yaml
kubectl apply -f postgres-service.yaml

# Wait for database to be ready
kubectl wait --for=condition=ready pod -l app=postgres -n webapp --timeout=60s

# 5. Deploy web application
kubectl apply -f webapp-deployment.yaml
kubectl apply -f webapp-service.yaml

# 6. Check everything
kubectl get all -n webapp
```

## Verification

```bash
# Check namespace
kubectl get ns webapp

# Check all resources in namespace
kubectl get all -n webapp

# Check ConfigMaps and Secrets
kubectl get configmap -n webapp
kubectl get secret -n webapp

# Check persistent volumes
kubectl get pv
kubectl get pvc -n webapp

# Describe deployments
kubectl describe deployment postgres -n webapp
kubectl describe deployment webapp -n webapp

# View logs
kubectl logs -l app=postgres -n webapp
kubectl logs -l app=webapp -n webapp

# Check service endpoints
kubectl get endpoints -n webapp
```

## Accessing the Application

```bash
# Using minikube
minikube service webapp-service -n webapp

# Or get the URL
URL=$(minikube service webapp-service -n webapp --url)
curl $URL

# Port forward (alternative)
kubectl port-forward -n webapp service/webapp-service 8080:80
# Then access http://localhost:8080
```

## Troubleshooting

### Database not ready
```bash
# Check Pod status
kubectl get pods -n webapp

# View logs
kubectl logs -l app=postgres -n webapp

# Describe Pod for events
kubectl describe pod -l app=postgres -n webapp

# Check PVC is bound
kubectl get pvc -n webapp
```

### Web app can't connect to database
```bash
# Verify database service exists
kubectl get service postgres-service -n webapp

# Check service endpoints
kubectl get endpoints postgres-service -n webapp

# Test connection from web app Pod
kubectl exec -it -n webapp deployment/webapp -- env | grep POSTGRES
```

### Can't access web app
```bash
# Check service
kubectl get service webapp-service -n webapp

# For minikube, ensure tunnel is running
minikube tunnel

# Or use NodePort
kubectl get service webapp-service -n webapp -o wide
```

## Cleanup

```bash
# Delete everything
kubectl delete namespace webapp

# This deletes:
# - All Pods
# - All Deployments
# - All Services
# - All ConfigMaps and Secrets
# - All PVCs

# Note: PersistentVolumes may need manual deletion
kubectl delete pv postgres-pv
```

## Key Takeaways

- Real applications require multiple resources working together
- Order matters when deploying (dependencies first)
- Use namespaces to organize related resources
- Always separate configuration from code
- Persistent storage ensures data survives Pod restarts
- Services enable communication between components
- Labels and selectors tie everything together

## Exercises

1. Scale the web app to 5 replicas
2. Add a Redis cache to the architecture
3. Change the web app service to LoadBalancer type
4. Add resource limits to both deployments
5. Create an Ingress resource to expose the app
6. Add liveness and readiness probes to both deployments
7. Implement a rolling update of the web app
8. Add a second database replica for high availability

## Next Steps

Congratulations on completing the Kubernetes learning project!

Continue your journey:
- Deploy to a cloud provider (EKS, GKE, AKS)
- Learn about Ingress controllers
- Explore Helm for package management
- Study StatefulSets for stateful applications
- Implement monitoring with Prometheus/Grafana
- Learn about service meshes (Istio, Linkerd)
- Practice for CKA/CKAD certification

You now have the foundational knowledge to work with Kubernetes!
