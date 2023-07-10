#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

if [[ $# -gt 0 && "$1" == *"-h"* ]]; then
    echo "$(basename $0) [timeout] [command] [arg...]"
    echo -e "e.g.\n    timeout: 5s, 1m, deault: 15s"
    echo -e "\n    Countdown.sh 15 mpv ~/Downloads/sounds/01.wav"
    exit 0
fi

secs=15s; [ $# -gt 0 ] && secs="$1"; shift; cmd="$*"

if [[ ! "$secs" =~ ^[0-9]+(m|s)$ ]]; then
    echo "invalid time interval" >&2
    exit 1
fi

if [[ "$secs" == *"s" ]]; then
    secs=${secs%s}
elif [[ "$secs" == *"m" ]]; then
    secs=$((${secs%m} * 60))
fi

sp="-\|/"; j=1
for i in $(seq 1 $secs | tac); do
    c=${sp:j++%${#sp}:1}
    echo -en "\r$c $c $c $(date +%FT%T%:z) $(printf "%03d" $i)"
    sleep 1
done

if [[ -z "$cmd" && ! -f ${_path}/Countdown.default.sh ]]; then
    echo -en "\r=== $(date +%FT%T:%:z) END\n"
elif [[ -z "$cmd" && -f ${_path}/Countdown.default.sh ]]; then
    echo ""
    bash ${_path}/Countdown.default.sh
else
    echo ""
    set -x
    $cmd
fi
