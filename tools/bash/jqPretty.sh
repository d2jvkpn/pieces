#! /usr/bin/env bash
set -eu -o pipefail

cat <<< $(jq . $1) > $1
