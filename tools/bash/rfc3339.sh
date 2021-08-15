#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)


date --rfc-3339=seconds | sed "s/ /T/"
