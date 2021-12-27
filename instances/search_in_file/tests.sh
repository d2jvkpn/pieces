#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


go run search_in_file_v1.go eee a01.txt
go run search_in_file_v1.go "eee"$'\n' a01.txt
go run search_in_file_v1.go "eem" a01.txt
go run search_in_file_v1.go "eee"$'\n'"f" a01.txt

go run search_in_file.go eee a01.txt
go run search_in_file.go "eee"$'\n' a01.txt
go run search_in_file.go "eem" a01.txt
go run search_in_file.go "eee"$'\n'"f" a01.txt

go run search_in_file.go -debug eee a01.txt

go run search_in_file.go -debug eee$'\n'ff a01.txt
