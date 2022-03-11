#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


## crontab -e
# @reboot bash -c /opt/bin/NerdctlCleanup.sh

nerdctl -n k8s.io ps --format=json -a |
  jq -rs '.[] | select( .Status == "Created").ID' |
  xargs -i nerdctl -n k8s.io rm {}
