#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

# cp /tmp/internal/sources.list /etc/apt/
apt-get -q update && apt-get -q -y install \
    git \
    make \
    cmake \
    gcc-4.8 \
    g++-4.8 \
    libevent-dev \
    libspeex-dev \
    libopus-dev \
    libssl-dev \
    python-dev \
    libhiredis-dev \
    libcurl4-openssl-dev \
    wget

update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.8 110
update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-4.8 110

apt-get clean && rm -rf /var/lib/apt/lists/*
