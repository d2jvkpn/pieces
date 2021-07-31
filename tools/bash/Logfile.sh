#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

logFile=$(date +"%FT%T_%s%N" | sed 's/:/-/g')
echo ${logFile::-6}.log
