#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})


mod=$1
mod=${mod%/}
d=$(basename $mod)

mkdir -p $d
cd $d

go mod init $mod

cat > main.go << EOF
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}
EOF

[[ -d ".git" ]] && exit 0
git init
git remote add origin https://$mod

cat > .gitignore << EOF
logs/
configs/
wk_*/
main
EOF

echo "static/" > .tokeignore

exit 0
git config user.name "YourName"
git config user.email "YourEmail"
git add .
git commit -m 'init'
git push -u origin master
