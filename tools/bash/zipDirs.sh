#! /usr/bin/env bash
set -eu -o pipefail

SAVEIFS=$IFS
IFS=$(echo -en "\n\b")
for d in $(ls -d */); do
    echo ">>> zip $d"
    tar -r "${d%/}.zip" "$d"
    rm -r "$d"
done
IFS=$SAVEIFS
