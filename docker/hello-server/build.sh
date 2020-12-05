#! /usr/bin/env bash
set -eu -o pipefail

## compile
go build -o hello hello.go

## build
docker build -t hello ./

## run
# lsof -i :8081
docker run -deamon --publish=8081:8080 --name=hello-server hello
docker logs hello-server

curl -i localhost:8081/Rover
docker logs hello-server

## Enter container
#   docer exec -it hello-server bash
#   curl -i localhost:8080
#   exit

## remove container
# docker stop hello-server
# docker rm hello-server
