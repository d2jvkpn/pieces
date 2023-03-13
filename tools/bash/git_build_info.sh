#! /usr/bin/env bash
# set -eu -o pipefail
# _wd=$(pwd)
# _path=$(dirname $0 | xargs -i readlink -f {})

BUILD_Time=$(date +'%FT%T%:z')

GIT_Branch=$(git rev-parse --abbrev-ref HEAD)

uncommitted=$(git status --short)
unpushed=$(git diff origin/$GIT_Branch..HEAD --name-status)

GIT_Commit=$(git rev-parse --verify HEAD) # git log --pretty=format:'%h' -n 1
GIT_Time=$(git log -1 --format="%at" | xargs -I{} date -d @{} +%FT%T%:z)
GIT_TreeState="clean"

uncommitted=$(git status --short)
unpushed=$(git diff origin/$GIT_Branch..HEAD --name-status)
[ ! -z "$uncommitted$unpushed" ] && GIT_TreeState="dirty"

cat <<EOF
export BUILD_Time=$BUILD_Time
export GIT_Branch=$GIT_Branch
export GIT_Commit=$GIT_Commit
export GIT_Time=$GIT_Time
export GIT_TreeState=$GIT_TreeState
EOF
