#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

if [[ $# -eq 0 || "$1" == *"-h"* ]]; then
    echo "$(basename $0) [timeout] [command] [args...]"
    echo -e "e.g.\n    timeout: 5s, 1m\n    Countdown.sh 15s mpv ~/Downloads/sounds/01.wav"
    exit 0
fi

secs=$1
shift
cmd="$*"

if [[ ! "$secs" =~ ^[0-9]+(m|s)$ ]]; then
    echo "invalid time interval" >&2
    exit 1
elif [[ "$secs" == *"s" ]]; then
    secs=${secs%s}
elif [[ "$secs" == *"m" ]]; then
    secs=$((${secs%m} * 60))
fi

sp='|/-\'
j=0

for i in $(seq 1 $secs | tac); do
    c=${sp:j++%${#sp}:1}
    echo -en "\r$c $(date +%FT%T%:z) $(printf "%03d" $i)"
    sleep 1
done

echo -e "\r= $(date +%FT%T:%:z) END\n"

if [[ ! -z "$cmd" ]]; then
    $cmd
elif [[ -f ${_path}/Countdown.default.sh ]]; then
    bash ${_path}/Countdown.default.sh
fi
