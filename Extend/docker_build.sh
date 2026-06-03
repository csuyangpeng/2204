# Script used to start lite5gc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
ALIREPOSITORY=registry.cn-hangzhou.aliyuncs.com/10_18_1_2_5000
TAGS=dev-x86_64-22.04-20240115
docker build --progress=plain -t lite5gc:${TAGS} . -f Dockerfile
docker tag lite5gc:${TAGS} ${REPOSITORY}/lite5gc:${TAGS}
docker push ${REPOSITORY}/lite5gc:${TAGS}
# docker tag lite5gc:${TAGS} ${ALIREPOSITORY}/lite5gc:${TAGS}
# docker push ${ALIREPOSITORY}/lite5gc:${TAGS}


docker run --detach --hostname 10.18.1.30 --publish 443:443 --publish 9999:9999 --publish 8888:22 --name gitlab --restart always -v /data/git/tmpgitlab/config:/etc/gitlab -v /data/git/tmpgitlab/logs:/var/log/gitlab -v /data/git/tmpgitlab/data:/var/opt/gitlab docker.m.daocloud.io/gitlab/gitlab-ee:16.11.0-ee.0
