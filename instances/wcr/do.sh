#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

mkdir -p tests/inputs

touch tests/inputs/empty.txt

echo "The quick brown fox jumps over the lazy dog." > tests/inputs/fox.txt

echo "hello, world!" > tests/inputs/hello.en.txt
echo "你好, 世界!" > tests/inputs/hello.cn.txt

cat > tests/inputs/atlamal.txt << EOF
Frétt hefir öld óvu, þá er endr of gerðu
seggir samkundu, sú var nýt fæstum,
æxtu einmæli, yggr var þeim síðan
ok it sama sonum Gjúka, er váru sannráðnir.
EOF
