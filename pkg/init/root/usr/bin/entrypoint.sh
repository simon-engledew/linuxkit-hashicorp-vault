#!/bin/sh

set -e

if [ ! -f /var/disk/certs/https.pem ]; then
    openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
        -subj "/CN=$(addr eth0)" \
        -keyout /var/disk/certs/https.key  -out /var/disk/certs/https.crt

    (umask 0077; cat /var/disk/certs/https.key /var/disk/certs/https.crt > /var/disk/certs/https.pem)
fi
