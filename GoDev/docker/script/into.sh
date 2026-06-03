# Script used to into lite5gc-dev containter
#!/usr/bin/env bash

DOCKER_USER="${USER}"
LITE5GC_DEV="lite5gc-dev-${USER}"
#set -x
xhost +local:root 1>/dev/null 2>&1

docker exec \
    -u "${DOCKER_USER}" \
    -it "${LITE5GC_DEV}" \
    /bin/bash

xhost -local:root 1>/dev/null 2>&1
#set +x
