package iprange

import (
	"fmt"
	"net"
	"testing"
)

func Test_Allocate_01(t *testing.T) {
	ipRange, err := ParseIPRange("10.55.0.0/29")

	//only ip/mask format supported
	_, net1, err := net.ParseCIDR("10.55.0.0/29")
	if err == nil {
		ipRange.End = GetEndIP(net1)
		ipRange.Start = GetStartIP(net1)

		ipAlloctor := NewAllocator(ipRange)

		fmt.Println("ip pool remaining: ", ipAlloctor.remaining)
		rip := net.ParseIP("10.55.0.3")
		fmt.Println("reserve ip: ", rip)
		ipAlloctor.Reserve(rip.To4())
		fmt.Println("ip pool remaining: ", ipAlloctor.remaining)

		for i := 0; i < 13; i++ {
			ueip := ipAlloctor.Allocate()
			fmt.Println(ueip)
			ipAlloctor.Release(ueip.To4())
		}
		//fmt.Println("ip pool remaining: ", ipAlloctor.remaining)
		//ueip := ipAlloctor.Allocate()
		//fmt.Println(ueip)

		fmt.Println("ip pool remaining: ", ipAlloctor.remaining)

	} else {

	}
}
