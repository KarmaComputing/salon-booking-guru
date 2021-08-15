#!/bin/bash

SALON_BOOKING_GURU_DB_HOST="localhost" \
SALON_BOOKING_GURU_DB_PORT="5432" \
SALON_BOOKING_GURU_DB_USER="postgres" \
SALON_BOOKING_GURU_DB_PASSWORD="changeme" \
go test -v -cover ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
