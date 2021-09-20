#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

logFile=$(date +"%F_%T_%Z_%s%N" | sed 's/:/-/g')
logFile=${logFile::-6}
echo $logFile
