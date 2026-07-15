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

FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/kamailio:dev-26.04-20260624

ARG TARGETARCH
ENV TARGETARCH=${TARGETARCH}
ENV DEBIAN_FRONTEND=noninteractive

ENV PIP_DEFAULT_TIMEOUT=100
ENV PIP_DISABLE_PIP_VERSION_CHECK=1
ENV PIP_BREAK_SYSTEM_PACKAGES=1
ENV PYTHONWARNINGS="ignore:Unverified HTTPS request"
ENV http_proxy=http://10.18.11.52:7890 https_proxy=http://10.18.11.52:7890 \
    no_proxy=mirrors.aliyun.com,registry.cn-hangzhou.aliyuncs.com,10.18.1.2,pypi.tuna.tsinghua.edu.cn

COPY internal /tmp/internal

# ---- 系统包（合并为一个 RUN + apt cache mount，避免 arm64 QEMU 下重复 update） ----
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    bash /tmp/internal/prepare_apt.sh && \
    apt-get -y install mysql-server tcpdump screen tmux ntpsec ntpsec-ntpdate git-core dkms \
        flex bison libmysqlclient-dev make libssl-dev libcurl4-openssl-dev \
        libxml2-dev libpcre2-dev bash-completion g++ autoconf libmnl-dev \
        libsctp-dev libradcli-dev libradcli4 libjson-c-dev pkg-config iproute2 net-tools \
        iputils-ping libgcrypt20-dev libjansson-dev libevent-dev libnghttp2-dev \
        systemctl rsyslog gdb logrotate lua5.1 lua5.1-0-dev \
        build-essential luarocks libmicrohttpd-dev \
        python3 python3-pip python3-requests python3-yaml python3-setuptools python3-wheel \
        supervisor && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# ---- Python / Lua 包安装 ----
# mysql-connector-python（configparser 为 Python3 标准库，无需 pip 安装）
RUN pip3 install --no-cache-dir --index-url https://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com --timeout 60 mysql-connector-python

# etcd3 + protobuf（etcd3 需要 protobuf<4）
RUN pip3 install --no-cache-dir --index-url https://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com --timeout 180 etcd3 && \
    pip3 install --no-cache-dir protobuf==3.19.6

RUN pip3 install redis

# Lua cjson
RUN luarocks install lua-cjson

# ---- 目录与配置 ----
RUN mkdir -p /var/log /var/run/supervisor /coredumps

RUN mkdir -p /root/.pip && \
    echo "[global]" > /root/.pip/pip.conf && \
    echo "index-url = https://pypi.tuna.tsinghua.edu.cn/simple" >> /root/.pip/pip.conf && \
    echo "trusted-host = pypi.tuna.tsinghua.edu.cn" >> /root/.pip/pip.conf && \
    echo "verify = false" >> /root/.pip/pip.conf

# bind9 + 清理
RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    bash /tmp/internal/prepare_apt.sh && \
    apt-get -y install bind9 ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/internal

ENV http_proxy= https_proxy=

# 创建必要的目录并设置权限
RUN mkdir -p /var/log/supervisor /var/run/named /var/cache/bind /etc/bind && \
    chown -R bind:bind /var/cache/bind /etc/bind /var/run/named
