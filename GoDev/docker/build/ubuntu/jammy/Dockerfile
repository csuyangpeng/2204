FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/ubuntu:22.04

ARG MAKEFLAGS=-j TARGETARCH
ENV GOROOT=/opt/go
ENV GOPATH=/home/sder/go
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GITUSER=peng.yang GITPASS=11asdASD
# ENV http_proxy="http://10.18.11.52:7890" https_proxy="http://10.18.11.52:7890"
ENV PATH=${GOROOT}/bin:${GOPATH}/bin:${PATH}

COPY internal /tmp/internal
RUN bash /tmp/internal/preinstall.sh
RUN bash /tmp/internal/install_sder.sh
RUN bash /tmp/internal/install_golang.sh
RUN bash /tmp/internal/install_3rd.sh
RUN bash /tmp/internal/install_tcpdump.sh
# RUN bash /tmp/internal/install_trace.sh
# RUN bash /tmp/internal/install_backend.sh
# RUN bash /tmp/internal/install_core.sh
RUN bash /tmp/internal/postinstall.sh

ENV GITUSER= GITPASS= http_proxy= https_proxy=
USER sder
WORKDIR ${GOPATH}/src/lite5gc
