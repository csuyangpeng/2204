#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

# cp /tmp/internal/sources.list /etc/apt/
apt-get -q update && apt-get -q -y install \
    make \
    cmake \
    git \
    curl \
    wget \
    sudo \
    gdb \
    vim \
    bc \
    libpcap-dev \
    pciutils \
    libsctp-dev \
    libboost-all-dev \
    libmysql++-dev \
    libmysqlcppconn-dev \
    libxml2-dev \
    libcrypto++-dev \
    libyaml-cpp-dev \
    libgrpc-dev \
    libgrpc++-dev \
    libprotobuf-dev \
    protobuf-compiler-grpc \
    libjsoncpp-dev \
    libglib2.0-dev \
    libreadline-dev \
    libpopt-dev \
    lshw \
    net-tools \
    iputils-ping \
    gcc \
    g++ \
    bash-completion
    # libhiredis-dev \

# mv /tmp/internal/nb/* /usr/local/bin/
# sed -i 's/SS/StreamT/g' /usr/include/yaml-cpp/traits.h
# sed -i 's/TT/ValueT/g' /usr/include/yaml-cpp/traits.h

# update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-7 110
# update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-7 110

apt-get clean && rm -rf /var/lib/apt/lists/*
