#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_self=$(readlink -f $0)
_path=$(dirname ${_self})


echo -e 'Rover\nooo\nAlice: Hey!\n' | nc localhost 8080

exit
Rover
ooo
Alic: Hey!
