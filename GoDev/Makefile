
GO = CGO_ENABLED=0 go
GOCGO = CGO_ENABLED=1 go

MSERVICES = amf n2proc udm smf

OUT_BIN =
ifneq ($(GOPATH),)
        OUT_BIN = $(GOPATH)/bin
endif

SRC_DIR = $(shell pwd)
BUILD_DIR = $(SRC_DIR)/build

.PHONY: build clean install $(MSERVICES) $(DOCKERS)

VERSION=$(shell cat .version)
GIT_SHA=$(shell git rev-parse HEAD)

GOFLAGS=

build: $(MSERVICES)

oss:
	rm -rf $(OUT_BIN)/build $(OUT_BIN)/lib
	mkdir -p $(OUT_BIN)/build
	cd $(OUT_BIN)/build && cmake $(SRC_DIR)/cmn/message/ngap/ngaposs && make && mv lib $(OUT_BIN)
	rm -rf $(OUT_BIN)/build
	cp $(SRC_DIR)/cmn/message/ngap/ngaposs/lib/* $(OUT_BIN)/lib

amf:
	$(GOCGO) build $(GOFLAGS) -o $(OUT_BIN)/$@ $(SRC_DIR)/amf/

smf:
	$(GOCGO) build $(GOFLAGS) -o $(OUT_BIN)/$@ $(SRC_DIR)/smf/

n2proc:
	$(GOCGO) build $(GOFLAGS) -o $(OUT_BIN)/$@ $(SRC_DIR)/amf/n2proc

udm:
	$(GOCGO) build $(GOFLAGS) -o $(OUT_BIN)/$@ $(SRC_DIR)/udm/

upf:
	cd $(SRC_DIR)/3rdParty/github.com/intel-go/nff-go && make -j4
	cd $(SRC_DIR)/upf && make && cp upf $(OUT_BIN)/$@ -f

clean:
	@echo $(OUT_BIN)
	rm -rf $(OUT_BIN)/build
	rm -rf $(OUT_BIN)/lib
	rm -rf $(OUT_BIN)/amf
	rm -rf $(OUT_BIN)/smf
	rm -rf $(OUT_BIN)/upf
	rm -rf $(OUT_BIN)/udm
	rm -rf $(OUT_BIN)/n2proc

install:
	rm -rf $(BUILD_DIR)/install
	mkdir -p $(BUILD_DIR)/install/5gc

	# sys
	cp $(SRC_DIR)/oam/cm/data/cm_sys*.yaml $(BUILD_DIR)/install/5gc/
	 
	# amf
	mkdir -p $(BUILD_DIR)/install/5gc/amf/config $(BUILD_DIR)/install/5gc/amf/logs
	cp $(OUT_BIN)/amf $(BUILD_DIR)/install/5gc/amf/	
	cp $(OUT_BIN)/n2proc $(BUILD_DIR)/install/5gc/amf/	
	cp $(SRC_DIR)/oam/cm/data/cm_amf*.yaml $(BUILD_DIR)/install/5gc/amf/config
	cd $(BUILD_DIR)/install/5gc/amf/config && ln -snf ../../cm_sys_conf.yaml cm_sys_conf.yaml

	# smf
	mkdir -p $(BUILD_DIR)/install/5gc/smf/config $(BUILD_DIR)/install/5gc/smf/logs
	cp $(OUT_BIN)/smf $(BUILD_DIR)/install/5gc/smf/
	cp $(SRC_DIR)/oam/cm/data/cm_smf*.yaml $(BUILD_DIR)/install/5gc/smf/config

	# udm
	mkdir -p $(BUILD_DIR)/install/5gc/udm/config $(BUILD_DIR)/install/5gc/udm/logs
	cp $(OUT_BIN)/udm $(BUILD_DIR)/install/5gc/udm/
	cp $(SRC_DIR)/oam/cm/data/cm_udm*.yaml $(BUILD_DIR)/install/5gc/udm/config
	cd $(BUILD_DIR)/install/5gc/udm/config && ln -snf ../../cm_sys_conf.yaml cm_sys_conf.yaml

    # upf
	# mkdir -p $(BUILD_DIR)/install/5gc/upf/config $(BUILD_DIR)/install/5gc/upf/logs
	# mkdir -p $(BUILD_DIR)/install/5gc/upf/dpdk/build/kmod $(BUILD_DIR)/install/5gc/upf/dpdk/usertools
	# cp $(OUT_BIN)/upf $(BUILD_DIR)/install/5gc/upf/
	# cp $(SRC_DIR)/oam/cm/data/cm_upf*.yaml $(BUILD_DIR)/install/5gc/upf/config
	# cp $(SRC_DIR)/3rdParty/github.com/intel-go/nff-go/dpdk/dpdk/usertools/* $(BUILD_DIR)/install/5gc/upf/dpdk/usertools
	# cp $(SRC_DIR)/3rdParty/github.com/intel-go/nff-go/dpdk/dpdk/build/kmod/* $(BUILD_DIR)/install/5gc/upf/dpdk/build/kmod

	# script
	cp $(BUILD_DIR)/script/core_mgr.sh $(BUILD_DIR)/install/5gc/
	chmod +x $(BUILD_DIR)/install/5gc/*.sh
	cp $(BUILD_DIR)/docker/docker-compose.yml $(BUILD_DIR)/install/5gc/
	cp $(BUILD_DIR)/docker/.env $(BUILD_DIR)/install/5gc/
	mkdir -p $(BUILD_DIR)/install/5gc/mysql/conf
	cp $(BUILD_DIR)/script/conf/my.cnf  $(BUILD_DIR)/install/5gc/mysql/conf 
	mkdir -p $(BUILD_DIR)/install/5gc/mysql/db
	mkdir -p $(BUILD_DIR)/install/5gc/mysql/script
	cp $(SRC_DIR)/udm/script/*.sql $(BUILD_DIR)/install/5gc/mysql/script
	mkdir -p $(BUILD_DIR)/install/5gc/redis/data
	cp $(BUILD_DIR)/script/conf/redis.conf  $(BUILD_DIR)/install/5gc/redis/

	# create a tar ball
	cd $(BUILD_DIR)/install && tar cf lite5gc_$(VERSION).tgz 5gc

