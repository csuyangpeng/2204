#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

# apt 走代理访问 archive.ubuntu.com 时大量并发下载易触发 502，改用国内镜像并直连
unset http_proxy https_proxy HTTP_PROXY HTTPS_PROXY
rm -f /etc/apt/sources.list.d/ubuntu.sources
cp /tmp/internal/sources.list /etc/apt/sources.list
echo 'Acquire::Retries "5";' > /etc/apt/apt.conf.d/80-retries
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
