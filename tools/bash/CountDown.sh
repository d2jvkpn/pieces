#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

if [[ $# -gt 0 && "$1" == *"-h"* ]]; then
    echo "$(basename $0) [timeout] [command] [arg...]"
    echo -e "e.g.\n    timeout: 5, 1m, deault: 15"
    exit 0
fi

secs=15
[ $# -gt 0 ] && secs=$1
if [[ "$secs" == *"m" ]]; then
    secs=$((${secs%m} * 60))
fi

for i in $(seq 1 $secs | tac); do
    echo -en "\r==> $(date +%FT%T%:z) $(printf "%03d" $i)"
    sleep 1
done

shift

if [ $# -le 1 ]; then
    echo -en "\r=== $(date +%FT%T:%:z) END\n"
else
    echo ""
    set -x
    $*
fi

# mpv 01.wav
