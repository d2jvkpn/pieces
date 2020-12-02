#! /usr/bin/env bash
set -eu -o pipefail

ssh-keygen
ssh-copy-id -p 22 rover@hostname_or_ip

pip3 install --upgrade -r requirements.txt

python3 daily_remote_backup.py -toml config.toml
python3 daily_remote_backup.py -toml config.toml -once true
