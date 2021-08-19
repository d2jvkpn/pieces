#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

# remove ansi color codes, e.g. \x1b, \x1b[30m, \x1b[31m
sed 's/\\x1b\[[0-9;]*m//g'

# remove ansi color codes in container log(/var/lib/docker/containers/?/?-json.log)
## e.g. \u001b[1m, \u001b[37m
sed 's/\\u001b\[[0-9;]*m//g'
