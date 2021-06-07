#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

app=$1
mkdir -p ~/Apps/$app
cd ~/Apps/$app

npm install --global live-server || rm -r ~/Apps/$app
cd $wd

exit

##### set registry
npm config set registry http://registry.npm.taobao.org/

##### list project installed packages 
npm list

##### list global installed packaged 
npm list --global
