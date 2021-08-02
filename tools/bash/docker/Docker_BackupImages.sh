#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

# sudo apt install gridsite-clients
images=$(docker images | awk 'NR>1 && $1!="<none>" && $2!="none"{print $1":"$2}')

for image in $images; do
    output=image_$(urlencode $image).tgz
    echo ">>> saving $image => $output"
    docker save $image | pigz -c > $output
done

# for tgz in $(ls image_*.tgz); do pigz -dc $tgz | docker load; done
