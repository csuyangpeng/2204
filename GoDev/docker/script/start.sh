# Script used to start lite5gc-dev containter
#!/usr/bin/env bash

#set -x
LITE5GC_ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"
source ${LITE5GC_ROOT_DIR}/docker/script/base.sh
DEV_INSIDE="in-dev-docker"
LITE5GC_DEV="lite5gc-dev-${USER}"
#LITE5GC_REPO="lite5gc"
#UBUNTU_LTS="18.04"
#TIMESTAMP=20201223_1358
LITE5GC_IMAGE="lite5gc:dev-x86_64-18.04-20201223_1502"
#LITE5GC_IMAGE="${LITE5GC_REPO}:dev-x86_64-${UBUNTU_LTS}-${TIMESTAMP}"
check_host_environment
check_target_arch
function remove_existing_dev_container() {
    if docker ps -a --format '{{.Names}}' | grep -q "${LITE5GC_DEV}"; then
        docker stop "${LITE5GC_DEV}" >/dev/null
        docker rm -v -f "${LITE5GC_DEV}" >/dev/null
    fi
}

info "Check and remove existing lite5gc dev container ..."
remove_existing_dev_container

bash ${LITE5GC_ROOT_DIR}/docker/script/registry.sh
#set +x

docker pull 10.18.1.2:5000/${LITE5GC_IMAGE}
docker tag 10.18.1.2:5000/${LITE5GC_IMAGE} ${LITE5GC_IMAGE}
GOPATH=/home/sder/go
docker run -itd --privileged --name "${LITE5GC_DEV}" \
    -e DOCKER_USER="${USER}" \
    --net host \
    --add-host "${DEV_INSIDE}:127.0.0.1" \
    --hostname "${DEV_INSIDE}" \
    --pid=host \
    -v `pwd`:${GOPATH}/src/lite5gc \
    -v /dev/null:/dev/raw1394 \
    "${LITE5GC_IMAGE}" \
    /bin/bash

    if [ $? -ne 0 ]; then
        error "Failed to start docker container \"${LITE5GC_DEV}\" based on image: ${LITE5GC_IMAGE}"
        exit 1
    fi

    ok "Congratulations! You have successfully finished setting up lite5gc Dev Environment."
    ok "To login into the newly created ${LITE5GC_DEV} container, please run the following command:"
    ok "  bash docker/script/into.sh"
    ok "Enjoy!"
