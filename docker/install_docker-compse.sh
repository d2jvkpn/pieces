#! /usr/bin/env bash
set -eu -o pipefail

sudo curl -L \
  "https://github.com/docker/compose/releases/download/1.28.5/docker-compose-$(uname -s)-$(uname -m)" \
  -o /usr/local/bin/docker-compose
