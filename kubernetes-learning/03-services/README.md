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
