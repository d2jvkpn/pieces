#! /usr/bin/env bash
set -eu -o pipefail

go build -o main main.go
docker build -t hello-img:latest ./

## run
# lsof -i :8081
docker run --detach --publish=18080:8080 --name=hello-inst hello-img:latest
docker logs hello-inst

curl -i localhost:18080
docker logs hello-inst

exit 0
## Enter container
docker exec -it hello-inst bash
curl -i localhost:8080/rover

## remove container
docker stop hello-inst
docker rm hello-inst
