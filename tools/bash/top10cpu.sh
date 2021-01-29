#! /usr/bin/env bash
set -eu -o pipefail

ps -eo pid,ppid,cmd,%mem,%cpu --sort=-%cpu | head -n 10
