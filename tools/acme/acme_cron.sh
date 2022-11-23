#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {} )

# cron: 0 0 * * *
# location: ${HOME}/.acme.sh/${DOMAIN}

domain=YOUR_Domain.xyz
target=$HOME/nginx/cert

acme_home=$HOME/.acme.sh # directory
domain_home=$acme_home/$domain

{
    date +">>> %FT%T%:z run acme.sh"
    s1=$(md5sum $domain_home/$domain.cer | awk '{print $1}')

    ${acme_home}/acme.sh --cron --home $acme_home --server letsencrypt
    s2=$(md5sum $domain_home/$domain.cer | awk '{print $1}')

    if [[ "$s1" != "$s2" ]]; then
        date +"    %FT%T%:z renew ssl and reload nginx"
        rsync ${domain_home}/$domain.{key,cer} $target/
        nginx -s reload
        # sudo nginx -s reload
    fi
} >> ${_path}/acme.$(date +"%Y").log 2>&1
