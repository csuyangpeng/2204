# Script used to start epc-dev containter
#!/usr/bin/env bash
set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

docker build --progress=plain -t epc:base18 . -f Dockerfile
