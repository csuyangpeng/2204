ARG dist=ubuntu
ARG tag=22.04
FROM 10.18.1.2:5000/${dist}:${tag}

MAINTAINER cnlab <csuyangpeng@gmail.com>

COPY sources.list /etc/apt/sources.list

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
        autoconf \
        automake \
        libtool \
        gcc-7 \
        g++-7 \
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
        update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-7 100 && \
        update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-7 100 && \
    apt-get clean
