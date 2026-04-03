package ueippool

import (
	"fmt"
	"lite5gc/cmn/iprange"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/syncmap"
	"lite5gc/cmn/types"
	"net"
)

//key:dnn name,value:ip range allocator
var ueIpPoolTable syncmap.SyncMap

type DnnNameKey string

func CreateUeIpPool(key interface{}, ipRange *iprange.IPRange) error {
	rlogger.FuncEntry(types.ModuleSmf, nil)
	var err error
	switch key.(type) {
	case DnnNameKey:
		ipAllocator := iprange.NewAllocator(ipRange)
		err = ueIpPoolTable.Set(key, ipAllocator)
		if err != nil {
			err = fmt.Errorf("failed to create the ip allocator for dnn(%s), error(%s)", key.(DnnNameKey), err)
		}
	default:
		err = fmt.Errorf("invalid key")
	}

	return err
}

func GetUeIpPool(key interface{}) (*iprange.IPRangeAllocator, error) {
	rlogger.FuncEntry(types.ModuleSmf, nil)

	switch key.(type) {
	case DnnNameKey:
		val := ueIpPoolTable.Get(key)
		if val == nil {
			return nil, fmt.Errorf("failed to find Ip Allocator with dnn key(%s)", key.(DnnNameKey))
		}
		allocator, ok := val.(*iprange.IPRangeAllocator)
		if !ok {
			return nil, fmt.Errorf("failed to convert ip allocator with dnn key(%s)", key.(DnnNameKey))
		}
		return allocator, nil

	default:
		return nil, fmt.Errorf("invalid key")
	}
}

func AllocateByDNN(dnn DnnNameKey) (ip net.IP, err error) {
	rlogger.FuncEntry(types.ModuleSmf, nil)
	allocator, err := GetUeIpPool(dnn)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to get allocator by dnn: %s", dnn)
		return nil, fmt.Errorf("fail to get allocator by dnn: %s", dnn)
	}
	return allocator.Allocate(), nil
}

func ReleaseIPByDNN(dnn DnnNameKey, ip net.IP) (err error) {
	rlogger.FuncEntry(types.ModuleSmf,nil)
	allocator, err := GetUeIpPool(dnn)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to get allocator by dnn: %s", dnn)
		return fmt.Errorf("fail to get allocator by dnn: %s", dnn)
	}
	allocator.Release(ip.To4())
	return nil
}

func UpdateUeIpRange(key interface{}, ipRange *iprange.IPRange) error {
	rlogger.FuncEntry(types.ModuleSmf, nil)
	var err error
	switch key.(type) {
	case DnnNameKey:
		ipAlloctor := iprange.NewAllocator(ipRange)
		ueIpPoolTable.Update(key, ipAlloctor)
	default:
		err = fmt.Errorf("invalid key")
	}

	return err
}

func UpdateUeIpPool(key interface{}, pool *iprange.IPRangeAllocator) error {
	rlogger.FuncEntry(types.ModuleSmf, nil)
	if pool == nil {
		return fmt.Errorf("invalid input parameter, nil pool")
	}

	switch key.(type) {
	case DnnNameKey:
		ueIpPoolTable.Update(key, pool)
	default:
		return fmt.Errorf("invalid key")
	}

	return nil
}

func DeleteUeIpPool(key interface{}) error {
	rlogger.FuncEntry(types.ModuleSmf, nil)

	switch key.(type) {
	case DnnNameKey:
		ueIpPoolTable.Del(key)
	default:
		return fmt.Errorf("invalid key")
	}

	return nil
}

func LengthOfUeIpPoolTbl() uint64 {
	return ueIpPoolTable.Length64()
}
