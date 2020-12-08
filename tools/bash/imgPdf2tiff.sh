#! /bin/bash
set -eu

for i in $@; do
    if [ -s $i ]; then
        echo "$i: file not exists, skip"
        continue
    fi

    convert -density 300 "$i" ${i%.pdf}.tiff
done
