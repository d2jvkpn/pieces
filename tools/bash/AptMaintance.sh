#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

apt update && apt -y upgrade
# reboot
apt clean && apt -y autoclean
apt remove && apt -y autoremove


dpkg -l | awk '/^rc/{print $2}' | xargs -i dpkg -P {}
