#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


## crontab -e
# @reboot bash /opt/bin/DockerCleanup.sh

docker ps -f status=exited -q | xargs -i docker rm {}
