#! /usr/bin/env bash
set -eu -o pipefail

# reference: https://jupyter-notebook.readthedocs.io/en/stable/public_server.html

# Create a self-signed certificate can be generated with openssl

# $ openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
#   -keyout configs/jupyter.key -out configs/jupyter.pem

# $ jupyter notebook --certfile=configs/jupyter.pem --keyfile configs/jupyter.key


## Setup
echo ">>> setup"

read -p "jupyter port (9100): " port
test -z "$port" && port="9100"

read -s -p "enter password: " password; echo ""
test -z "$password" && { echo "empty password"; exit 1; }
read -s -p "verify password: " password2; echo ""
test -z "$password2" && { echo "empty password"; exit 1; }
test "$password" == "$password2" || { echo "password not match"; exit 1; }

read -p "container name (pyspark-ws1): " container
test -z "$container" && container="pyspark-ws1"

read -p "workpath (./):  " workpath
test -z "$workpath" && workpath="./"
test -d $workpath || { echo "workpath not exists"; exit 1; }
workpath=$(readlink -f $workpath)

## Run container
echo ">>> run container $container"

# --restart=always
docker run --detach --publish=$port:$port --name="${container}" \
    --volume="${workpath}":/mnt/WorkPath --workdir=/mnt/WorkPath \
    pyspark:latest python3 /root/configs/jupyter.py "$port" "$password"
