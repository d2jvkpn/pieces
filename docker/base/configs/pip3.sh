#! /usr/bin/env bash
set -eu -o pipefail

pip3 install --no-cache --upgrade pip
pip3 install --no-cache -r ~/configs/requirements.txt
python3 -m bash_kernel.install
jupyter kernelspec list
jupyter kernelspec uninstall unwanted-kernel
