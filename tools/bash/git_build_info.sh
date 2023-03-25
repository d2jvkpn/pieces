#! /usr/bin/env bash
# set -eu -o pipefail
# _wd=$(pwd)
# _path=$(dirname $0 | xargs -i readlink -f {})

BUILD_Time=$(date +'%FT%T%:z')
GIT_Branch=$(git rev-parse --abbrev-ref HEAD)

GIT_Commit=$(git rev-parse --verify HEAD) # git log --pretty=format:'%h' -n 1
GIT_Time=$(git log -1 --format="%at" | xargs -I{} date -d @{} +%FT%T%:z)
GIT_TreeState="clean"

uncommitted=$(git status --short)
unpushed=$(git diff origin/$GIT_Branch..HEAD --name-status)
[ ! -z "$uncommitted$unpushed" ] && GIT_TreeState="dirty"

# src/git-build-info.yaml, .git-build-info.yaml
cat <<EOF
build_time: $BUILD_Time
git_branch: $GIT_Branch
git_commit: $GIT_Commit
git_time: $GIT_Time
git_tree_state: $GIT_TreeState
EOF
