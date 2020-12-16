#! /usr/bin/env bash
set -eu -o pipefail

docker build --squash --build-arg TZ="Asia/Shanghai" -t jupyter:latest .

docker inspect jupyter:latest                |
    jq -r ".[0].Comment"                     |
    awk '{sub("sha256:", "", $2); print $2}' |
    xargs -i docker rmi {}
