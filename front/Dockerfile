FROM node:13.10.1-alpine3.10

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli@4.1.2 \
    npm ci

CMD npm run serve
