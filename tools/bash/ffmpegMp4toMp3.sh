#! /bin/bash
set -eu

mp4=$1

ffmpeg -i "$mp4" -vn -ar 44100 -ac 2 -ab 192k -f mp3 "${mp4%.mp4}.mp3"
