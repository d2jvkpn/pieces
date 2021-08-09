#! /usr/bin/env bash
set -eu -o pipefail

go build auth2_json.go
./auth2_json &
p1=$!
sleep 5

echo ">>> Sign up"
curl -X POST http://localhost:9000/signup \
-d '{"username": "johndoe", "password": "mysecurepassword"}'
echo

echo ">>> Log in"
curl -X POST http://localhost:9000/login \
-d '{"username": "johndoe", "password": "mysecurepassword"}'
echo

echo ">>> Log in with wrong password"
curl -X POST http://localhost:9000/login \
-d '{"username": "johndoe","password": "incorrect"}'
echo


echo ">>> Sign up again"
curl -X POST http://localhost:9000/signup \
-d '{"username": "johndoe", "password": "mysecurepassword"}'
echo

sleep 10
kill $p1
rm ./auth2_json
