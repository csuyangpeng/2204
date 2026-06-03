package session

import (
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"net"
)

// 3GPP TS 29.244 V15.3.0 (2018-09)
// 5.6.2	Session Endpoint Identifier Handling
// 8.2.37	F-SEID Fully Qualified SEID
/*type FSEID struct {
	V4Flag uint8 // V4:2,V6:1
	SEID     uint64
	IPv4Addr net.IP // 4 byte
	IPv6Addr net.IP // 16 byte
}*/

type FSEID = pfcp.IEFSEID

// allocation of F-SEID
func NewFSEID(seid uint64, ipType uint8, ip net.IP) *FSEID {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	var FSeid *FSEID
	if ipType == pfcp.IEFSEID_IPv4_address {
		FSeid = &FSEID{
			V4Flag:   ipType,
			SEID:     seid,
			IPv4Addr: ip,
			IPv6Addr: nil,
		}
	} else if ipType == pfcp.IEFSEID_IPv6_address {
		FSeid = &FSEID{
			V4Flag:   ipType,
			SEID:     seid,
			IPv4Addr: ip,
			IPv6Addr: nil,
		}
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Fully Qualified SEID:%v", *FSeid)

	return FSeid
}
