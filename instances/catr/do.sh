#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

####
mkdir -p tests/inputs/
touch tests/inputs/empty.txt

echo "The quick brown fox jumps over the lazy dog." > tests/inputs/fox.txt

cat > tests/inputs/spiders.txt << EOF
Don't worry, spiders,
I keep house
casually.
EOF

cat > tests/inputs/the-bustle.txt << EOF
The bustle in a house
The morning after death
Is solemnest of industries
Enacted upon earth,—

The sweeping up the heart,
And putting love away
We shall not want to use again
Until eternity.
EOF

####
cargo add clap flate2

cargo add --dev assert_cmd predicates rand
