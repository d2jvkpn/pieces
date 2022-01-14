#! /usr/bin/env bash

_wd=$(pwd)
_self=$(readlink -f $0)
_path=$(dirname ${_self})

## crontab -l
# @reboot bash /path/to/autossh_socks5_proxy.sh

export AUTOSSH_LOGFILE="${_self}.log"
export AUTOSSH_PIDFILE="${_self}.pid"

REMOTE_User=hello
REMOTE_SSHPort=22
REMOTE_IP=1.2.3.4
LOCAL_Port=1080

autossh -f -N -C -D $LOCAL_Port -p $REMOTE_SSHPort \
  -o "ServerAliveInterval 5"    \
  -o "ServerAliveCountMax 2"    \
  -o "ExitOnForwardFailure yes" \
  $REMOTE_User@$REMOTE_IP

## curl with socks5 proxy
curl -x socks5h://localhost:$LOCAL_Port https://github.com/d2jvkpn
