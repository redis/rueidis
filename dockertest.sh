#!/usr/bin/env bash

set -ev

package=${1:-./...}
pkgbase="github.com/redis/rueidis"

[[ "/om /rueidiscompat /rueidislock /rueidisotel ./..." =~ "${package#$pkgbase}" || -z "${package#$pkgbase}" ]] && \
docker-compose up -d
go test -coverprofile=coverage.out -v -race -timeout 30m $package
docker-compose down -v
