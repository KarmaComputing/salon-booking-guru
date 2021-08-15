#!/bin/bash
go get -d -v ./... && \
go install -v ./... && \
go build && \
SALON_BOOKING_GURU_DB_HOST="localhost" \
SALON_BOOKING_GURU_DB_PORT="5432" \
SALON_BOOKING_GURU_DB_USER="postgres" \
SALON_BOOKING_GURU_DB_PASSWORD="changeme" \
SALON_BOOKING_GURU_DB_DBNAME="salon_booking_guru" \
./salon-booking-guru
