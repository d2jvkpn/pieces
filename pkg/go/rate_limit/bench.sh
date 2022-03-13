#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

mkdir -p wk_v1 wk_v2

go test  -run none  -bench ^BenchmarkLimiterV1_b1$ -outputdir wk_v1 \
  -cpuprofile=cpu.out -memprofile=mem.out -blockprofile=block.out

go test  -run none  -bench ^BenchmarkLimiterV2_b1$ -outputdir wk_v2 \
  -cpuprofile=cpu.out -memprofile=mem.out -blockprofile=block.out

for f in wk*/*.out; do
    go tool pprof -svg $f > ${f%.out}.svg
done
