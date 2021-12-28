#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


####
go run search_in_file_v1.go eee a01.txt
go run search_in_file_v1.go "eee"$'\n' a01.txt
go run search_in_file_v1.go "eem" a01.txt
go run search_in_file_v1.go "eee"$'\n'"f" a01.txt

go run search_in_file.go eee a01.txt
go run search_in_file.go zzz a01.txt
go run search_in_file.go "eee"$'\n' a01.txt
go run search_in_file.go "eem" a01.txt
go run search_in_file.go "eee"$'\n'"f" a01.txt

go run search_in_file.go -debug eee a01.txt

go run search_in_file.go -debug eee$'\n'ff a01.txt

####
mkdir -p target
rustfmt search_in_file.rs
rustc search_in_file.rs -o target/search_in_file

APP_Debug=true target/search_in_file eee$'\n'ff a01.txt
APP_Debug=true target/search_in_file e a01.txt
APP_Debug=true target/search_in_file ee a01.txt

APP_Debug=true target/search_in_file ex a01.txt

rustc search_in_file.rs -O -o target/search_in_file # release
