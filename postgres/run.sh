#!/bin/bash

docker stop postgres-salon-booking-guru > /dev/null | true
docker run -it \
    --rm \
    -d \
    --name postgres-salon-booking-guru \
    -e POSTGRES_PASSWORD="veryinsecure" \
    -e POSTGRES_DB="postgres" \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v postgres:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres
