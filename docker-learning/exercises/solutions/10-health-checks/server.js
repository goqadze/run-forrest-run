const http = require('http');
const fs = require('fs');

// Simulate unhealthy state with a flag file
const isHealthy = () => !fs.existsSync('/tmp/unhealthy');

const server = http.createServer((req, res) => {
  if (req.url === '/health') {
    if (isHealthy()) {
      res.writeHead(200, { 'Content-Type': 'application/json' });
      res.end(JSON.stringify({ status: 'healthy', timestamp: new Date().toISOString() }));
    } else {
      res.writeHead(503, { 'Content-Type': 'application/json' });
      res.end(JSON.stringify({ status: 'unhealthy', timestamp: new Date().toISOString() }));
    }
    return;
  }

  res.writeHead(200, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({
    message: 'Hello from health-checked service!',
    healthy: isHealthy()
  }));
});

server.listen(3000, '0.0.0.0', () => {
  console.log('Server with health check running on port 3000');
  console.log('Health endpoint: http://localhost:3000/health');
  console.log('To make unhealthy: touch /tmp/unhealthy');
});
