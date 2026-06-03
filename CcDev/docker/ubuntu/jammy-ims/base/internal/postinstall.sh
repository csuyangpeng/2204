#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

echo "Asia/Shanghai" > /etc/timezone
mv setup.sh /usr/local/bin/setup
