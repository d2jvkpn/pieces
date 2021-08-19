#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

echo "date +'%FT%T%z %N' > now.txt" | at 15:32
