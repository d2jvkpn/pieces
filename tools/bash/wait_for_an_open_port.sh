#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

port=7860

echo "==> Waiting port $port to open ..."
while ! nc -z localhost 7860; do
    # wait for 1/10 of the second before check again
    sleep 0.1 && echo -n .
done
echo "==> Port $port is open"
