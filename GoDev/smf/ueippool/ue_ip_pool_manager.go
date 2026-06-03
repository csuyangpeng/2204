package ueippool

import (
	"lite5gc/cmn/iprange"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"net"
)

func SmfCreateUeIpPool(dnnKey string, iprangeValue string) {
	rlogger.FuncEntry(types.ModuleSmf, nil)
	// parse ip range
	ipRange, err := iprange.ParseIPRange(iprangeValue)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "failed to parse ip range in dnnIpRangeMap, "+
			"error(%s)", err)
		ipRange, _ = iprange.ParseIPRange(iprangeValue)
	}
	//only ip/mask format supported
	_, net1, err := net.ParseCIDR(iprangeValue)
	if err == nil {
		ipRange.End = iprange.GetEndIP(net1)
		ipRange.Start = iprange.GetStartIP(net1)

		rlogger.Trace(types.ModuleSmf, rlogger.INFO, nil,
			"dnn =", dnnKey, ", ipRange =", iprangeValue,
			", startIp=", ipRange.Start, ", endIp=", ipRange.End)
		// create ip range for dnn
		err = CreateUeIpPool(DnnNameKey(dnnKey), ipRange)
		if err != nil {
			rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil,
				"create ue ip pool failed.(dnn=%s,ip range=%s)", dnnKey, iprangeValue)
		}
	} else {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil,
			"parse ip range failed. (%s)", iprangeValue)
	}
	return
}
