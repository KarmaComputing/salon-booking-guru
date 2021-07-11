#!/bin/bash
go get -d -v ./... && \
go install -v ./... && \
go build && \
SBG_DB_HOST="localhost" \
SBG_DB_PORT="5432" \
SBG_DB_USER="postgres" \
SBG_DB_PASSWORD="veryinsecure" \
SBG_DB_DBNAME="postgres" \
./salon-booking-guru
