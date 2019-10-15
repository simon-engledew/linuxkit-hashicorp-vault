FROM alpine:3.10.2 AS dhparam

RUN apk --no-cache add openssl \
  && openssl dhparam -out /etc/ssl/dhparam.pem 2048 \
  && apk del openssl \
  ;

FROM haproxy:1.8.20-alpine

COPY --from=dhparam /etc/ssl/dhparam.pem /etc/ssl/dhparam.pem

RUN apk add --no-cache tini
ENTRYPOINT ["/sbin/tini", "/docker-entrypoint.sh"]
CMD ["haproxy", "-f", "/usr/local/etc/haproxy/haproxy.cfg"]

COPY root/ /

EXPOSE 80 443

VOLUME /var/disk/certs /dev/log
