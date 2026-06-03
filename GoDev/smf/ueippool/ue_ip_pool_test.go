package ueippool

import (
	"fmt"
	"lite5gc/cmn/iprange"
	"testing"
)

func Test1(t *testing.T) {
	var alloctor *iprange.IPRangeAllocator
	ipr, _ := iprange.ParseIPRange("10.55.0.0-255.254/16")
	alloctor = iprange.NewAllocator(ipr)
	fmt.Println(alloctor.Allocate())
}

func Test2(t *testing.T) {
	//dnn1 := "cmnet"
	//dnn2 := "5gnet"
	//
	//d1 := DnnNameKey(dnn1)
	//d2 := DnnNameKey(dnn2)
	//
	iprange1 := "10.55.0.22/16"
	iprange2 := "10.55.0.2/16"
	//iprange1 := "10.55.0.1/16"
	//iprange2 := "10.2.2.11/16"
	//b,err := iprange.OverlappingSubnets(iprange1,iprange2)
	//fmt.Println(b,err)
	r1, err := iprange.ParseIPRange(iprange1)
	if err != nil {
		fmt.Println("read ip range 1 err: ", err)
	}
	r2, err := iprange.ParseIPRange(iprange2)
	if err != nil {
		fmt.Println("read ip range 2 err: ", err)
	}
	fmt.Println(r1.Overlaps(r2))
	//err = CreateUeIpPool(d1, r1)
	//if err != nil {
	//	fmt.Println("get ip pool 1 err: ", err)
	//}
	//err = CreateUeIpPool(d2,r2)
	//if err != nil{
	//	fmt.Println("get ip pool 2 err: ",err)
	//}
	//

	//for i := 0; i < 11; i++ {
	//	a1, err := GetUeIpPool(d1)
	//	if err != nil {
	//		fmt.Println("fail to get ip from ip pool 1", err)
	//	}
	//	fmt.Println("ip1 = ", a1.Allocate())
	//}

	//
	//a2, err := GetUeIpPool(d2)
	//if err != nil{
	//	fmt.Println("fail to get ip from ip pool 2",err)
	//}
	//fmt.Println("ip2 = ",a2.Allocate())

	//UpdateUeIpPool(DnnNameKey("d2"),a1)
	//a2, err := GetUeIpPool(DnnNameKey("d2"))
	//fmt.Println(a2.Allocate())
	//
	//DeleteUeIpPool(DnnNameKey("d2"))
	//_, err = GetUeIpPool(DnnNameKey("d2"))
	//if err != nil{
	//	fmt.Println(err)
	//}

	//fmt.Println(LengthOfUeIpPoolTbl())
}

func Test3(t *testing.T) {
	var kv map[string]string
	kv = make(map[string]string)
	kv["1"] = "10.50.0.2-255.254/16"
	kv["2"] = "10.51.0.1/16"
	//kv["3"] = "10.52.0.2/16"
	//kv["4"] = "10.50.0.3/16"
	var valueSlice []string
	for _, v := range kv {
		valueSlice = append(valueSlice, v)
	}
	for i := 0; i < len(kv); i++ {
		for j := i + 1; j < len(kv); j++ {
			overlapping, err := iprange.OverlappingSubnets(valueSlice[i], valueSlice[j])
			fmt.Println(i, j, valueSlice[i], valueSlice[j], overlapping, err)
			//fmt.Println(i,j,valueSlice[i],valueSlice[j],";")
		}
		fmt.Println()
	}
}
