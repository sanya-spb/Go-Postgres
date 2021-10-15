#!/bin/bash -x

TASK='task05'
DATADIR=$(pwd)/../postgres/data/
INITDIR=$(pwd)/../${TASK}/init

docker build -t pg-ext ../${TASK}

docker run \
    --rm -it \
    -p 5432:5432 \
    --name pg-ext \
    -v ${DATADIR}:/var/lib/postgresql/data \
    pg-ext
