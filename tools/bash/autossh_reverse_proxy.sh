#! /bin/bash

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

## cronjob
# @reboot bash /path/to/autossh_reverse_proxy.sh

export AUTOSSH_LOGFILE="$0.log"

REMOTE_IP=1.2.3.4
REMOTE_SSHPort=22
REMOTE_AccessPort=2001
LOCAL_SSHPort=10022

autossh -p $REMOTE_SSHPort -f -N -R localhost:$REMOTE_AccessPort:localhost:$LOCAL_SSHPort \
  -o "ServerAliveInterval 5"   \
  -o "ServerAliveCountMax 2"    \
  -o "ExitOnForwardFailure yes" \
  root@$REMOTE_IP

exit
# on remote machine
ssh -p $REMOTE_AccessPort hello@localhost
