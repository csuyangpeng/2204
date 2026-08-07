# ---------- builder：编译安装，层不进入最终镜像 ----------
FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/cc:dev-26.04 AS builder

ARG TARGETARCH
ENV GOROOT=/opt/go
ENV GOPATH=/home/sder/go
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}
ENV DEBIAN_FRONTEND=noninteractive

# Install build dependencies（系统 FFmpeg 原生支持 rtpengine 12.4）
RUN apt-get update && \
	apt-get -y install wget git unzip tcpdump screen tmux gperf ntpsec ntpsec-ntpdate git-core dkms \
  gcc flex bison libmysqlclient-dev make libssl-dev libcurl4-openssl-dev bzip2 \
	libxml2-dev libpcre2-dev bash-completion g++ autoconf libmnl-dev \
	libsctp-dev libradcli-dev libradcli4 libjson-c-dev pkg-config iproute2 net-tools \
	iputils-ping libgcrypt20-dev libxmlrpc-core-c3-dev libjson-glib-dev libwebsockets-dev \
  libiptc-dev libspandsp-dev libhiredis-dev libbcg729-0 debhelper-compat module-assistant \
  nfs-common libbencode-perl libconfig-tiny-perl libcrypt-rijndael-perl \
  libio-socket-inet6-perl libsocket6-perl netcat-openbsd libbcg729-dev markdown \
  libcrypt-openssl-rsa-perl libdigest-crc-perl libio-multiplex-perl libnet-interface-perl \
  libsystemd-dev libxtables-dev markdown libdigest-hmac-perl \
  git vim tmux dpkg-dev debhelper libxtables-dev default-libmysqlclient-dev gperf \
  libevent-dev libhiredis-dev libio-multiplex-perl libio-socket-inet6-perl libiptc-dev libjson-glib-dev libnet-interface-perl \
  libpcap0.8-dev libsocket6-perl libspandsp-dev libssl-dev libsystemd-dev libxmlrpc-core-c3-dev \
  dkms module-assistant keyutils libnfsidmap1 nfs-common rpcbind libglib2.0-dev zlib1g-dev \
  libcurl4-openssl-dev netcat-openbsd iptables libxtables12 iproute2 net-tools iputils-ping libconfig-tiny-perl libwebsockets-dev \
  libopus-dev libnftnl-dev pandoc liblua5.1-0-dev libmicrohttpd-dev lsb-release \
  clang llvm libbpf-dev libelf-dev linux-libc-dev \
  libavcodec-dev libavfilter-dev libavformat-dev libavutil-dev libswresample-dev \
  && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY internal/install_golang.sh /tmp/internal/
RUN bash /tmp/internal/install_golang.sh

# ETSI EVS codec（rtpengine 运行时依赖）
ENV http_proxy=http://10.18.11.52:7890 https_proxy=http://10.18.11.52:7890
RUN wget https://www.etsi.org/deliver/etsi_ts/126400_126499/126443/17.00.00_60/ts_126443v170000p0.zip && \
	unzip ts_126443v170000p0.zip && \
  wget https://raw.githubusercontent.com/sipwise/rtpengine/refs/heads/mr12.4.1/evs/etsi-ts-126.443-3gpp-ts-26.443-v17.0.0/patch && \
	cd c-code && \
  patch -p1 < ../patch && \
	make && \
  ls && cp lib3gpp-evs.so /usr/local/lib/ && \
	cd .. && rm ts_126443v170000p0.zip patch c-code -rf

# ---------- 最终镜像：只保留运行时产物 ----------
FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/cc:dev-26.04

ARG TARGETARCH
ENV GOROOT=/opt/go
ENV GOPATH=/home/sder/go
ENV GO111MODULE=on
ENV GOPROXY="http://10.18.1.2:8081,https://goproxy.cn,direct" \
    GOSUMDB="sum.golang.org http://10.18.1.2:8081/sumdb/sum.golang.org"
ENV PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}
ENV DEBIAN_FRONTEND=noninteractive

# Runtime dependencies（系统 FFmpeg 原生支持 rtpengine 12.4）
RUN apt-get update && \
	apt-get -y install wget git unzip tcpdump screen tmux gperf ntpsec ntpsec-ntpdate git-core dkms \
  gcc flex bison libmysqlclient-dev make libssl-dev libcurl4-openssl-dev bzip2 \
	libxml2-dev libpcre2-dev bash-completion g++ autoconf libmnl-dev \
	libsctp-dev libradcli-dev libradcli4 libjson-c-dev pkg-config iproute2 net-tools \
	iputils-ping libgcrypt20-dev libxmlrpc-core-c3-dev libjson-glib-dev libwebsockets-dev \
  libiptc-dev libspandsp-dev libhiredis-dev libbcg729-0 debhelper-compat module-assistant \
  nfs-common libbencode-perl libconfig-tiny-perl libcrypt-rijndael-perl \
  libio-socket-inet6-perl libsocket6-perl netcat-openbsd libbcg729-dev markdown \
  libcrypt-openssl-rsa-perl libdigest-crc-perl libio-multiplex-perl libnet-interface-perl \
  libsystemd-dev libxtables-dev markdown libdigest-hmac-perl \
  git vim tmux dpkg-dev debhelper libxtables-dev default-libmysqlclient-dev gperf \
  libevent-dev libhiredis-dev libio-multiplex-perl libio-socket-inet6-perl libiptc-dev libjson-glib-dev libnet-interface-perl \
  libpcap0.8-dev libsocket6-perl libspandsp-dev libssl-dev libsystemd-dev libxmlrpc-core-c3-dev \
  dkms module-assistant keyutils libnfsidmap1 nfs-common rpcbind libglib2.0-dev zlib1g-dev \
  libcurl4-openssl-dev netcat-openbsd iptables libxtables12 iproute2 net-tools iputils-ping libconfig-tiny-perl libwebsockets-dev \
  libopus-dev libnftnl-dev pandoc liblua5.1-0-dev libmicrohttpd-dev lsb-release \
  clang llvm libbpf-dev libelf-dev linux-libc-dev \
  libavcodec-dev libavfilter-dev libavformat-dev libavutil-dev libswresample-dev \
  && apt-get clean && rm -rf /var/lib/apt/lists/*

# 从 builder 复制编译产物
COPY --from=builder /opt/go /opt/go
COPY --from=builder /usr/local /usr/local

RUN ldconfig 2>/dev/null || true

USER root
WORKDIR /home/root/
