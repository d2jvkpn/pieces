#! /usr/bin/env bash
set -eu -o pipefail

#### config
wd=$(pwd)
path=$(dirname $0)
registry=$(printenv DOCKER_Registry)
BRANCH="$1"
PORT=""
test -z $registry && { echo "DOCKER_Registry is unset"; exit 1; }

case $BRANCH in
  "test")
    PORT="1032"
    ;;
  "main")
    PORT="1030"
    ;;
  *)
    echo "inlvaid BRANCH for build!!!"
    exit -1
    ;;
esac
image="$registry/datetime:${BRANCH}"


#### deploy
docker pull $image
export DOCKER_Registry=${DOCKER_Registry}
export BRANCH=${BRANCH}
export PORT=${PORT}
envsubst < $path/deployment.yaml > docker-compose.yaml
docker-compose up -d

echo ">>> http server port: $PORT"
docker logs datetime_${BRANCH}_service
