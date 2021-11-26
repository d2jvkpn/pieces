#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})/

branch=latest
name="registry.cn-shanghai.aliyuncs.com/d2jvkpn/node-web"
image="$name:${branch}"


####
echo ">>> docker build $image"
docker pull node:14-alpine &> /dev/null
docker build --no-cache -f "${_path}"/Dockerfile --tag "$image" .

docker push $image

imgs=$(docker images --filter=dangling=true --quiet $name)
[[ -z "$imgs" ]] || docker rmi $imgs &> /dev/null
