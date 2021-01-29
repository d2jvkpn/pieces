#! /usr/bin/env bash
set -eu -o pipefail

ps -eo pid,ppid,cmd,%mem,%cpu --sort=-%mem | head -n 10
