const http = require('http');

const server = http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({
    appName: process.env.APP_NAME || 'Unknown',
    nodeEnv: process.env.NODE_ENV || 'development',
    debug: process.env.DEBUG || 'false',
    database: {
      host: process.env.DB_HOST || 'not set',
      port: process.env.DB_PORT || 'not set',
      poolSize: process.env.DB_POOL_SIZE || 'not set'
    }
  }, null, 2));
});

server.listen(3000, '0.0.0.0', () => {
  console.log(`${process.env.APP_NAME || 'App'} running on port 3000`);
  console.log('Environment:', process.env.NODE_ENV);
});
