# BSD 2-Clause License

# Copyright (c) 2020, Supreeth Herle
# All rights reserved.

# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions are met:

# 1. Redistributions of source code must retain the above copyright notice, this
#    list of conditions and the following disclaimer.

# 2. Redistributions in binary form must reproduce the above copyright notice,
#    this list of conditions and the following disclaimer in the documentation
#    and/or other materials provided with the distribution.

# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/cc:dev-22.04

ARG MAKEFLAGS=-j TARGETARCH
ENV GOROOT=/opt/go
ENV GOPATH=/home/sder/go
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GITUSER=peng.yang GITPASS=11asdASD
# ENV http_proxy="http://10.18.11.52:7890" https_proxy="http://10.18.11.52:7890"
ENV PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}
ENV DEBIAN_FRONTEND=noninteractive

COPY internal /tmp/internal
# Install updates and dependencies
RUN apt-get update && \
	apt-get -y install git tcpdump screen tmux gperf ntp ntpdate git-core dkms \
  gcc flex bison libmysqlclient-dev make libssl-dev libcurl4-openssl-dev \
	libxml2-dev libpcre2-dev bash-completion g++ autoconf libmnl-dev \
	libsctp-dev libradcli-dev libradcli4 libjson-c-dev pkg-config iproute2 net-tools \
	iputils-ping libgcrypt20-dev libxmlrpc-core-c3-dev libjson-glib-dev libwebsockets-dev \
  libiptc-dev libavcodec-dev libavformat-dev libavutil-dev libswresample-dev \
  libspandsp-dev libhiredis-dev libbcg729-0 debhelper-compat module-assistant \
  nfs-common libbencode-perl libconfig-tiny-perl libcrypt-rijndael-perl \
  libio-socket-inet6-perl libsocket6-perl netcat-openbsd libbcg729-dev markdown \
  libcrypt-openssl-rsa-perl libdigest-crc-perl libio-multiplex-perl libnet-interface-perl \
  libsystemd-dev libxtables-dev markdown libavfilter-dev libdigest-hmac-perl \
   \
  git vim tmux dpkg-dev debhelper libxtables-dev default-libmysqlclient-dev gperf libavcodec-dev libavfilter-dev libavformat-dev \
  libavutil-dev libbencode-perl libcrypt-openssl-rsa-perl libcrypt-rijndael-perl libdigest-crc-perl libdigest-hmac-perl \
  libevent-dev libhiredis-dev libio-multiplex-perl libio-socket-inet6-perl libiptc-dev libjson-glib-dev libnet-interface-perl \
  libpcap0.8-dev libpcre3-dev libsocket6-perl libspandsp-dev libssl-dev libswresample-dev libsystemd-dev libxmlrpc-core-c3-dev \
  markdown dkms module-assistant keyutils libnfsidmap1 nfs-common rpcbind libglib2.0-dev zlib1g-dev libavcodec-extra \
  libcurl4-openssl-dev netcat-openbsd netcat iptables libxtables12 iproute2 net-tools iputils-ping libconfig-tiny-perl libwebsockets-dev \
  libopus-dev libnftnl-dev pandoc \
  && apt-get clean && rm -rf /var/lib/apt/lists/*
  # netcat libnfsidmap2

RUN /tmp/internal/install_golang.sh

# Install rtpengine
# dpkg-checkbuilddeps && dpkg-buildpackage -b -uc -us
ENV http_proxy= https_proxy=
RUN git clone http://${GITUSER}:${GITPASS}@10.18.1.2:9999/develop/ims.git && \
	cd ims && git checkout cc_yp && cd rtpengine && dpkg-checkbuilddeps && \
	dpkg-buildpackage -b -uc -us && cd .. && \
	dpkg -i *.deb && ldconfig && cd .. && rm ims -rf

# # Fetch Kamailio code
# RUN mkdir -p /usr/local/src/ && cd /usr/local/src/ && \
# 	git clone https://github.com/kamailio/kamailio && \
# 	cd kamailio && git checkout 4fb8accc6747ad56fec3dc84d70cb2b8bbd7316e

# # Build and Install Kamailio
# RUN cd /usr/local/src/kamailio && make cfg
# COPY modules.lst /usr/local/src/kamailio/src
# RUN cd /usr/local/src/kamailio && \
# 	make -j`nproc` Q=0 all | tee make_all.txt && \
# 	make install | tee make_install.txt && \
# 	ldconfig && \
# 	cd / && rm -rf /usr/local/src/kamailio

ENV GITUSER= GITPASS= 
ENV http_proxy= https_proxy=

# COPY kamailio_init.sh /
# CMD /kamailio_init.sh
