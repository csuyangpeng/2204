# Script used to into lite5gc-dev containter
#!/usr/bin/env bash

set -e

LITE5GC_ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"
source ${LITE5GC_ROOT_DIR}/docker/script/base.sh

function buildamf() {
	info "compiling oss .........................."
	make oss -j

    sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.8 100
	sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-4.8 100
	info "compiling amf .........................."
	make amf -j
	sudo update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-4.8 10
	sudo update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-4.8 10
}

function buildupf() {
	info "compiling upf .........................."
	make upf -j
}

AVAILABLE_COMMANDS=" build test coverage lint doc clean format usage "

function usage() {
    echo -e "\n${GREEN}${BOLD}Usage${NO_COLOR}:
    .${BOLD}/lite5gc.sh${NO_COLOR} [OPTION]"
    echo -e "\n${GREEN}${BOLD}Options${NO_COLOR}:
    ${BLUE}build [module]${NO_COLOR}: run build for lite5gc (<module> = ) or modules/<module>.  If <module> unspecified, build all.
    ${BLUE}test [module]${NO_COLOR}: run unittest for lite5gc (module='') or modules/<module>. If unspecified, test all.
    ${BLUE}coverage [module]${NO_COLOR}: run coverage test for lite5gc (module='') or modules/<module>. If unspecified, coverage all.
    ${BLUE}lint${NO_COLOR}: run code style check
    ${BLUE}doc${NO_COLOR}: generate doxygen document
    ${BLUE}clean${NO_COLOR}: cleanup bazel output and log/coredump files
    ${BLUE}format${NO_COLOR}: format C++/Python/Bazel/Shell files
    ${BLUE}usage${NO_COLOR}: show this message and exit
    "
}

function main() {
	echo $#
    if [ $# -eq 0 ]; then
        usage
        exit 0
    fi

    check_host_environment
    check_target_arch

    local cmd="$1"
    shift
    case "${cmd}" in
        build)
            buildamf
            buildupf
            ;;
        test)
            ;;
        coverage)
            ;;
        lint)
            ;;
        clean)
            ;;
        doc)
            ;;
        format)
            ;;
        usage)
            ;;
        -h|--help)
            ;;
        *)
        	usage
            ;;
    esac
}

main "$@"
