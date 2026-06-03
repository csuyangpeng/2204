#!/usr/bin/env bash


set -e

cd "$(dirname "${BASH_SOURCE[0]}")"

git clone http://${GITUSER}:${GITPASS}@10.18.1.2:9999/develop/lite5gc.git 
pushd lite5gc
git checkout master
go mod init lite5gc
# echo "replace github.com/intel-go/nff-go v0.9.2 => ./3rdParty/github.com/intel-go/nff-go" >> go.mod
# echo "replace github.com/upf v0.0.1 => ./3rdParty/github.com/xdp-upf" >> go.mod
# echo "replace github.com/eupf v0.0.1 => ./3rdParty/github.com/eupf" >> go.mod
# echo "replace github.com/dpi v0.0.1 => ./3rdParty/github.com/dpi" >> go.mod
# echo "require github.com/spf13/cobra v1.7.0" >> go.mod
# echo "require github.com/Shopify/sarama v1.36.0" >> go.mod
# echo "require github.com/gin-gonic/gin v1.9.0" >> go.mod
# echo "require github.com/smallnest/rpcx v1.7.11" >> go.mod
# echo "require github.com/mdlayher/ndp v0.8.0" >> go.mod
# echo "require github.com/xuri/excelize/v2 v2.6.0" >> go.mod
# echo "require gorm.io/plugin/opentelemetry v0.1.1" >> go.mod
# echo "replace github.com/fiorix/go-diameter/v4 v4.0.4 => ./3rdParty/github.com/fiorix/go-diameter/v4" >> go.mod
echo "replace github.com/upf v0.0.1 => ./3rdParty/github.com/xdp-upf" >> go.mod
echo "replace github.com/dpi v0.0.1 => ./3rdParty/github.com/dpi" >> go.mod
echo "replace github.com/eupf v0.0.1 => ./3rdParty/github.com/eupf" >> go.mod
echo "require github.com/spf13/cobra v1.7.0" >> go.mod
echo "require github.com/gin-gonic/gin v1.9.0" >> go.mod
echo "require github.com/smallnest/rpcx v1.7.11" >> go.mod
echo "require github.com/mdlayher/ndp v0.8.0" >> go.mod
echo "require github.com/xuri/excelize/v2 v2.6.0" >> go.mod
echo "require gorm.io/gorm v1.25.11" >> go.mod
echo "replace github.com/fiorix/go-diameter/v4 v4.0.4 => ./3rdParty/github.com/fiorix/go-diameter/v4" >> go.mod
go mod tidy && bash lite5gc.sh build
mv build/output/oam ${GOPATH}/bin/
mv build/output/cli ${GOPATH}/bin/
popd

rm lite5gc /root/.cache ${GOPATH}/pkg -rf
