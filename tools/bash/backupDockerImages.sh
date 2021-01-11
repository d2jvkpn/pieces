#! /usr/bin/env bash
set -eu -o pipefail

for img in $(docker images | awk 'NR>1{print $1":"$2}'); do
  echo ">>> $img"
  output=$(echo image_$img.tar | sed 's#/#%2F#g; s#:#%3A#g')
  docker save $img -o $output
  pigz $output
done
