#!/bin/bash -x

docker run \
    --rm -it \
    -p 5432:5432 \
    --name postgres \
    -e POSTGRES_PASSWORD=passwd \
    -e PGDATA=/var/lib/postgresql/data \
    -v $(pwd)/postgres/data/:/var/lib/postgresql/data \
    -v $(pwd)/init:/docker-entrypoint-initdb.d \
    postgres:13.4
