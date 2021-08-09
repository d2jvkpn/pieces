#! /usr/bin/env bash
set -eu -o pipefail

echo ">>> Sign up"
curl -X POST http://localhost:8000/signup \
-d '{"username": "johndoe", "password": "mysecurepassword"}'
echo

echo ">>> Sign in"
curl -X POST http://localhost:8000/signin \
-d '{"username": "johndoe", "password": "mysecurepassword"}'
echo

echo ">>> Sign in with wrong password"
curl -X POST http://localhost:8000/signin \
-d '{"username": "johndoe","password": "incorrect"}'
echo


echo ">>> Sign up again"
curl -X POST http://localhost:8000/signup \
-d '{"username": "johndoe", "password": "mysecurepassword"}'
echo
