#! /usr/bin/env bash
set -eu -o pipefail

SAVEIFS=$IFS
IFS=$(echo -en "\n\b")
for d in $(ls -d */); do
    echo ">>> tar $d"
    tar -cf "${d%/}.tar" "$d" && rm -r "$d"
done
IFS=$SAVEIFS


exit 0

## alternative
find * -prune -type d | while IFS= read -r d; do 
    tar -cf "$d.tar" "$d" && rm -r "$d"
done
# find . -printf "%y %p\n"
