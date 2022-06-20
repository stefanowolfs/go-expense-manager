#!/bin/bash
SERVER_ADDRESS=localhost \
SERVER_PORT=3000 \
DB_USER=root \
DB_PASSWD=password \
DB_ADDR=localhost \
DB_PORT=3306 \
DB_NAME=expenses \
go run main.go
