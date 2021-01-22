#! /bin/bash

__author__=d2jvkpn
__version__=2.1.2
__release__=2019-02-21
__project__=https://github.com/d2jvkpn/dksh/

USAGE="Start a container workspace, usage:
  $(basename $0)  [OPTIONS]  <docker_image>

  Options:
    -n  <string>  set the container name, E.g. hello, if not, a random string
      (16 chars) will be assigned.

    -v  <quoted string>  mount directories to container, e.g. a:/x/a:ro, a::ro, a

    -w  <string>  set work directory replace current directory

    -u  <string>  set login user for container, default: $USER
      note: without mount any directry and set environment variable by set
      empty string, namely -u ""

    -U  <string>  set login exisiting user in container

    -s  <string>  set shell, default is bash

    -g  <string>  create additional groups and add user to groups

    -d <qutoted string> additional docker arguments

    -c <int> set max cpu to use

    -m <string> set max memory, followed by a suffix of m, g.
"

if [ $# -eq 0 ] ||  [[ "$1" == "-h" ]] || [[ "$1" == "--help" ]]; then
    echo "$USAGE"
    exit
fi

####
StartTime=$(date +"%Y-%m-%d %H:%M:%S %z")

n=0; SetMemory=""; SetCPU=""

while getopts "v:n:w:u:U:s:d:g:m:c:" arg; do
    case $arg in
        v) mdir="$mdir $OPTARG"; n=$((n+1));;
        n) CONTAINER="$OPTARG"; n=$((n+1));;
        w) HostPath="$OPTARG"; n=$((n+1));;
        u) user="$OPTARG"; n=$((n+1));;
        U) euser="$OPTARG"; n=$((n+1));;
        s) Shell="$OPTARG"; n=$((n+1));;
        d) dkopts="$OPTARG"; n=$((n+1));;
        g) group="$group $OPTARG"; n=$((n+1));;
        m) SetMemory="-m $OPTARG"; n=$((n+1));;
        c) SetCPU="--cpu-period 100000 --cpu-quota $((100000*OPTARG))"; n=$((n+1));;
    esac
done

for i in $(seq $n); do shift; shift; done

####
if [ $# -ne 1 ]; then
    echo "$USAGE"
    exit
fi

####
img=$1

IMGAGE="$(docker images | sed 's/   */\t/g' |
awk -v i=$img 'NR>1{k=$1":"$2;
if(i==k || i":latest"==k || i==$3) {print $1":"$2" "$3; exit}}')"

if [ -z "$IMGAGE" ]; then
    echo "Cann't found \"$img\", available local image(s):"

    docker images | sed 's/   */\t/g' | awk 'BEGIN{FS=OFS="\t"}
    NR>1{$2=$1":"$2; $1=""; for(i=1; i<=NF; i++) $i=$i"  "; print}' |
    column -t -s $'\t'

    exit
fi

img=$(echo $IMGAGE | awk '{print $1}')

###
test -z "$HostPath" && HostPath="$(pwd)"
test -d "$HostPath" && HostPath=$(readlink -f "$HostPath") ||
{ echo "Error: directory \"$HostPath\" not exists."; exit; }

mountdir=$(for i in $mdir; do echo $i; done | sed 's/:/ /g' |
while read s d m; do
    test -d $s || { echo "Error: directory \"$s\" not exists"; exit; }
    realpath=$(readlink -f $s)

    if [ -z $d ]; then
        d=$(cd $(dirname $s) && pwd)/$(basename $s)
    elif [[ "$d" == "/*" ]]; then
        echo "Error: \"$d\" is not an absolute path"; exit
    fi

    test -z $m && m="ro"

    echo "-v $realpath:$d:$m"
done)

####
if [ -z $CONTAINER ]; then
    CONTAINER=dksh_$(printf '%X\n' $(date +%s%N))
else
    CONTAINER=dksh_$CONTAINER
fi

test -z "$(echo $CONTAINER | sed 's/[\.0-9a-zA-Z_-]//g')" ||
{ echo "Error: illegal container name \"$CONTAINER\""; exit; }

userid=$(id -u $user)
test -z $userid &&
{ echo "Error: user $user not exists"; exit; }

test "${user-set}" = set && user=$(whoami)
test -z $Shell && Shell=/bin/bash

gid=$(id -u $user)
test -z $gid && { echo "Error: not match Gid for user $user"; exit; }
gn=$(awk -F ":" -v gid=$gid '$3==gid{print $1; exit}' /etc/group)
test -z $gn && gn=$user ## group may not exists
AddUG="groupadd -g $gid $gn; useradd -u $userid -g $gn $user"

for g in $group; do
    i=$(awk -F ":" -v g=$g '$1==g{print $3; exit}' /etc/group)
    test -z $i && { echo "Error: group $g not exists"; exit; }
    AddUG="$AddUG; groupadd -g $i $g; usermod -a -G $g $user"
done

####
echo "Creating container \"$CONTAINER\"..."

if [ -z $user ]; then
    docker run --name=$CONTAINER --rm -it $img $Shell
    exit 0
elif [[ "$user" == "root" ]]; then
    cmd="cd /mnt/HostPath; $Shell"
elif [[ "$euser" != "" ]]; then
    cmd="cd /mnt/HostPath; su $euser -s $Shell"
else
    # cmd="$AddUG &> /dev/null; cd /mnt/HostPath; su $user -s $Shell"
    cmd="$AddUG; cd /mnt/HostPath; su $user -s $Shell"
fi

docker run $dkopts --name=$CONTAINER --rm -it $SetMemory $SetCPU $mountdir \
-v "$HostPath":/mnt/HostPath -e StartTime="$StartTime" \
-e PoweredBy=dksh -e CONTAINER=$CONTAINER -e HostPath="$HostPath" \
-e IMAGE="$IMGAGE" $img /bin/sh -c "$cmd"

## "su $user -s $Shell" will make sure user environment loaded without -l,
## but docker -w "workpath" is ineffective, so "cd /mnt/HostPath" firstly
