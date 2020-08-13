FROM golang:1.10.3-alpine
MAINTAINER John McKee <john@subledger.com>

ENV DE_MIGRATIONS_PATH="/migrations"

RUN apk update                && \
    apk add bash              && \
    apk add postgresql-client && \
    apk add curl

COPY migrator /
COPY executables/* /migrations/

WORKDIR /

LABEL com.subledger.deployer.push="true"
LABEL com.subledger.deployer.migrator="true"

CMD ["/migrator"]
