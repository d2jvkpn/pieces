#! /usr/bin/env bash
set -eu -o pipefail

yum -y install nodejs npm

npm config set registry https://registry.npm.taobao.org
npm install -g nodemon

npm config list
