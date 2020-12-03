#! /usr/bin/env bash
set -eu -o pipefail

SAVEIFS=$IFS
IFS=$(echo -en "\n\b")
for d in $(ls -d */); do
    echo ">>> tar $d"
    tar -cf "${d%/}.tar" "$d"
    rm -r "$d"
done
IFS=$SAVEIFS
