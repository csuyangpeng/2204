package session

import (
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"net"
)

const Allocation_CN_Tunnel_Info_IN_SMF = 1

// 3GPP TS 23.501 V15.3.0 (2018-09)
// 5.8.2.3	Management of CN Tunnel Info in the SMF
// It comprises the TEID and the IP address which is used by the UPF for the PDU Session.
type CNTunnelInfo struct {
	TEID      types3gpp.Teid
	UpfIpAddr net.IP //[]byte ,upf N3 Ip address
}

// allocation of CN Tunnel Info
func NewCNTunnelInfo(teid types3gpp.Teid, ip net.IP) *CNTunnelInfo {
	rlogger.FuncEntry(types.ModuleSmfN4, &logger.TraceV{types.TEID_TRACE, uint32(teid)})
	tunnelInfo := &CNTunnelInfo{
		TEID:      teid,
		UpfIpAddr: ip,
	}

	return tunnelInfo
}

// release of CN Tunnel Info
func (this *CNTunnelInfo) Release() error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	this = nil
	return nil
}

// SMF assign the Local F-TEID
// uint32 类型，核心网内唯一
func GetTEID() types3gpp.Teid {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	idmgr.GetInst().RegisterIDMgr(string(types.SmfTEID), types.MaxNumSmfTeid)
	teid, err := idmgr.GetInst().BorrowID(string(types.SmfTEID))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "SMF assign the Local TEID failed:%s", err)
	}
	return types3gpp.Teid(teid)
}
