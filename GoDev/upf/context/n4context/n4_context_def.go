package n4context

import (
	"container/list"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/metric"
	"lite5gc/cmn/rlogger"
)

const moduleTag rlogger.ModuleTag = "n4context"

type N4SessionIDKey uint64

type KeyType uint8

const (
	//ImsiType      KeyType = 0
	//SessionIdType KeyType = 1

	// 3GPP TS 23.501 V15.3.0 (2018-09)
	// 5.8.2.11.2	N4 Session Context
	N4SessionIDCxtType KeyType = 2 //zj
)

// 3GPP TS 23.501 V15.3.0 (2018-09)
// 5.8.2.11.2	N4 Session Context
type N4SessionContext struct {
	SEID                     uint64
	SmfSEID                  pfcp.IEFSEID
	PDRs                     []*pfcp.IECreatePDR
	URRs                     []*pfcp.IECreateURR
	QERs                     []*pfcp.IECreateQER
	FARs                     []*pfcp.IECreateFAR
	BAR                      *pfcp.IECreateBAR
	CreateTrafficEndpoints   []*pfcp.IECreateTrafficEndpoint
	PDNType                  *pfcp.IEPDNType
	UserPlaneInactivityTimer *pfcp.IEUserPlaneInactivityTimer
	UserID                   *pfcp.IEUserID
	TraceInformation         *pfcp.IETraceInformation
	//	CN tunnel info.
	LocalFTEID map[uint16]*pfcp.IEFTEID
	//-	Network instance.
	NetworkInstance map[uint16]*pfcp.IENetworkInstance
	//-	QFIs.
	PDRQFIs map[uint16][]*pfcp.IEQFI
	//-	IP Packet Filter Set
	SDFFilters map[uint16][]*pfcp.IESDFFilter
	//Application Identifier
	ApplicationID map[uint16]*pfcp.IEApplicationID
	//Ethernet Packet Filter Set
	EthPacketFilters map[uint16][]*pfcp.IEEthernetPacketFilter

	// counter table
	MetricItems         metric.Registry // 16byte
	MetricItemsSnapshot metric.Registry

	// buffer
	Buffer *list.List // []byte
	//BufferFar pfcp.IECreateFAR
}

/// types.IGetAMFTraceFunc
/*func (p *N4SessionContext) ConstructUPFTraceObj() interface{} {
	if p == nil {
		return nil
	}

	traceObj := rlogger.CreateUPFTraceValueCtxt()

	traceObj.UPFSEID = types.Uint64Valid{V: p.SEID, Valid: true}
	traceObj.SMFSEID = types.Uint64Valid{V: p.SmfSEID.SEID, Valid: true}

	return traceObj
}
*/
