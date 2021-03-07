#! /usr/bin/env bash
set -eu -o pipefail

export GIT_REPO="https://github.com/d2jvkpn/hello-service"
export BRANCH="main"
export PROG="hello-service"
export PORT=8080
export GO_VERSION="1.16"
export TAG=$(date +%Y%m%dT%H%M%S)

docker build --squash              \
  --build-arg GIT_REPO="$GIT_REPO" \
  --build-arg PROG="$PROG"         \
  --build-arg BRANCH="$BRANCH"     \
  --build-arg PORT=$PORT           \
  --build-arg GO_VERSION="$GO_VERSION"  \
  -t $PROG:$TAG .
#  --no-cache

docker inspect $PROG:$TAG |
    jq -r ".[0].Comment"  |
    awk '{sub("sha256:", "", $2); print $2}' |
    xargs -i docker rmi {}

docker run --detach --publish=$PORT:$PORT --name=$PROG $PROG:$TAG

curl localhost:$PORT

exit 0

#### docker support --squash
mkdir -p /etc/docker/

cat > /etc/docker/daemon.json << EOF
{
  "experimental": true
}
EOF

systemctl restart docker
