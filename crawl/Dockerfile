FROM node:13.10.1-alpine3.10

WORKDIR /app

COPY package.json .
COPY package-lock.json .
COPY index.js .

RUN apk update && \
    npm install -g npm && \
    npm ci

CMD node index.js