FROM registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000/ubuntu:26.04

ENV DEBIAN_FRONTEND=noninteractive
ENV http_proxy="http://10.18.11.52:7890" https_proxy="http://10.18.11.52:7890" \
    no_proxy="mirrors.aliyun.com,registry.cn-hangzhou.aliyuncs.com,npmmirror.com,10.18.1.2"

COPY internal /tmp/internal
RUN bash /tmp/internal/preinstall.sh
RUN bash /tmp/internal/install_node.sh
RUN bash /tmp/internal/postinstall.sh

ENV http_proxy= https_proxy=
ENV NODE_OPTIONS=--max-old-space-size=8192
ENV PNPM_HOME=/root/.local/share/pnpm
ENV PATH=${PNPM_HOME}:${PATH}
WORKDIR /root/app
CMD ["/bin/bash"]
