#! /usr/bin/env bash
set -eu -o pipefail

docker ps -q |
  xargs docker inspect --format "{{.Name}}  {{.NetworkSettings.IPAddress}}" |
  sed '1i Name IPAddress' |
  column -t -s'  '
