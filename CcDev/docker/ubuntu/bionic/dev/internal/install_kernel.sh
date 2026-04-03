#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

tar -xzvf 4.15.0-128-generic.tar.gz && mkdir -p /lib/modules/ && mv 4.15.0-128-generic /lib/modules/ \
    && tar -xzvf linux-headers-4.15.0-128.tar.gz && mv linux-headers-4.15.0-128 /usr/src/ \
    && tar -xzvf linux-headers-4.15.0-128-generic.tar.gz && mv linux-headers-4.15.0-128-generic /usr/src/ \
    && rm 4.15.0-*.tar.gz  linux-headers-4.15.0-*.tar.gz
