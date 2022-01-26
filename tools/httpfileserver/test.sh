#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_self=$(readlink -f $0)
_path=$(dirname ${_self})


mkdir -p target/tmp
echo "hello, world!" > target/tmp/hello.txt
truncate -s 1G target/tmp/a.data

go build -o target/httpfileserver
GOOS=windows GOARCH=amd64 go build -o target/httpfileserver.exe

target/httpfileserver -path target/tmp -address :8000

curl localhost:8000
wget -O target/hello.txt http://localhost:8000/hello.txt
