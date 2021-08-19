#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

go run server.go :8080

go run client.go -addr http://127.0.0.1:8080 -delay 10
