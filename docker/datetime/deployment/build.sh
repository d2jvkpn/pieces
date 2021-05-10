#! /usr/bin/env bash
set -eu -o pipefail

#### config and check
registry=$(printenv DOCKER_Registry)
BRANCH="$1"
path=$(dirname $0)
test -z $registry && { echo "DOCKER_Registry is unset"; exit 1; }

if [[ "$BRANCH" != "test" && "$BRANCH" != "main" ]]; then
    echo "inlvaid app env for deployment!!!"
    exit -1
fi
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
# sudo docker push $image
