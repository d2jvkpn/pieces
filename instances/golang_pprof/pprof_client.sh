#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)/
_path=$(dirname $0)/

d=$(date +"%s_%FT%T" | sed 's/:/-/g')"_pprof"
addr=http://localhost:1030
secs=30

mkdir -p $d
cd $d

# go tool pprof -http=:8080 allocs.out
# go tool pprof -png allocs.out > allocs.png
wget -O allocs.out       $addr/debug/pprof/allocs?seconds=$secs &
wget -O block.out        $addr/debug/pprof/block?seconds=$secs &
wget -O goroutine.out    $addr/debug/pprof/goroutine?seconds=$secs &
wget -O heap.out         $addr/debug/pprof/heap?seconds=$secs &
wget -O mutex.out        $addr/debug/pprof/mutex?seconds=$secs &
wget -O profile.out      $addr/debug/pprof/profile?seconds=$secs &
wget -O threadcreate.out $addr/debug/pprof/threadcreate?seconds=$secs &

# go tool trace trace.out
wget -O trace.out        $addr/debug/pprof/trace?seconds=$secs &

# go runtime
wget -O status.json      $addr/debug/runtime/status &

wait
echo "done"
