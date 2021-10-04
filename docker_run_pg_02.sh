#!/bin/bash -x

DATADIR=$(pwd)/postgres/data/
INITDIR=$(pwd)/task02/init

docker build -t pg-ext ./task02

docker run \
    --rm -it \
    -p 5432:5432 \
    --name pg-ext \
    -v ${DATADIR}:/var/lib/postgresql/data \
    -v ${INITDIR}:/docker-entrypoint-initdb.d \
    pg-ext
