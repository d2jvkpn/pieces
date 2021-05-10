#! /usr/bin/env bash
set -eu -o pipefail

#### config and check
registry=$(printenv DOCKER_Registry)
TAG="$1"
path=$(dirname $0)
BRANCH=""
test -z $registry && { echo "DOCKER_Registry is unset"; exit 1; }

case $TAG in
  "test")
    BRANCH="$TAG"
    ;;
  "prod")
    BRANCH="main"
    ;;
  *)
    echo "inlvaid tag for build!!!"
    exit -1
    ;;
esac
image="$registry/datetime:${BRANCH}"

#### build local image
echo ">>> building image: $image"
git checkout -f $BRANCH
docker pull golang:1.16-alpine
docker pull alpine
docker build -f $path/Dockerfile --no-cache -t "$image" .
docker image prune --force --filter label=stage=datetime_builder


#### push to registry
echo ">>> pushing image: $image"
docker push $image
