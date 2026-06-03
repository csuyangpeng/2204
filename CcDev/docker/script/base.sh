#!/usr/bin/env bash

BOLD='\033[1m'
RED='\033[0;31m'
GREEN='\033[32m'
WHITE='\033[34m'
YELLOW='\033[33m'
NO_COLOR='\033[0m'

function info() {
    (>&2 echo -e "[${WHITE}${BOLD}INFO${NO_COLOR}] $*")
}

function error() {
    (>&2 echo -e "[${RED}ERROR${NO_COLOR}] $*")
}

function warning() {
    (>&2 echo -e "${YELLOW}[WARNING] $*${NO_COLOR}")
}

function ok() {
    (>&2 echo -e "[${GREEN}${BOLD} OK ${NO_COLOR}] $*")
}

SUPPORTED_OS=("Linux")
function check_host_environment() {
    local os="$(uname -s)"
    for ent in "${SUPPORTED_OS[@]}"; do
        if [[ `uname -s` == "Linux" ]]; then
            ok "target os ${os} is supported"
            return 0
        fi
    done
    warning "Running lite5gc dev container on os ${HOST_OS} is UNTESTED, exiting..."
    exit 1
}

SUPPORTED_ARCHS=(x86_64)
function check_target_arch() {
    local arch="$(uname -m)"
    for ent in "${SUPPORTED_ARCHS[@]}"; do
        if [[ "${ent}" == "${arch}" ]]; then
            ok "target arch ${arch} is supported"
            return 0
        fi
    done
    error "Running lite5gc dev container on arch ${arch} is UNTESTED, exiting..."
    exit 1
}
