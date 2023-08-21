#! /bin/bash
set -eu

for pdf in $@; do
    echo "Convert $pdf"
    png=${pdf%.pdf}.png

    if  [ ! -s "$pdf" ]; then
        echo "    $pdf, file not exists, skip"
        continue
    fi

    if [[ "$png" == "$pdf" ]]; then
        echo "    $pdf, convert is trying to overwritten itself, skip"
        continue
    fi

    gs -q -o $png -sDEVICE=pngalpha -dLastPage=1 -r144 $pdf 2
done
