FROM mysql:8.0

RUN apt update && \
  apt install -y curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.10.0/migrate.linux-amd64.tar.gz | tar xvz && \
  mv migrate.linux-amd64 /usr/local/sbin/migrate

CMD /bin/bash
