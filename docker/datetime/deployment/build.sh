#! /usr/bin/env bash
set -eu -o pipefail

#### config and check
registry=$(printenv DOCKER_Registry)
BRANCH="$1"
path=$(dirname $0)
test -z $registry && { echo "DOCKER_Registry is unset"; exit 1; }

image="$registry/datetime:${BRANCH}"

#### build local image
echo ">>> Building image: $image"
docker pull golang:1.16-alpine
docker pull alpine
git checkout -f $BRANCH
docker build -f $path/Dockerfile --no-cache -t "$image" .
docker image prune --force --filter label=stage=datetime_builder


#### push to registry
echo ">>> Pushing image: $image"
docker push $image
# sudo docker push $image
