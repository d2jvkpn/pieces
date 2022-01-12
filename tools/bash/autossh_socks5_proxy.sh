#! /usr/bin/env bash

## crontab -l
# @reboot bash /path/to/autossh_socks5_proxy.sh

export AUTOSSH_LOGFILE="$0.log"
export AUTOSSH_PIDFILE="$0.pid"

REMOTE_SSHPort=22
REMOTE_IP=1.2.3.4
LOCAL_Port=1080

autossh -f -N -C -D $LOCAL_Port -p $REMOTE_SSHPort \
  -o "ServerAliveInterval 5"    \
  -o "ServerAliveCountMax 2"    \
  -o "ExitOnForwardFailure yes" \
  root@$REMOTE_IP

## curl with socks5 proxy
curl -x socks5h://localhost:$LOCAL_Port https://github.com/d2jvkpn
