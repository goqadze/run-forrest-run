# Exercise 7: Optimize a Dockerfile to Reduce Image Size

## Goal
Learn techniques to optimize Dockerfiles for smaller, more efficient images.

## Optimization Techniques

1. **Use smaller base images** (alpine vs ubuntu)
2. **Multi-stage builds** (separate build from runtime)
3. **Minimize layers** (combine RUN commands)
4. **Clean up in same layer** (remove caches)
5. **Use .dockerignore** (exclude unnecessary files)
6. **Order matters** (put frequently changing files last)

## Files

- `Dockerfile.bad` - Unoptimized (~900MB)
- `Dockerfile.good` - Optimized (~150MB)

## Solution

```bash
cd exercises/solutions/07-optimize-dockerfile

# Build both versions
docker build -f Dockerfile.bad -t app:unoptimized .
docker build -f Dockerfile.good -t app:optimized .

# Compare sizes
docker images | grep app

# Cleanup
docker rmi app:unoptimized app:optimized
```
