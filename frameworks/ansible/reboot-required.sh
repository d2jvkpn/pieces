#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

host=$1

ansible-playbook -vv ${_path}/reboot-required.yaml --inventory=hosts.ini --extra-vars="host=$1"

#
exit
ansible w01 -i hosts.ini -m reboot
