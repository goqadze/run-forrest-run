const http = require('http');
const fs = require('fs');

// Read secrets from files
function readSecret(name) {
  const path = `/run/secrets/${name}`;
  try {
    return fs.readFileSync(path, 'utf8').trim();
  } catch (err) {
    return `Secret '${name}' not found`;
  }
}

const server = http.createServer((req, res) => {
  // Read secrets (in production, cache these)
  const dbPassword = readSecret('db_password');
  const apiKey = readSecret('api_key');

  res.writeHead(200, { 'Content-Type': 'application/json' });
  res.end(JSON.stringify({
    message: 'Secrets are loaded securely!',
    secrets: {
      db_password: dbPassword ? '***loaded***' : 'not found',
      api_key: apiKey ? '***loaded***' : 'not found',
      // Show lengths for debugging (never show actual values!)
      db_password_length: dbPassword.length,
      api_key_length: apiKey.length
    },
    tip: 'Secrets are stored in /run/secrets/ as files'
  }, null, 2));
});

server.listen(3000, '0.0.0.0', () => {
  console.log('Secrets demo server running on port 3000');
});
