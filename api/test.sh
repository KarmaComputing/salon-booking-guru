#!/bin/bash

SBG_DB_HOST="localhost" \
SBG_DB_PORT="5432" \
SBG_DB_USER="postgres" \
SBG_DB_PASSWORD="veryinsecure" \
go test -v -cover ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
