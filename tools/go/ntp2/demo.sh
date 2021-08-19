#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

go run main.go server --addr :8080

go run main.go client --addr http://127.0.0.1:8080 --delay 10
