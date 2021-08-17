#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

sudo apt install sshfs

mkdir sshfs_dev

sshfs -p PORT USER@HOSTNAME:/home/USER/work $PWD/sshfs_dev

fusermount -u $PWD/sshfs_dev
