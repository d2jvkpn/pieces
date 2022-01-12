#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

date --rfc-3339=seconds | sed "s/ /T/"

date +%FT%T%:z
