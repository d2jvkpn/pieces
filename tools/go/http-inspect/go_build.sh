#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

prog=$(basename ${_path})

mkdir -p target
go build -ldflags="-w -s" -o target/$prog main.go
GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o target/$prog.exe main.go
ls -al target
