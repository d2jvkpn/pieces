#! /usr/bin/env bash
set -eu -o pipefail

# export HISTTIMEFORMAT="%Y-%m-%dT%H:%M:%S%z "
# export HISTSIZE=1000000

out=history_$(date +"%F_%s").txt
history > $out
