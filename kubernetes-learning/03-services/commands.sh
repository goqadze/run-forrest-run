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
