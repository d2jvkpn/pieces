#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)


branch=$(git rev-parse --abbrev-ref HEAD)

git add -A # include untracked files for git diff, undo with git reset && git checkout .
test -z "$(git diff HEAD)" || { echo "You have uncommitted changes!"; exit 1; }

test -z "$(git diff origin/$branch..HEAD --name-status)" ||
  { echo "You have unpushed commits!"; exit 1; }
