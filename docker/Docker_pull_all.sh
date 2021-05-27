#! /usr/bin/env bash
set -eu -o pipefail

docker images | awk 'NR>1 && $2!="local" && $2!="<none>"{print $1":"$2}' | xargs -i docker pull {}

docker images | awk 'NR>1 && $2=="<none>"{print $3}' | xargs -i docker rmi {}
