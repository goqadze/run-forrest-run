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
