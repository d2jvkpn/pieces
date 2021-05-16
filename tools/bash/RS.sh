#! /usr/bin/env bash
set -eu -o pipefail

script=$1

rustfmt $script && rustc $script && ./${script%.rs}
