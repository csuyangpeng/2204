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

ENV DEBIAN_FRONTEND=noninteractive

ENV PIP_DEFAULT_TIMEOUT=100
ENV PIP_DISABLE_PIP_VERSION_CHECK=1
ENV PIP_BREAK_SYSTEM_PACKAGES=1
ENV PYTHONWARNINGS="ignore:Unverified HTTPS request"
ENV http_proxy=http://10.18.11.52:7890 https_proxy=http://10.18.11.52:7890 \
    no_proxy=mirrors.aliyun.com,registry.cn-hangzhou.aliyuncs.com,10.18.1.2,pypi.tuna.tsinghua.edu.cn

COPY internal /tmp/internal

# Install updates and dependencies
RUN bash /tmp/internal/prepare_apt.sh && \
	apt-get -y install mysql-server tcpdump screen tmux ntpsec ntpsec-ntpdate git-core dkms \
					flex bison libmysqlclient-dev make libssl-dev libcurl4-openssl-dev \
					libxml2-dev libpcre2-dev libpcre3-dev bash-completion g++-11 autoconf libmnl-dev \
					libsctp-dev libradcli-dev libradcli4 libjson-c-dev pkg-config iproute2 net-tools \
					iputils-ping libgcrypt20-dev libjansson-dev libevent-dev libnghttp2-dev \
					systemctl rsyslog gdb logrotate lua5.1 lua5.1-0-dev


# Install luarocks, lua-cjson, Python and Supervisor
# Optimize space usage by cleaning up after each step
RUN bash /tmp/internal/prepare_apt.sh && \
    apt-get install -y --no-install-recommends build-essential lua5.1 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN bash /tmp/internal/prepare_apt.sh && \
    apt-get install -y --no-install-recommends luarocks && \
    luarocks install lua-cjson && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN bash /tmp/internal/prepare_apt.sh && \
    apt-get install -y --no-install-recommends libmicrohttpd-dev && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install Python and required Python packages
# Use system packages when possible to avoid network issues
RUN bash /tmp/internal/prepare_apt.sh && \
    apt-get install -y --no-install-recommends \
        python3 \
        python3-pip \
        python3-requests \
        python3-yaml \
        python3-setuptools \
        python3-wheel && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# mysql-connector-python（configparser 为 Python3 标准库，无需 pip 安装）
RUN pip3 install --no-cache-dir --index-url https://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com --timeout 60 mysql-connector-python

# Install supervisor
RUN bash /tmp/internal/prepare_apt.sh && \
    apt-get install -y --no-install-recommends supervisor && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Install etcd3-py
RUN pip3 install --no-cache-dir --index-url https://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com --timeout 180 etcd3 && \
    pip3 install --no-cache-dir protobuf==3.19.6

RUN pip3 install redis

# Create necessary directories
RUN mkdir -p /var/log /var/run/supervisor /coredumps
# 复制并设置初始化脚本

RUN mkdir -p /root/.pip && \
    echo "[global]" > /root/.pip/pip.conf && \
    echo "index-url = https://pypi.tuna.tsinghua.edu.cn/simple" >> /root/.pip/pip.conf && \
    echo "trusted-host = pypi.tuna.tsinghua.edu.cn" >> /root/.pip/pip.conf && \
    echo "verify = false" >> /root/.pip/pip.conf

# 安装系统依赖
RUN bash /tmp/internal/prepare_apt.sh && \
    apt-get -y install \
        bind9 \
        ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /tmp/internal

ENV http_proxy= https_proxy=
# 创建必要的目录并设置权限
RUN mkdir -p /var/log/supervisor /var/run/named /var/cache/bind /etc/bind && \
    chown -R bind:bind /var/cache/bind /etc/bind /var/run/named

# 复制配置文件
#COPY requirements.txt /tmp/

# 安装Python依赖
#RUN pip3 install --no-cache-dir --trusted-host pypi.tuna.tsinghua.edu.cn -r /tmp/requirements.txt && \
#   rm -f /tmp/requirements.txt
