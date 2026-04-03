# Script used to start epc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

pushd internal/epc
# git checkout master
git checkout 834b7c56
# git pull
VER=$(git rev-parse --short HEAD)
popd
REPOSITORY=10.18.1.2:5000
# TAGS=dev-x86_64-18.04-`date +%Y%m%d_%H%M`-${VER}
TAGS=dev-2.5.7
docker build --progress=plain -t epc:${TAGS} . -f Dockerfile
docker tag epc:${TAGS} ${REPOSITORY}/epc:${TAGS}
docker push ${REPOSITORY}/epc:${TAGS}
