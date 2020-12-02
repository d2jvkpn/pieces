#! /bin/bash
#  project =https://github.com/d2jvkpn/pyds
set -eu -o pipefail

HELP='''Start jupyter using docker image:
$ Jupyter_pyds.sh  [-c command]  [-p port]  [-i image]
  default: command=lab, port=9999, image=pyds:latest

project: https://github.com/d2jvkpn/pyds'''

if [ $# -gt 0 ] && 
  ([[ "$1" == "-h" ]] || [[ "$1" == "-help" ]] || [[ "$1" == "--help" ]]); then
  echo "$HELP"
  exit 2
fi

command=lab; image=pyds:latest; port=9999

while [ $# -ge 2 ]; do
  case "$1" in
  -c) command=$2 && { shift; shift; };;
  -i) image=$2   && { shift; shift; };;
  -p) port=$2    && { shift; shift; };;
  *)  echo "Option $1 not recognized" && exit 1;;
  esac
done

docker run -it -p $port:$port -v $PWD:/mnt/HostPath $image \
  su root -l -c "cd /mnt/HostPath; \
  jupyter $command --generate-config -y; jupyter $command password; \
  jupyter $command --ip=0.0.0.0 --no-browser --allow-root --port $port"
