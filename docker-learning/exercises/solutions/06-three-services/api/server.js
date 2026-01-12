const http = require('http');

const server = http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({
    service: 'API',
    status: 'running',
    database: process.env.DB_HOST || 'not configured',
    timestamp: new Date().toISOString()
  }));
});

server.listen(3000, '0.0.0.0', () => {
  console.log('API server running on port 3000');
});
