#!/usr/bin/env bash

export SERV_ADDR=localhost:22080
export CERT_FILE=~/.keys/localhost.cert
export KEY_FILE=~/.keys/localhost.key

go run ./cmd/runserver.go
