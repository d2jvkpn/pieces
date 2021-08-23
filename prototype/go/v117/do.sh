#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

go build --gcflags=-G=3 generic_addable.go

./generic_addable
