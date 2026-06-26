#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

echo "Asia/Shanghai" > /etc/timezone
rm -rf /tmp/internal
