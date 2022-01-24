#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_self=$(readlink -f $0)
_path=$(dirname ${_self})


mkdir -p target/tmp
echo "hello, world!" > target/tmp/hello.txt
truncate -s 1G target/tmp/a.data

go build -o target/httpfileserver

target/httpfileserver -path target/tmp -address :8000
