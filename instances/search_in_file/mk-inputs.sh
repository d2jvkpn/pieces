#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


for x in {a..z}; do
    echo $x$x$x
done > a01.txt
