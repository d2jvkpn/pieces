#! /usr/bin/env bash
set -eu -o pipefail

path=$1
output=$(echo $path | sed 's#https://##; s#http://##; s#/#_#; s#/#%2F#').zip

wget -O $output $path/archive/master.zip
