#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

sudo apt install sshfs

mkdir -p ~/Remote/sshfs_dev

sshfs -p ${PORT} ${USER}@${HOSTNAME}:/home/${USER}/work ~/Remote/sshfs_dev

fusermount -u ~/Remote/sshfs_dev
