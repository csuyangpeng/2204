#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

# cp /tmp/internal/sources.list /etc/apt/
apt-get -q update && apt-get -q -y install \
    make \
    cmake \
    autoconf \
    libtool \
    git \
    curl \
    wget \
    sudo \
    vim \
    libpcap-dev \
    pciutils \
    libsctp-dev \
    libmaxminddb-dev \
    swig \
    net-tools \
    iputils-ping \
    iputils-tracepath \
    bash-completion \
    gcc \
    g++

# update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-7 100
# update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-7 100

apt-get clean && rm -rf /var/lib/apt/lists/*
