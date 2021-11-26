#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)/
_path=$(dirname $0 | xargs -i readlink -f {})/

echo '{}' > package.json

npm install express ws yargs

exit
curl -i https://node-web.domain.example/api/time

node client.js --addr wss://node-web.domain.example/ws/talk
