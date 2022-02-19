#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

host=$1

# -vv -vvv
ansible-playbook -v ${_path}/apt-update.yaml --inventory=hosts.ini --extra-vars="host=$host"
