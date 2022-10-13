#!/usr/bin/env bash

set -ev

docker-compose up -d
go test -coverprofile=./c.out -v -race -timeout 30m ./...
cp c.out coverage.txt
docker-compose down -v