FROM golang:1.18.2-alpine

ENV GO111MODULE=on

ENV MYSQL_USER=${MYSQL_USER}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
ENV MYSQL_DATABASE=${MYSQL_DATABASE}

RUN mkdir /app

WORKDIR /app
COPY . /app

# Install & Run migrations
RUN wget -O /tmp/migrations.tar.gz "https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-386.tar.gz" && \
    tar -xzvf /tmp/migrations.tar.gz -C /tmp/ && \
    mv /tmp/migrate /usr/bin/migrate

EXPOSE 8080

CMD ["sh", "run.sh"]