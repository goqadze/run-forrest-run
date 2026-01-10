#!/bin/bash
# Kubernetes Basics - Practice Commands
# Run these commands to learn Pod and Namespace management

echo "=== Kubernetes Basics - Hands-On Practice ==="
echo ""

# 1. View available namespaces
echo "1. Viewing namespaces..."
kubectl get namespaces
echo ""

# 2. Create simple Pod
echo "2. Creating simple Pod..."
kubectl apply -f simple-pod.yaml
echo ""

# 3. Check Pod status
echo "3. Checking Pod status..."
kubectl get pods
echo ""
sleep 3

# 4. Get detailed Pod information
echo "4. Describing Pod (check Events section)..."
kubectl describe pod nginx-pod
echo ""

# 5. View Pod logs
echo "5. Viewing Pod logs..."
kubectl logs nginx-pod
echo ""

# 6. Execute command in Pod
echo "6. Testing nginx inside the Pod..."
kubectl exec nginx-pod -- curl -s localhost
echo ""

# 7. Get Pod details in YAML format
echo "7. Getting Pod configuration..."
kubectl get pod nginx-pod -o yaml
echo ""

# 8. Create multi-container Pod
echo "8. Creating multi-container Pod..."
kubectl apply -f multi-container-pod.yaml
echo ""
sleep 3

# 9. View logs from specific container
echo "9. Viewing logs from sidecar container..."
kubectl logs multi-container-pod -c sidecar --tail=10
echo ""

# 10. Exec into main container
echo "10. Generating traffic to create logs..."
kubectl exec multi-container-pod -c nginx -- curl -s localhost > /dev/null
kubectl exec multi-container-pod -c nginx -- curl -s localhost/test > /dev/null
echo ""

# 11. Check sidecar logs again (should see access logs)
echo "11. Checking sidecar logs (should show access entries)..."
kubectl logs multi-container-pod -c sidecar --tail=5
echo ""

# 12. Create custom namespace
echo "12. Creating custom namespace..."
kubectl apply -f namespace.yaml
kubectl get namespaces
echo ""

# 13. Create Pod in custom namespace
echo "13. Creating Pod in 'learning' namespace..."
kubectl apply -f simple-pod.yaml -n learning
kubectl get pods -n learning
echo ""

# 14. View all Pods across all namespaces
echo "14. Viewing all Pods..."
kubectl get pods --all-namespaces | grep -E 'NAMESPACE|nginx-pod'
echo ""

# Cleanup instructions
echo "=== Cleanup ==="
echo "To clean up resources, run:"
echo "  kubectl delete -f simple-pod.yaml"
echo "  kubectl delete -f multi-container-pod.yaml"
echo "  kubectl delete namespace learning  # Deletes namespace and all resources in it"
echo ""
echo "Or delete all at once:"
echo "  kubectl delete -f ."
echo "  kubectl delete namespace learning"
