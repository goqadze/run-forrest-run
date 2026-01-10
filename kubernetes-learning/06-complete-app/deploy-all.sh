#!/bin/bash
# Deploy complete application

set -e

echo "=========================================="
echo "Deploying Complete Application"
echo "=========================================="
echo ""

# 1. Create namespace
echo "1. Creating namespace..."
kubectl apply -f namespace.yaml
echo ""

# 2. Create ConfigMap and Secret
echo "2. Creating configuration..."
kubectl apply -f postgres-configmap.yaml
kubectl apply -f postgres-secret.yaml
echo ""

# 3. Create storage
echo "3. Creating persistent storage..."
kubectl apply -f postgres-pv.yaml
kubectl apply -f postgres-pvc.yaml
echo ""

# Wait for PVC to be bound
echo "Waiting for PVC to be bound..."
kubectl wait --for=jsonpath='{.status.phase}'=Bound pvc/postgres-pvc -n webapp --timeout=60s
echo ""

# 4. Deploy database
echo "4. Deploying PostgreSQL database..."
kubectl apply -f postgres-deployment.yaml
kubectl apply -f postgres-service.yaml
echo ""

# Wait for database to be ready
echo "Waiting for database to be ready..."
kubectl wait --for=condition=ready pod -l app=postgres -n webapp --timeout=120s
echo ""

# 5. Deploy web application
echo "5. Deploying web application..."
kubectl apply -f webapp-deployment.yaml
kubectl apply -f webapp-service.yaml
echo ""

# Wait for webapp to be ready
echo "Waiting for web app to be ready..."
kubectl wait --for=condition=ready pod -l app=webapp -n webapp --timeout=120s
echo ""

# 6. Display status
echo "=========================================="
echo "Deployment Complete!"
echo "=========================================="
echo ""
echo "Checking deployment status..."
kubectl get all -n webapp
echo ""

echo "ConfigMaps and Secrets:"
kubectl get configmap,secret -n webapp
echo ""

echo "Storage:"
kubectl get pv,pvc -n webapp
echo ""

echo "=========================================="
echo "Access the application:"
echo "=========================================="
echo ""
echo "Option 1 - Minikube service:"
echo "  minikube service webapp-service -n webapp"
echo ""
echo "Option 2 - Port forward:"
echo "  kubectl port-forward -n webapp service/webapp-service 8080:80"
echo "  Then visit: http://localhost:8080"
echo ""
echo "Option 3 - Get NodePort URL:"
echo "  minikube service webapp-service -n webapp --url"
echo ""

echo "To view logs:"
echo "  kubectl logs -l app=postgres -n webapp"
echo "  kubectl logs -l app=webapp -n webapp"
echo ""

echo "To cleanup:"
echo "  kubectl delete namespace webapp"
echo "  kubectl delete pv postgres-pv"
echo ""
