#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"
cp sources.list /etc/apt/sources.list

apt-get -q update && apt-get -q -y install \
    gcc \
    make \
    sudo \
    git \
    curl \
    tcpdump \
    psmisc \
    iptables \
    net-tools \
    pkg-config \
    libgmp-dev \
    libpcap-dev \
    libelf-dev \
    hugepages  \
    libnuma-dev \
    libhyperscan-dev \
    liblua5.3-dev \
    python \
    pciutils \
    libmnl-dev \
    libibverbs-dev \
    libsctp-dev \
    gcc-aarch64-linux-gnu && \
    rm -rf /var/lib/apt/lists/*

apt-get clean
