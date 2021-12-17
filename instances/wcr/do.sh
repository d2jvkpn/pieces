#! /usr/bin/env bash
set -eu -o pipefail

_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


####
mkdir -p tests/{inputs,expected}

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

wc tests/inputs/fox.txt > tests/expected/fox.txt.out
wc -c tests/inputs/fox.txt > tests/expected/fox.txt.c.out

####
cargo test --tests tests::test_format_field

cargo test -- tests::test_format_field --exact

echo -e "hello, world" | cargo run -- tests/inputs/* -
