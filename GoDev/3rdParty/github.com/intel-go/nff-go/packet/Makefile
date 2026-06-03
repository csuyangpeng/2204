# Copyright 2017 Intel Corporation.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.
export NFF_GO_NO_MLX_DRIVERS=1
export NFF_GO_NO_BPF_SUPPORT=1
PATH_TO_MK = ../mk
include $(PATH_TO_MK)/include.mk


.PHONY: testing
testing: check-pktgen
	go test -tags "${GO_BUILD_TAGS}"

.PHONY: coverage
coverage:
	go test -cover -coverprofile=c.out
	go tool cover -html=c.out -o packet_coverage.html
