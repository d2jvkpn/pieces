#! /bin/bash

export AUTOSSH_LOGFILE="$0.log"
export AUTOSSH_PIDFILE="$0.pid"

## cronjob
# @reboot bash /path/to/autossh_reverse_proxy.sh

REMOTE_IP=1.2.3.4
REMOTE_SSHPort=22
REMOTE_AccessPort=2001
LOCAL_SSHPort=10022

autossh -p $REMOTE_SSHPort -f -N -R localhost:$REMOTE_AccessPort:localhost:$LOCAL_SSHPort \
  -o "ServerAliveInterval 10"   \
  -o "ServerAliveCountMax 3"    \
  -o "ExitOnForwardFailure yes" \
  root@$REMOTE_IP

exit
# on remote machine
ssh -p $REMOTE_AccessPort hello@localhost
