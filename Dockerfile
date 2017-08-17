FROM alpine:3.6

CMD ["/app/boot.sh"]

RUN set -xeuo pipefail ;\
    apk update ;\
    apk add gcc git go musl-dev nginx wget nginx-mod-http-geoip geoip zlib ;\
    \
    apk add --virtual builddeps \
        build-base \
        curl curl-dev \
        geoip-dev \
        zlib-dev \
    ;\
    mkdir gotmp ;\
    GOPATH=$PWD/gotmp go get -v github.com/kardianos/govendor ;\
    mv gotmp/bin/govendor /usr/local/bin/ ;\
    rm -fr gotmp ;\
    \
    wget -nv -O/usr/bin/gosu "https://github.com/tianon/gosu/releases/download/1.10/gosu-amd64" ;\
    chmod +x /usr/bin/gosu ;\
    \
    wget -nv -O- https://github.com/maxmind/geoipupdate/releases/download/v2.4.0/geoipupdate-2.4.0.tar.gz | tar -xvz ;\
    cd geoipupdate-2.4.0/ ;\
    ./configure --enable-silent-rules ;\
    make install -j 4 V=0 ;\
    src=$PWD && cd .. && rm -fr "$src" ;\
    \
    apk del --purge builddeps ;\
    sh -c 'rm -fr /tmp/* /var/cache/apk/*' ;\
    \
    true

RUN set -xeuo pipefail ;\
    mkdir -p /app/src/api

WORKDIR /app
ENV GIN_MODE=release API_LISTEN=127.0.0.1:8001 GOPATH=/app

RUN mkdir -p /app/src/api/vendor/
COPY src/api/vendor/vendor.json /app/src/api/vendor/

RUN set -xue ;\
    cd /app/src/api ;\
    govendor sync -v ;\
    govendor build -v +vendor

RUN set -xueo pipefail ;\
    mkdir -p /usr/local/share/GeoIP ;\
    sed -i -e 's/^ProductIds/### ProductIds/' /usr/local/etc/GeoIP.conf ;\
    echo 'ProductIds 506 517 533' >> /usr/local/etc/GeoIP.conf ;\
    /usr/local/bin/geoipupdate -v

COPY assets/ assets/
COPY boot.sh /app/boot.sh
COPY src /app/src

RUN set -xeuo pipefail ;\
    echo chmod -R a+rX /app/ /etc/ ;\
    chmod -R go-rwx /etc/crontabs/ ;\
    cd /app/src/api ;\
    go build -v -o /usr/local/bin/api-cli cli/cli.go ;\
    true

COPY etc/ /etc/
RUN nginx -t

ARG COMMIT
ENV COMMIT=$COMMIT
