#! /usr/bin/env bash
set -eu -o pipefail

SAVEIFS=$IFS
IFS=$(echo -en "\n\b")
for d in $(ls -d */); do
    echo ">>> zip $d"
    d=${d%/}
    zip -r "$d.zip" $d && rm -rf "$d"
done
IFS=$SAVEIFS
