FROM node:20-alpine

WORKDIR /app

# Copy package files first for better layer caching
COPY package.json .

# Install dependencies
RUN npm install --production

# Copy application code
COPY server.js .

# Document the port
EXPOSE 3000

# Run the application
CMD ["node", "server.js"]
