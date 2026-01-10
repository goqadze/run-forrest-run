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
