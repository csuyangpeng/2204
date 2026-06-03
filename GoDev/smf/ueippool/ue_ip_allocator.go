package ueippool

import (
	"lite5gc/cmn/iprange"
	"net"
)

var smfUeIpAlloctor *iprange.IPRangeAllocator

func UeIpAlloctorInit(ipRange *iprange.IPRange) {
	smfUeIpAlloctor = iprange.NewAllocator(ipRange)
}

func UeIpReserve(ip net.IP) {
	smfUeIpAlloctor.Reserve(ip)
}

func UeIpReserveRange(ipRange *iprange.IPRange) {
	smfUeIpAlloctor.Subtract(ipRange)
}

func UeIpAllocate() net.IP {
	return smfUeIpAlloctor.Allocate()
}

func UeIpRelease(ip net.IP) {
	smfUeIpAlloctor.Release(ip)
}

func UeIpRemain() int64 {
	return smfUeIpAlloctor.Remaining()
}

func GetUeIpAllocator() *iprange.IPRangeAllocator {
	return smfUeIpAlloctor
}
