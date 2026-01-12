# Exercise 5: Write a Multi-Stage Dockerfile

## Goal
Create a multi-stage Dockerfile to build a Go application with a minimal final image.

## Solution

```bash
cd exercises/solutions/05-multi-stage

# Build the image
docker build -t multi-stage-app .

# Check the image size (should be ~10MB!)
docker images multi-stage-app

# Run it
docker run --rm multi-stage-app

# Compare with single-stage (would be ~300MB+)
```

## Key Concepts

- Multi-stage builds use multiple `FROM` statements
- `AS builder` names a build stage
- `COPY --from=builder` copies from a previous stage
- Final image only contains runtime necessities
- Dramatically reduces image size (300MB -> 10MB)
- Build tools stay in build stage, not final image
