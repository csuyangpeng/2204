FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/ubuntu:26.04

ARG MAKEFLAGS=-j
# GITUSER/GITPASS/GIT_PROXY 由 git_env.sh 通过 BuildKit secret 注入，不写入镜像层
ENV LD_LIBRARY_PATH=/usr/local/lib
ENV http_proxy="http://10.18.11.52:7890" https_proxy="http://10.18.11.52:7890" \
    no_proxy="mirrors.aliyun.com,10.18.1.2,registry.cn-hangzhou.aliyuncs.com,goproxy.cn"

COPY internal/preinstall.sh /tmp/internal/
COPY internal/sources.list /etc/apt/sources.list
RUN bash /tmp/internal/preinstall.sh

# 各 install 脚本单独 COPY，改一个只重编后续层（不用每次从 preinstall 重来）
# RUN bash /tmp/internal/install_msgpack.sh
COPY internal/git_env.sh internal/install_wfrest.sh /tmp/internal/
RUN --mount=type=secret,id=gitpass bash /tmp/internal/install_wfrest.sh

COPY internal/git_env.sh internal/install_cpp-httplib.sh /tmp/internal/
RUN --mount=type=secret,id=gitpass bash /tmp/internal/install_cpp-httplib.sh

COPY internal/git_env.sh internal/install_cpprestsdk.sh /tmp/internal/
COPY internal/patches/cpprestsdk-fix-asio-error.patch /tmp/internal/patches/
RUN --mount=type=secret,id=gitpass bash /tmp/internal/install_cpprestsdk.sh

COPY internal/git_env.sh internal/install_etcd-cpp-apiv3.sh /tmp/internal/
RUN --mount=type=secret,id=gitpass bash /tmp/internal/install_etcd-cpp-apiv3.sh

# RUN bash /tmp/internal/install_cppkafka.sh
COPY internal/git_env.sh internal/install_nlohmannjson.sh /tmp/internal/
RUN --mount=type=secret,id=gitpass bash /tmp/internal/install_nlohmannjson.sh

# RUN bash /tmp/internal/install_zlog.sh
# RUN bash /tmp/internal/install_core.sh
COPY internal/postinstall.sh internal/setup.sh /tmp/internal/
RUN bash /tmp/internal/postinstall.sh

USER root
WORKDIR /home/root/
