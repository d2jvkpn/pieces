#! /bin/bash
set -eu

for i in $@; do
    if [ ! -s $i ]; then
        echo "$i: file not exists, skip"
        continue
    fi

    jpg=${i%.*}.jpg

    if [[ "$jpg" == "$i" ]]; then
        echo "    $i, convert is trying to overwritten itself, skip"
        continue
    fi
    
    convert -verbose -density 150 -trim  $i -quality 100 \
    -flatten -sharpen 0x1.0 $jpg
done
