#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

#### install, https://github.com/acmesh-official/acme.sh
curl https://get.acme.sh  | sh

access_key_file=ALIYUN_access_key.json
email=YOUR_Email@domain.xyz
domain=YOUR_Domain.xyz

#### get or generate Ali_Key and Ali_Secret from https://ram.console.aliyun.com/manage/ak
export Ali_Key="$(jq -r '.AccessKeyId' $access_key_file)"
export Ali_Secret="$(jq -r '.AccessKeySecret' $access_key_file)"

#### register account
export PATH=$HOME/.acme.sh/$PATH

acme.sh --register-account -m $email
acme.sh --issue --dns dns_ali --server letsencrypt -d $domain -d *.$domain

#### edit cron jobs
mkdir -p ~/cron && cp acme_cron.sh ~/cron

# $ crontabl -e
# remove default and add following line
# 0 0 * * * /path-to-your-home/cron/acme_cron.sh
