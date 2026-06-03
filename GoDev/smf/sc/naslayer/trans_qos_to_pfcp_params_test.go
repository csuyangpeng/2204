package naslayer

import (
	"fmt"
	"lite5gc/cmn/nas/nasie"
	"testing"
)

func TestPduSessionModRejectMsg_Encode(t *testing.T) {
	c := nasie.PacketFilterContent{}
	c.PacketFilterContentID = []nasie.PacketFilterContentIdentifier{}
	//c.IPV4RemoteAddr.IP = []byte{0,0,0,0}
	//c.V4ProtocolIdOrV6NextHeader = 0
	//c.RemotePortRangeLowLimit = 3
	//c.RemotePortRangeHighLimit = 4
	sdfs, _ := TransFilterToSDF(c)
	fmt.Println(sdfs)
}
