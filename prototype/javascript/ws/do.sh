#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

echo '{}' > package.json

npm install express ws
