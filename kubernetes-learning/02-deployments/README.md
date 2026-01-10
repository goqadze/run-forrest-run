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
