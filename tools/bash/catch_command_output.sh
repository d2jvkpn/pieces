#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

out=$(ls zzzzzzzz 2>&1 || true)
echo $?
echo $out
