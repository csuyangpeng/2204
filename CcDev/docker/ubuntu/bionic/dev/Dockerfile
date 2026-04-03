ARG dist=epc
ARG tag=base18
FROM ${dist}:${tag}

MAINTAINER cnlab <csuyangpeng@gmail.com>

RUN apt-get update && \
    apt-get upgrade -y && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y --no-install-recommends \
    git-buildpackage \
    debhelper \
    devscripts \
    osc \
    dput \
    dh-autoreconf \
    dh-systemd \
    manpages-dev \
    fakeroot \
    cscope \
    vim \
    sudo \
    unzip \
    iputils-ping \
    wireshark \
    curl \
    pciutils \
    lshw \
    gdb \
    libevent-dev \
    cmake \
    net-tools && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get clean

RUN cp /bin/rm /usr/bin/rm
COPY internal/setup.sh /usr/local/bin/setup
COPY internal /tmp/internal
RUN bash /tmp/internal/install_ipsec.sh
RUN bash /tmp/internal/install_kernel.sh
RUN bash /tmp/internal/install_epc.sh

ARG username=root
RUN mkdir -p /home/${username}/epc
WORKDIR /home/${username}/epc
ENV PATH ${PATH}:/usr/local/epc/install/bin:/usr/local/ipsec/install/sbin
RUN ldconfig
