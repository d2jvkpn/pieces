#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_self=$(readlink -f $0)
_path=$(dirname ${_self})

rustup run nightly cargo run
rustup run nightly cargo build --release
