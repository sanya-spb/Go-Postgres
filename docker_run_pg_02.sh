#!/bin/bash -x

DATADIR=$(pwd)/postgres/data/
INITDIR=$(pwd)/task02/init

docker run \
    --rm -it \
    -p 5432:5432 \
    --name postgres \
    -e POSTGRES_PASSWORD=passwd \
    -e PGDATA=/var/lib/postgresql/data \
    -v ${DATADIR}:/var/lib/postgresql/data \
    -v ${INITDIR}:/docker-entrypoint-initdb.d \
    postgres:13.4
