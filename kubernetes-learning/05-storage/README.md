# Storage

Persistent storage in Kubernetes using Volumes, PersistentVolumes, and PersistentVolumeClaims.

## Volume Types

**emptyDir** - Temporary storage (deleted with Pod)
**hostPath** - Mount from node's filesystem
**PersistentVolume (PV)** - Cluster-wide storage resource
**PersistentVolumeClaim (PVC)** - User request for storage

## Files

1. **emptydir-volume.yaml** - Temporary storage
2. **persistent-volume.yaml** - PV definition
3. **persistent-volume-claim.yaml** - PVC definition
4. **pod-with-pvc.yaml** - Pod using PVC
5. **commands.sh** - kubectl commands

## Next: Complete App (06-complete-app/)
