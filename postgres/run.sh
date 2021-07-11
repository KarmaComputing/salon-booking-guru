#!/bin/bash

docker run -it \
    --name sbg-postgres \
    -e POSTGRES_PASSWORD="veryinsecure" \
    -e POSTGRES_DB="postgres" \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v sbg-postgres:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres
