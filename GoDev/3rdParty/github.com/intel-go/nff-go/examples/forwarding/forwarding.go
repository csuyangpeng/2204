// Copyright 2017-2019 Intel Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/intel-go/nff-go/types"
	"net"

	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"
)

var l3Rules *packet.L3Rules

// Main function for constructing packet processing graph.
func main() {
	inport := flag.Uint("inport", 0, "Port for receiving packets.")
	//numflows := flag.Uint("numflows", 5, "Number of output flows to use. First flow with number zero is used for dropped packets.")
	numflows := flag.Uint("numflows", 2, "Number of output flows to use. First flow with number zero is used for dropped packets.")
	nostats := flag.Bool("nostats", false, "Disable statics HTTP server.")
	noscheduler := flag.Bool("no-scheduler", true, "Disable scheduler.")
	dpdkLogLevel := flag.String("dpdk", "--log-level=8", "Passes an arbitrary argument to dpdk EAL.")

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
		StatsHTTPAddress:  statsServerAddres,
		SchedulerInterval: 100,
		DPDKArgs:          []string{*dpdkLogLevel},
		DisableScheduler:  *noscheduler,
	}
	flow.CheckFatal(flow.SystemInit(&config))

	// Get splitting rules from access control file.
	l3Rules, err = packet.GetL3ACLFromTextTable("forwarding.config")
	flow.CheckFatal(err)

	// Receive packets from zero port. Receive queue will be added automatically.
	inputFlow, err := flow.SetReceiver(uint16(*inport))
	//inputFlow, err := flow.SetReceiverOS("enp2s0f1")
	flow.CheckFatal(err)

	//arp
	//大端存
	ipv4 := types.IPv4Address(20)<<24 | types.IPv4Address(20)<<16 | types.IPv4Address(20)<<8 | types.IPv4Address(172)

	flow.CheckFatal(flow.SetIPForPort(uint16(*inport), ipv4))
	fmt.Println(types.IPv4ToBytes(ipv4))
	fmt.Println(flow.GetNameByPort(uint16(*inport)))
	fmt.Printf("%#x", flow.GetPortMACAddress(uint16(*inport)))

	flow.CheckFatal(flow.DealARPICMP(inputFlow))

	// Split packet flow based on ACL.
	flowsNumber := uint16(*numflows)
	//outputFlows, err := flow.SetSplitter(inputFlow, l3Splitter, uint(flowsNumber), nil)
	secondFlow, err := flow.SetSeparator(inputFlow, NffN3Handler, nil)
	flow.CheckFatal(err)

	// "0" flow is used for dropping packets without sending them.
	flow.CheckFatal(flow.SetStopper(secondFlow))

	// Send each flow to corresponding port. Send queues will be added automatically.
	for i := uint16(1); i < flowsNumber; i++ {
		//todo zj
		flow.CheckFatal(flow.SetHandler(inputFlow, HandePacket, nil))
		//flow.CheckFatal(flow.SetSender(outputFlows[i], i-1))
		flow.CheckFatal(flow.SetSender(inputFlow, 1))
		//flow.CheckFatal(flow.SetSenderOS(inputFlow, "enp2s0f2"))
	}

	// Begin to process packets.
	flow.CheckFatal(flow.SystemStart())
}

// User defined function for splitting packets
func l3Splitter(currentPacket *packet.Packet, context flow.UserContext) uint {
	// Return number of flow to which put this packet. Based on ACL rules.
	return currentPacket.L3ACLPort(l3Rules)
}
func NffN3Handler(currentPacket *packet.Packet, context flow.UserContext) bool {
	ipv4, _, _ := currentPacket.ParseAllKnownL3()
	if ipv4 == nil {
		return false
	}
	return true // 保留在原流中
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
		if pkt.GetIPv4().NextProtoID == types.UDPNumber {
			pkt.ParseL4ForIPv4()
			pkt.GetUDPForIPv4().DstPort = packet.SwapUDPPortGTPU
		}
	}
}
