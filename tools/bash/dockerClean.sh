#! /bin/bash
set -eu -o pipefail

# remove exited containers
exitedContainers=$(docker ps --all -q -f status=exited)
test -z "$exitedContainers" || docker rm $exitedContainers

# remove images named <none>
docker images | awk 'NR>1 && $1=="<none>"{print $3}' |
perl -e 'print reverse <>' | xargs -i docker rmi {}
