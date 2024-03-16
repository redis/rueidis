#!/usr/bin/env bash

set -ev

trap "docker-compose down -v" EXIT
docker-compose up -d
sleep 5
go install gotest.tools/gotestsum@v1.10.0
gotestsum --format standard-verbose --junitfile unit-tests.xml -- -coverprofile=coverage.out -race -timeout 30m "$@"
