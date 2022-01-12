#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

python3 -c "import time; print('start'); time.sleep(10); print('done')"
