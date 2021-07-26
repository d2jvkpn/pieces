#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

function call() {
    date +"%FT%T%z"
    sleep 10
}

export -f call

while true; do
    timeout 10 bash -c -- "call" || true
done
