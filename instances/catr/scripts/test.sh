#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

touch tests/inputs/cant-touch-this && chmod 000 tests/inputs//cant-touch-this

cargo run -- blargh tests/inputs/cant-touch-this tests/inputs/fox.txt
