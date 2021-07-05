#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

app=$1
mkdir -p ~/Apps/$app && cd ~/Apps/$app
npm install -g $app
cd $wd
