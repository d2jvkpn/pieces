#! /usr/bin/env bash
set -eu -o pipefail

test $# -eq 0 && { echo "Please provide a key word for searching!"; exit 1; }

key=$1; suffix="go"
test $# -gt 1 && suffix=$2
path="./"
test $# -gt 2 && path=$3

n=$(find "$path" -type f -name "*.$suffix" 2> /dev/null | head -n1 || true)
test -z "$n" && { echo "no *.$suffix file found!"; exit 1; }

grep --color=always -n "$key" $(find "$path" -type f -name "*.$suffix") |
  sed 's/\t/    /g'          |                   ## replcace tab in code with 4 spaces
  sed '1i #FILE\tLINE\tCODE' |                   ## add column names
  sed 's/:/\t\[/; s/:/\]\t/' |                   ## filename and linenumber
  column -s $'\t' -t         |                   ## cloumn align text output
  sed "1i  ~~~~~~ Found \"$key\" ~~~~~~"  ## add title
