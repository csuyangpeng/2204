#!/usr/bin/env bash

set -e

ulimit -c unlimited
echo "/usr/local/epc/install/var/corefile/core-%e-%p-%t" > /proc/sys/kernel/core_pattern

