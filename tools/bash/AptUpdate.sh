#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

apt update && apt -y upgrade
apt clean && apt autoclean
apt remove && apt autoremove

dpkg -l | awk '/^rc/{print $2}' | xargs -i dpkg -P {}
