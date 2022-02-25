#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

name=$1

mkdir -p $name

cp ~/Templates/sh.sh $name/do.sh
cp ~/Templates/md.md $name/think.md

echo "cd $name"
