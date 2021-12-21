#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


url=$1
output=$(2>&1 curl -s -I $url  | awk 'NR==1{print; exit}')

if [[ "$output" =~ "200 OK" ]]; then
    echo true
else
    echo false
fi
