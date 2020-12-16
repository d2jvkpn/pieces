#! /usr/bin/env bash
set -eu -o pipefail

export GIT_REPO="https://github.com/d2jvkpn/hello-service"
export PROG="hello-service"
export PORT=8080
export TAG=`date +%Y%m%dT%H%M`

docker build --squash              \
  --build-arg GIT_REPO="$GIT_REPO" \
  --build-arg PROG="$PROG"         \
  --build-arg BRANCH="main"        \
  --build-arg PORT=$PORT           \
  --build-arg GO_VERSION="1.15.6"  \
  -t $PROG:$TAG .
# --no-cache

docker inspect $PROG:$TAG                    |
    jq -r ".[0].Comment"                     |
    awk '{sub("sha256:", "", $2); print $2}' |
    xargs -i docker rmi {}

docker run --detach --publish=$PORT:$PORT --name=$PROG $PROG:$TAG

curl localhost:$PORT/rover

exit 0
#### docker support --squash
cat > /etc/docker/daemon.json << 'EOF'
{
  "experimental": true
}
EOF

systemctl restart docker
