const http = require('http');
const os = require('os');

const APP_NAME = process.env.APP_NAME || 'App';
const APP_PORT = process.env.APP_PORT || 3000;

const server = http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({
    service: APP_NAME,
    hostname: os.hostname(),
    path: req.url,
    headers: {
      'x-real-ip': req.headers['x-real-ip'] || 'direct',
      'x-forwarded-for': req.headers['x-forwarded-for'] || 'direct',
      'x-forwarded-proto': req.headers['x-forwarded-proto'] || 'http'
    },
    timestamp: new Date().toISOString()
  }, null, 2));
});

server.listen(APP_PORT, '0.0.0.0', () => {
  console.log(`${APP_NAME} running on port ${APP_PORT}`);
});
