#!/usr/bin/env bash

set -e

docker-compose up -d
go test -coverprofile=./c.out -v -race ./...
docker-compose down -v