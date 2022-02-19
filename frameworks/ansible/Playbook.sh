#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

script=$1
host=$2

# -vv -vvv
ansible-playbook -v ${_path}/$script.yaml --inventory=hosts.ini --extra-vars="host=$host"
