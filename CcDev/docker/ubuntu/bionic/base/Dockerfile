ARG dist=ubuntu
ARG tag=18.04
FROM ${dist}:${tag}

MAINTAINER cnlab <csuyangpeng@gmail.com>

COPY sources.list /etc/apt/sources.list

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
        autoconf \
        automake \
        libtool \
        gcc \
        flex \
        bison \
        git \
        libsctp-dev \
        libgnutls28-dev \
        libgcrypt-dev \
        libssl-dev \
        libidn11-dev \
        libmongoc-dev \
        libbson-dev \
        libyaml-dev \
        libncurses5-dev \
        libnuma-dev \
        libpcap-dev \
        build-essential \
        iproute2 \
        ca-certificates \
        netbase \
        pkg-config && \
    apt-get clean
