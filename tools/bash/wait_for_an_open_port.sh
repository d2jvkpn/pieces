#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

echo "Waiting port 7860 to open ..."
while ! nc -z localhost 7860; do   
  sleep 0.1 # wait for 1/10 of the second before check again
done
echo "Port 7860"
