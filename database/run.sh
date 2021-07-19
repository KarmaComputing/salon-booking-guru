#!/bin/bash

docker stop postgres-salon-booking-guru > /dev/null | true
docker run -it \
    --rm \
    -d \
    --name postgres-salon-booking-guru \
    -e POSTGRES_PASSWORD="changeme" \
    -e POSTGRES_DB="salon_booking_guru" \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v postgres:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres
