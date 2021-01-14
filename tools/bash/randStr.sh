#! /usr/bin/env bash
set -eu -o pipefail

#### random string
date +"%s%N" | md5sum

head -c 32 /dev/random | base64

openssl rand -hex 32

head -c 32 /dev/urandom | base64

cat /dev/urandom | head -n 10 | md5sum | cut -c 1-32

tr -cd '[:alnum:]' </dev/urandom | fold -w32 | head -n1

strings /dev/urandom | grep -o '[[:alnum:]]' | head -n 32 | tr -d '\n'; echo

< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c32; echo

</dev/urandom tr -dc a-z-0-9|head -c32;echo

cat /proc/sys/kernel/random/uuid  # uuid

< /dev/urandom tr -dc a-z| head -c ${1:-32}; echo

< /dev/urandom tr -dc A-Z | head -c ${1:-32};echo

< /dev/urandom tr -dc 0-9 | head -c ${1:-32}; echo

< /dev/urandom tr -dc 0-9-A-Z | head -c ${1:-32}; echo

< /dev/urandom tr -dc 0-9-A-Z-a-z | head -c ${1:-32}; echo

< /dev/urandom tr -dc 0-9-A-Z-a-z- | head -c ${1:-32}; echo

< /dev/urandom tr -dc 0-9-A-Z-a-z-/ | head -c ${1:-32}; echo

</dev/urandom tr -dc '12345!@#$%qwertQWERTasdfgASDFGzxcvbZXCVB' | head -c32; echo

dd if=/dev/urandom bs=1 count=32 2>/dev/null | base64 -w 0 | rev | cut -b 2- | rev

dd if=/dev/urandom bs=1 count=10 2>/dev/null | base64 -w 0 | rev | cut -b 2- | rev


# random number
head -20 /dev/urandom | cksum

array=(0 1 2 3 4 5 6 7 8 9)
num=${#array[*]} 
randnum=${array[$((RANDOM%num))]}
echo $randnum

seq 10 32 | shuf -n 1
