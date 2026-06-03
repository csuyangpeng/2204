package pfcp

import (
	"fmt"
	"net"
	"testing"
)

func TestA(t *testing.T) {
	vIE := &IEActivatePredefinedRules{}
	vIE.Set("active a")

	fmt.Println(*vIE)

	var MACAddrValue net.HardwareAddr
	MACAddrValue = net.HardwareAddr{1, 2, 3, 4, 5, 6}
	fmt.Println(len(MACAddrValue), MACAddrValue)
}
