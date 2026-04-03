#!/usr/bin/env bash

set -e

cd "$(dirname "${BASH_SOURCE[0]}")"
source "$(dirname "${BASH_SOURCE[0]}")/git_env.sh"

git clone --recursive https://github.com/wfrest/wfrest
pushd wfrest/workflow && make -j8 && make install && popd 
pushd wfrest && make -j8 && make install && popd
rm -rf wfrest
