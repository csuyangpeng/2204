package up

import (
	"fmt"
	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"
	flowT "github.com/intel-go/nff-go/types"
	"lite5gc/cmn/rlogger"
	"sync/atomic"
)

func ReceiveArpHandler(pkt *packet.Packet, cxt flow.UserContext) bool {
	hc := cxt.(HandlerContext)

	pkt.ParseL3()
	protocol := pkt.Ether.EtherType

	if protocol == flowT.SwapARPNumber {
		arp := pkt.GetARPNoCheck()
		// Filtering large number of non local ARP packets in promiscuous mode
		if packet.SwapBytesUint16(arp.Operation) == packet.ARPRequest {
			return false
		}

		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "arp response,ether %s", pkt.Ether)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "arp response,l3 %s", arp)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "arp response,ARPReply %v", packet.SwapBytesUint16(arp.Operation))
		err := hc.Ipport.NeighCache.HandleIPv4ARPPacket(pkt)
		if err != nil {
			fmt.Println(err)
		}
		return false
	}
	atomic.AddUint64(&hc.Ipport.PacketCount, 1)
	atomic.AddUint64(&hc.Ipport.BytesCount, uint64(pkt.GetPacketLen()))
	return true
}
