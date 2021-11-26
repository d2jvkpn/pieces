#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)/
_path=$(dirname $0 | xargs -i readlink -f {})/

PORT=$1


export PORT=${PORT}
envsubst < ${_path}/deploy.yaml > docker-compose.yaml

docker-compose pull
[[ ! -z "$(docker ps --all --quiet --filter name=node-web)" ]] && docker rm -f node-web
# docker-compose down for running containers only, not stopped containers

docker-compose up -d
