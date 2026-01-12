FROM node:20-alpine

WORKDIR /app

COPY server.js .

ARG APP_NAME="App"
ARG APP_PORT="3000"

ENV APP_NAME=${APP_NAME}
ENV APP_PORT=${APP_PORT}

EXPOSE ${APP_PORT}

CMD ["node", "server.js"]
