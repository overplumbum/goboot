#!/bin/sh
set -xe
# /usr/bin/gosu nobody:nobody /app/manage.py migrate -v3 --noinput

#mkdir -p /tmp/django_cache/
#chown -R nobody:nobody /tmp/django_cache/

exec init
