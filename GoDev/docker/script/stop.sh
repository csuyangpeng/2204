# Script used to stop lite5gc-dev containter
#!/usr/bin/env bash

#set -x
LITE5GC_ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"
source ${LITE5GC_ROOT_DIR}/docker/script/base.sh
LITE5GC_DEV="lite5gc-dev-${USER}"
function remove_existing_dev_container() {
    if docker ps -a --format '{{.Names}}' | grep -q "${LITE5GC_DEV}"; then
        docker stop "${LITE5GC_DEV}" >/dev/null
        docker rm -v -f "${LITE5GC_DEV}" >/dev/null
    fi
}

info "Check and remove existing lite5gc dev container ..."
remove_existing_dev_container

#set +x
