#! /bin/bash
set -eu

input=$1
start=$2    # 00:13:48
duration=$3 # 00:04:18
output=$4

ffmpeg -i $input -ss $start -t $duration -acodec copy -vcodec copy $output
