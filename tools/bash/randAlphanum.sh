#! /usr/bin/env bash
set -eu -o pipefail

width=$1
count=1
test $# -gt 1 && count=$2

cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w ${1:-$width} | head -n $count
