#! /usr/bin/env bash
set -eu -o pipefail

playListURL="$1"

youtube-dl -j --flat-playlist "$playListURL" > list.txt
# jq -r '.id' | sed 's_^_https://youtube.com/_' > download_url.txt

jq -r '(.title + "-" + .id)' list.txt | while read line; do
   line="$(echo "$line" | sed 's/|/_/g; s/\//_/g; s/:/ -/')"
   ls -1 "$line".mp4 "$line".en.vtt
done

n=0
jq -r '(.title + "-" + .id)' list.txt | while read line; do
    n=$((n+1))
    line="$(echo "$line" | sed 's/|/_/g; s/\//_/g; s/:/ -/')"
    mv "$line".mp4 "$n $line".mp4
    test -f "$line".en.vtt && mv "$line".en.vtt "$n $line".en.vtt || true
done
