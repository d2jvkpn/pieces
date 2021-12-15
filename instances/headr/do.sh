#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


####
cargo add clap
cargo add --dev assert_cmd predicates rand

####
mkdir -p tests/inputs

cat > tests/inputs/three.txt << EOF
Three
lines,
four words.
EOF

touch tests/inputs/empty.txt

####
cargo test parse_positive_int
