# Script used to start epc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

REPOSITORY=10.18.1.2:5000
# TAGS=dev-x86_64-22.04-`date +%Y%m%d_%H%M`-${VER}
TAGS=base
docker build -t epc:${TAGS} . -f Dockerfile
docker tag epc:${TAGS} ${REPOSITORY}/epc:${TAGS}
docker push ${REPOSITORY}/epc:${TAGS}
