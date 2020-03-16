FROM node:12.14.0-alpine

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli@4.1.2 \
    npm install

CMD ["/bin/sh"]