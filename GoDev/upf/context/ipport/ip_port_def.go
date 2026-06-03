package ipport

import (
	"github.com/intel-go/nff-go/packet"
	"github.com/intel-go/nff-go/types"
)

type NetworkSubnet struct {
	IPv4 types.IPv4Subnet `json:"ipv4"`
	IPv6 types.IPv6Subnet `json:"ipv6"`
}

type IpPort struct {
	Index  uint16        `json:"index"`
	Subnet NetworkSubnet `json:"subnet"`
	//static dst mac
	DstMacAddress types.MACAddress `json:"dst-mac"`

	StaticARP  bool // 配置DstMacAddress，不用查询arp
	NeighCache *packet.NeighboursLookupTable
	// self mac
	MacAddress types.MACAddress

	PacketCount uint64
	BytesCount  uint64
}

// n6 context
