#! /usr/bin/env bash
set -eu -o pipefail

script=$1
prog=$(basename $script | sed 's/.rs$//')
shift
args=$*

rustfmt $script && rustc -o /tmp/$prog $script

/tmp/$prog $args || true
rm /tmp/$prog
