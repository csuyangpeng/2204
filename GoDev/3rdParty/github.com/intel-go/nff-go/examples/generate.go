// Copyright 2017 Intel Corporation.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"
	"github.com/intel-go/nff-go/types"
	"time"
)

func main() {
	mode := flag.Int("mode", 2, "mode of generating:\n0 - fast generate that will be slowed in a second.\n1 - time-based generate send by 32 packets.\n2 - time-based generate send by 1 packet.")
	output := flag.Int("port", 1, "output port")
	flag.Parse()
	outputPort := uint16(*output)

	config := flow.Config{
		CPUList: "0-3",
		//NoPacketHeadChange: true,
		HWTXChecksum: true,
	}

	flow.SystemInit(&config)

	switch *mode {
	case 0:
		//firstFlow, genChannel, _ := flow.SetFastGenerator(generatePacket, 100, nil)
		firstFlow, _ := flow.SetReceiverOS("enp2s0f1")
		//flow.CheckFatal(flow.SetSender(firstFlow, outputPort))
		flow.CheckFatal(flow.SetSenderOS(firstFlow, "enp2s0f1"))
		//go updateSpeed(genChannel)
		flow.SystemStart()
	case 1:
		firstFlow := flow.SetGenerator(generatePacket1, nil)
		flow.CheckFatal(flow.SetSender(firstFlow, outputPort))
		flow.SystemStart()
	case 2:
		temp, _ := (flow.SetReceiver(outputPort))
		flow.SetStopper(temp)
		flow.SystemInitPortsAndMemory()
		generatePacket2(outputPort)
	}
}

func generatePacket(pkt *packet.Packet, context flow.UserContext) {
	//packet.InitEmptyIPv4Packet(pkt, 1300)
	packet.InitEmptyIPv4UDPPacket(pkt, 1300)

	//        52:54:00:07:a8:cc
	//pkt.Ether.DAddr = [6]uint8{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
	//todo zj test
	pkt.Ether.DAddr = [6]uint8{0x52, 0x54, 0x00, 0x07, 0xa8, 0xcc}
	// a0:36:9f:84:5c:bd
	pkt.Ether.SAddr = [6]uint8{0xa0, 0x36, 0x9f, 0x84, 0x5c, 0xbd}

	pkt.ParseL3()
	srcip := []byte{172, 20, 0, 18}
	dstip := []byte{172, 20, 0, 16}
	pkt.GetIPv4().SrcAddr = types.SliceToIPv4(srcip)
	pkt.GetIPv4().DstAddr = types.SliceToIPv4(dstip)

	pkt.ParseL4ForIPv4()
	pkt.GetUDPForIPv4().DstPort = packet.SwapUDPPortGTPU
}

func generatePacket1(pkt *packet.Packet, context flow.UserContext) {
	packet.InitEmptyIPv4Packet(pkt, 1300)
	pkt.Ether.DAddr = [6]uint8{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
	time.Sleep(175 * time.Microsecond)
}

func generatePacket2(port uint16) {
	for {
		pkt, _ := packet.NewPacket()
		packet.InitEmptyIPv4Packet(pkt, 1300)
		pkt.Ether.DAddr = [6]uint8{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
		pkt.SendPacket(port)
		time.Sleep(175 * time.Microsecond)
	}
}

func updateSpeed(genChannel chan uint64) {
	var load int
	for {
		// Can be file or any other source
		if _, err := fmt.Scanf("%d", &load); err == nil {
			genChannel <- uint64(load)
		}
	}
}
