FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache nodejs npm
COPY package.json .
COPY package-lock.json .
COPY index.js .
CMD ["/bin/sh"]