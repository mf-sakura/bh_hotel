FROM golang:1.12.6-stretch

WORKDIR /

RUN apt-get update \
    && apt-get install --no-install-recommends \
    -y build-essential git mysql-client

# Install goose
RUN go get bitbucket.org/liamstask/goose/cmd/goose

COPY . .

CMD ["sh", "docker-entrypoint.sh"]