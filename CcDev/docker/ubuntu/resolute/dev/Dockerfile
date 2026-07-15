# ---------- builder：编译安装，层不进入最终镜像 ----------
FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/cc:dev-26.04 AS builder

ARG TARGETARCH
ENV GOROOT=/opt/go
ENV GOPATH=/home/sder/go
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}
ENV DEBIAN_FRONTEND=noninteractive
# rtpengine 9.4.1 使用自编译 FFmpeg 4.4.2（与 22.04 一致），不用 26.04 系统 FFmpeg 8
ENV FFMPEG_PREFIX=/usr/local/ffmpeg-4.4.2
ENV PKG_CONFIG_PATH=/usr/local/ffmpeg-4.4.2/lib/pkgconfig
ENV LD_LIBRARY_PATH=/usr/local/ffmpeg-4.4.2/lib
ENV PATH=/usr/local/ffmpeg-4.4.2/bin:${PATH}
# GITUSER/GITPASS 不再写死 ENV，由 git_env.sh 通过 BuildKit secret 注入

RUN rm -f /etc/apt/sources.list && \
    # 确保 ubuntu.sources 已配置为阿里云源
    sed -i 's|http://archive.ubuntu.com/ubuntu/|http://mirrors.aliyun.com/ubuntu/|g' /etc/apt/sources.list.d/ubuntu.sources && \
    sed -i 's|http://security.ubuntu.com/ubuntu/|http://mirrors.aliyun.com/ubuntu/|g' /etc/apt/sources.list.d/ubuntu.sources

# Install updates and dependencies（不装系统 libav*-dev，避免与 FFmpeg 4.4 头文件冲突）
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
  && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY internal/install_ffmpeg442.sh /tmp/internal/
RUN bash /tmp/internal/install_ffmpeg442.sh

COPY internal/install_libpcre3_jammy.sh /tmp/internal/
RUN bash /tmp/internal/install_libpcre3_jammy.sh

COPY internal/install_golang.sh /tmp/internal/
RUN bash /tmp/internal/install_golang.sh

# rtpengine 需要克隆私有 GitLab，用 secret 注入凭证
COPY internal/git_env.sh internal/build_rtpengine.sh /tmp/internal/
RUN --mount=type=secret,id=gitpass \
    bash /tmp/internal/build_rtpengine.sh && \
    mkdir -p /tmp/debs && cp /tmp/*.deb /tmp/debs/ 2>/dev/null || true

# ETSI EVS codec
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
ENV FFMPEG_PREFIX=/usr/local/ffmpeg-4.4.2
ENV PKG_CONFIG_PATH=/usr/local/ffmpeg-4.4.2/lib/pkgconfig
ENV LD_LIBRARY_PATH=/usr/local/ffmpeg-4.4.2/lib
ENV PATH=/usr/local/ffmpeg-4.4.2/bin:${PATH}

# 运行时依赖（与 builder 同一组包，不含构建工具可自行裁减）
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
  && apt-get clean && rm -rf /var/lib/apt/lists/*

# 从 builder 复制自编译产物
COPY --from=builder /opt/go /opt/go
COPY --from=builder /usr/local /usr/local

# 安装 rtpengine debs（builder 已编译好）
COPY --from=builder /tmp/debs /tmp/debs
RUN if ls /tmp/debs/*.deb >/dev/null 2>&1; then \
      dpkg -i /tmp/debs/*.deb 2>/dev/null || true; \
      apt-get install -f -y; \
      rm -rf /tmp/debs; \
    fi

# 安装 libpcre3（jammy 兼容包，下载即可无需编译）
COPY internal/install_libpcre3_jammy.sh /tmp/internal/
RUN bash /tmp/internal/install_libpcre3_jammy.sh && rm -rf /tmp/internal

RUN ldconfig 2>/dev/null || true

USER root
WORKDIR /home/root/
