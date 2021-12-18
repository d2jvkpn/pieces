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

cargo test --test parse_positive_int

cargo test test_xx # runs test_xx_1 and test_xx_2

cargo test --test test_xx_1 # run test_xx_1 only

cargo test tests::test_xx_1 -- --exact # specify module

cargo test test_xx_2 -- --exact # specify module
