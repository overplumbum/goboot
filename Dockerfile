FROM golang:1.9-alpine

CMD ["/app/boot.sh"]

RUN set -xeuo pipefail ;\
    apk add --no-cache ca-certificates git nginx su-exec ;\
    \
    apk add --no-cache --virtual builddeps \
        wget \
    ;\
    go get -u github.com/golang/dep/cmd/dep ;\
    \
    apk del --purge builddeps ;\
    mkdir -p /app/src/api

WORKDIR /app
ENV GIN_MODE=release API_LISTEN=127.0.0.1:8001 GOPATH=/app

COPY src/api/Gopkg.toml src/api/Gopkg.lock /app/src/api/

RUN set -xue ;\
    cd /app/src/api ;\
    dep ensure -vendor-only -v

COPY assets/ assets/
COPY boot.sh /app/boot.sh
COPY etc/ /etc/
RUN nginx -t
COPY src /app/src

RUN set -xeuo pipefail ;\
    echo chmod -R a+rX /app/ /etc/ ;\
    chmod -R go-rwx /etc/crontabs/ ;\
    cd /app/src/api ;\
    go build -v -o /usr/local/bin/api-cli cli/cli.go

ARG COMMIT
ENV COMMIT=$COMMIT
