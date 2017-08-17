FROM golang:1.9-alpine

CMD ["/app/boot.sh"]

RUN set -xeuo pipefail ;\
    apk add --no-cache ca-certificates git nginx ;\
    \
    apk add --no-cache --virtual builddeps \
        wget \
    ;\
    go get -v github.com/kardianos/govendor ;\
    \
    wget -nv -O/usr/bin/gosu "https://github.com/tianon/gosu/releases/download/1.10/gosu-amd64" ;\
    chmod +x /usr/bin/gosu ;\
    \
    apk del --purge builddeps ;\
    mkdir -p /app/src/api/vendor

WORKDIR /app
ENV GIN_MODE=release API_LISTEN=127.0.0.1:8001 GOPATH=/app

COPY src/api/vendor/vendor.json /app/src/api/vendor/

RUN set -xue ;\
    cd /app/src/api ;\
    govendor sync -v ;\
    govendor build -v +vendor

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
