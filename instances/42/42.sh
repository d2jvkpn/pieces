#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

echo "Hello, world!"

ans=$(echo "(-80538738812075974)^3 + 80435758145817515^3 + 12602123297335631^3" | bc)
echo "Life, the Universe and Everything: $ans"
