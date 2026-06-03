// Copyright 2017-2019 Intel Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"github.com/intel-go/nff-go/types"
	"net"

	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"
)

var l3Rules *packet.L3Rules

// Main function for constructing packet processing graph.
func main() {
	//inport := flag.Uint("inport", 0, "Port for receiving packets.")
	//numflows := flag.Uint("numflows", 5, "Number of output flows to use. First flow with number zero is used for dropped packets.")
	numflows := flag.Uint("numflows", 2, "Number of output flows to use. First flow with number zero is used for dropped packets.")
	nostats := flag.Bool("nostats", false, "Disable statics HTTP server.")
	flag.Parse()

	var err error
	var statsServerAddres *net.TCPAddr = nil
	if !*nostats {
		// Set up address for stats web server
		statsServerAddres = &net.TCPAddr{
			Port: 8080,
			IP:   net.IP{10, 180, 9, 78},
		}
	}

	// Initialize NFF-GO library at 16 cores by default
	config := flow.Config{
		StatsHTTPAddress: statsServerAddres,
	}
	flow.CheckFatal(flow.SystemInit(&config))

	// Get splitting rules from access control file.
	l3Rules, err = packet.GetL3ACLFromTextTable("forwarding.config")
	flow.CheckFatal(err)

	// Receive packets from zero port. Receive queue will be added automatically.
	//inputFlow, err := flow.SetReceiver(uint16(*inport))
	inputFlow, err := flow.SetReceiverOS("enp2s0f1")
	flow.CheckFatal(err)

	// Split packet flow based on ACL.
	flowsNumber := uint16(*numflows)
	outputFlows, err := flow.SetSplitter(inputFlow, l3Splitter, uint(flowsNumber), nil)
	flow.CheckFatal(err)

	// "0" flow is used for dropping packets without sending them.
	flow.CheckFatal(flow.SetStopper(outputFlows[0]))

	// Send each flow to corresponding port. Send queues will be added automatically.
	for i := uint16(1); i < flowsNumber; i++ {
		//todo zj
		flow.CheckFatal(flow.SetHandler(outputFlows[i], HandePacket, nil))
		//flow.CheckFatal(flow.SetSender(outputFlows[i], i-1))
		flow.CheckFatal(flow.SetSenderOS(outputFlows[i], "enp2s0f2"))
	}

	// Begin to process packets.
	flow.CheckFatal(flow.SystemStart())
}

// User defined function for splitting packets
func l3Splitter(currentPacket *packet.Packet, context flow.UserContext) uint {
	// Return number of flow to which put this packet. Based on ACL rules.
	return currentPacket.L3ACLPort(l3Rules)
}

func HandePacket(pkt *packet.Packet, context flow.UserContext) {
	//packet.InitEmptyIPv4Packet(pkt, 1300)
	//packet.InitEmptyIPv4UDPPacket(pkt, 1300)

	//        52:54:00:07:a8:cc
	//pkt.Ether.DAddr = [6]uint8{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
	//todo zj test

	pkt.Ether.DAddr = [6]uint8{0x52, 0x54, 0x00, 0x07, 0xa8, 0xcc}
	// a0:36:9f:84:5c:bd
	pkt.Ether.SAddr = [6]uint8{0xa0, 0x36, 0x9f, 0x84, 0x5c, 0xbd}

	pkt.ParseL3()
	if pkt.Ether.EtherType == types.SwapIPV4Number {

		srcip := []byte{172, 20, 0, 18}
		dstip := []byte{172, 20, 0, 16}
		pkt.GetIPv4().SrcAddr = types.SliceToIPv4(srcip)
		pkt.GetIPv4().DstAddr = types.SliceToIPv4(dstip)

		pkt.ParseL4ForIPv4()
		pkt.GetUDPForIPv4().DstPort = packet.SwapUDPPortGTPU
	}
}
