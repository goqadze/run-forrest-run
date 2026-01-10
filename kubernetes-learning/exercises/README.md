# Practice Exercises

Test your Kubernetes skills with these challenges!

## Beginner Exercises

1. **Create a Pod**
   - Deploy a Pod running `httpd:2.4`
   - Name it `apache-pod`
   - Verify it's running

2. **Use a Namespace**
   - Create a namespace called `practice`
   - Deploy nginx in that namespace
   - List all Pods in the namespace

3. **Create a Deployment**
   - Create a deployment with 3 replicas of nginx
   - Scale it to 5 replicas
   - Scale it back to 2

4. **Expose with Service**
   - Create a deployment
   - Expose it with a ClusterIP service
   - Expose it with a NodePort service

## Intermediate Exercises

5. **ConfigMap Usage**
   - Create a ConfigMap with database settings
   - Create a Pod that reads those settings as env vars

6. **Secrets**
   - Create a Secret with username/password
   - Mount it as a volume in a Pod

7. **Rolling Update**
   - Create a deployment with nginx:1.20
   - Update to nginx:1.21
   - Check rollout status
   - Rollback to previous version

8. **Resource Limits**
   - Create a deployment with CPU and memory limits
   - Set requests and limits appropriately

## Advanced Exercises

9. **Multi-Tier Application**
   - Deploy a database (postgres)
   - Deploy a web application
   - Connect them with services
   - Use ConfigMaps for database connection info

10. **Troubleshooting**
    - Deploy an intentionally broken Pod
    - Use describe and logs to find the issue
    - Fix the problem

## Project: Deploy a Real Application

Deploy a complete 2-tier application:
- PostgreSQL database with persistent storage
- Node.js/Python web app connecting to database
- Service to expose the web app
- ConfigMaps for configuration
- Secrets for database credentials

Good luck!
