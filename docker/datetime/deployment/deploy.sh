#! /usr/bin/env bash
set -eu -o pipefail

#### config
wd=$(pwd)
path=$(dirname $0)
registry=$(printenv DOCKER_Registry)
TAG="$1"
PORT=""
BRANCH=""
test -z $registry && { echo "DOCKER_Registry is unset"; exit 1; }

case $TAG in
  "test")
    PORT="1032"
    BRANCH="$TAG"
    ;;
  "prod")
    PORT="1030"
    BRANCH="main"
    ;;
  *)
    echo "inlvaid tag for build!!!"
    exit -1
    ;;
esac
image="$registry/datetime:${BRANCH}"


#### deploy
docker pull $image
export TAG=${TAG}
export PORT=${PORT}
envsubst < $path/deployment.yaml > docker-compose.yaml
docker-compose up -d

docker logs datetime_${TAG}_service
