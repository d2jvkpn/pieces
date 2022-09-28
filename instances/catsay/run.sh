#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


cargo add colored predicates structopt
cargo add --dev assert_cmd predicates

cargo build

cargo run -- "Hello, World!"
cargo run -- "Hello, World!" -d
cargo run -- "Hello, World!" -d -f cat_file.txt

cargo test
cargo test tests::run_with_defaults -- --exact
cargo test run_with_defaults
cargo test with
