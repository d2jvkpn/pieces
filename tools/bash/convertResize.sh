#! /usr/bin/env bash
set -eu -o pipefail

# convert -resize 20% input.jpg output.small.jpg
convert  -resize  "$1"%  "$2"  "$3"
