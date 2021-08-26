#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

app=$1
target=~/Apps/$app

test -d $target && rm -r $target
mkdir -p $target && cd $target
npm install -g $app

cd $wd


exit
##### set registry
npm config set registry http://registry.npm.taobao.org/

##### list project installed packages
npm list

##### list global installed packaged
npm list --global
