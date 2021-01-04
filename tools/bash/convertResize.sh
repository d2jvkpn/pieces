#! /usr/bin/env bash
set -eu -o pipefail

# convert -resize 20% IMG_E4062.jpg IMG_E4062.small.jpg

convert  -resize  "$1"%  "$2"  "$3"
