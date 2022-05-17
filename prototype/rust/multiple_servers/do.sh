#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


cargo run

curl -i localhost:8080/api/v1/greet?name=Rover
curl -i localhost:8080/api/v1/greet/Rover
curl -i localhost:8080/api/v1/one
curl -i localhost:8080/hello/X-Men

curl -i -X POST localhost:8080/login/h5?timestamp=123 -H 'X-Language: English'


curl -i -X POST localhost:8080/auth/logout

curl -i localhost:8081/health
