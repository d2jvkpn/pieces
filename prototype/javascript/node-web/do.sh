#! /usr/bin/env bash
set -eu -o pipefail


echo '{}' > package.json

npm install express ws
