# Exercise 2: Build a Custom Image with a Simple Web App

## Goal
Create a Dockerfile and build a custom image with a simple web application.

## Solution

Build and run:
```bash
cd exercises/solutions/02-custom-image
docker build -t my-webapp .
docker run -d --name my-webapp -p 3000:3000 my-webapp

# Test it
curl http://localhost:3000

# Cleanup
docker stop my-webapp
docker rm my-webapp
docker rmi my-webapp
```

## Key Concepts

- `FROM` specifies the base image
- `WORKDIR` sets the working directory
- `COPY` copies files from host to image
- `RUN` executes commands during build
- `EXPOSE` documents the port (informational)
- `CMD` specifies the default command to run
