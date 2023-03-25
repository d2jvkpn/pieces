#! /bin/bash
set -eu -o pipefail
_wd=$(pwd)
_self=$(readlink -f $0)
_path=$(dirname $0 | xargs -i readlink -f {})

## cronjob
# @reboot bash /path/to/ssh_remote_forwarding.sh

# export AUTOSSH_LOGFILE="${_self}.$(date +%F_%s).log"
export AUTOSSH_LOGFILE=$(readlink -f $0).$(date +%Y-%m).log
export AUTOSSH_PIDFILE="${_self}.pid"

REMOTE_User=hello
REMOTE_IP=1.2.3.4
REMOTE_SSHPort=22
REMOTE_Port=2001
LOCAL_Port=10022

autossh -fN -R localhost:$REMOTE_Port:localhost:$LOCAL_Port -p $REMOTE_SSHPort \
  -o "ServerAliveInterval 5"   \
  -o "ServerAliveCountMax 2"    \
  -o "ExitOnForwardFailure yes" \
  $REMOTE_User@$REMOTE_IP

exit

# on remote machine
ssh -p $REMOTE_Port hello@localhost
