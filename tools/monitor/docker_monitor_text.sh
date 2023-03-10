#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

# crontab: */1 * * * *   bash ~/Apps/docker_monitor/docker_monitor.sh

outfile=$0.$(date +%Y-%m).log
timestamp=$(date +%FT%T:%:z)

{
    echo -e "\n#### $timestamp"
    docker stats --no-stream
} >> $outfile
