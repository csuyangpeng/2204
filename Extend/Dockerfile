# 基于原镜像
FROM 10.18.1.2:5000/lite5gc:dev-x86_64-22.04-20240115

# 设置环境变量
# ENV GOPROXY="http://10.18.1.2:8081,direct" \
#     HTTP_PROXY="http://10.18.11.52:7890" \
#     HTTPS_PROXY="http://10.18.11.52:7890" \
#     NO_PROXY="localhost,127.0.0.1,.internal"
ENV GOSUMDB="sum.golang.google.cn" \
    GOPROXY=http://10.18.1.2:8081,direct

