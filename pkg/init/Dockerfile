FROM alpine:3.10.2

RUN apk add --no-cache openssl \
    && apk add --no-cache --virtual .build-deps curl \
    && curl -L https://github.com/simon-engledew/addr/releases/download/1.0.0/addr-linux-amd64.tgz | tar xvzp -C /usr/bin \
    && apk del .build-deps

COPY root/ /

VOLUME /var/disk/certs

CMD ["/usr/bin/entrypoint.sh"]
