#! /usr/bin/env bash
set -eu -o pipefail

# Count Lines of Code

suffix="go"
test $# -gt 0 && suffix="${1}" || true

find -name "*.${suffix}"  | sed '/^[[:space:]]*$/d' | xargs -i cat {} | strings | wc -l
