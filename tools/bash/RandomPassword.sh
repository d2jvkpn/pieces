#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

len=32
[ $# -gt 1 ] && len=$1

tr -dc '0-9a-zA-Z!@#$%^&*()' < /dev/urandom |
  fold -w ${1:-$len} |
  head -n 5 || true

# head -c $len
