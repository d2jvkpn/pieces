#! /bin/bash
set -eu -o pipefail

docker pull centos:8
docker build -t centos:8e1 ./
