#! /usr/bin/env bash
set -eu -o pipefail

## compile
go build -o hello.x hello.go

## build
docker build -t hello ./

## run
# lsof -i :8081
docker run --publish=8081:8080 --name=hello.x -d hello
docker logs hello.x

curl -i localhost:8081/Rover
docker logs hello.x

## Enter container
#   docer exec -it hello.x bash
#   curl -i localhost:8080
#   exit

## remove container
# docker stop hello.x
# docker rm hello.x
