#!/usr/bin/env bash

set -e

ulimit -c unlimited
echo "core-%e-%p-$(date +%H:%M:%S)" > /proc/sys/kernel/core_pattern
