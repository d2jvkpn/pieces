#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

# crontab: */1 * * * *   bash ~/Apps/docker_monitor/docker_monitor.sh

# command -v jq || { >&2 echo "jq not found"; exit 1; }

# outfile=${_path}/$(date +%Y-%m).log
outfile=$0.$(date +%Y-%m).log
timestamp=$(date +%FT%T:%:z)

docker stats -a --no-stream --format json |
  jq -c --arg t $timestamp '. + {timestamp: $t}' >> $outfile
