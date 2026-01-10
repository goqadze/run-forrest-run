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
