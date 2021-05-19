#! /usr/bin/env bash
set -eu -o pipefail

script=$1
prog=$(basename $script | sed 's/.rs$//')

rustfmt $script && \
rustc -o /tmp/$prog $script && \
/tmp/$prog && \
rm /tmp/$prog
