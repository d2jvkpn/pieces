#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

jf=$1

n=0
for link in $(jq -r .links[] $jf); do
    n=$((n+1))
    name=$(basename $link)
    name=$(printf "%03d_%s" $n $name)
    # wget -cO $name $link
    echo $name $link
done | xargs -n 2 -P 8 wget -O
