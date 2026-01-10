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
