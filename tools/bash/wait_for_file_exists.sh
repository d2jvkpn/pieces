#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

until [ ! -f /data/ok ]; do
    echo ">>> $(date +%FT%T%z) file not exists: /data/ok"
    sleep 1
done
