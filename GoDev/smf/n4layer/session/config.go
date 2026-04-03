package session

import (
	"lite5gc/cmn/types3gpp"
)

var PDRID uint16 = 1

// the higher precedence values indicate lower precedence of the PDR when matching a packet.
var PrecedenceValue uint32 = 1

var SequenceNumber uint32 = 1

var LocalIp string = "127.0.0.1"

var UpfIpN4Port string = "10.180.8.47:8805"

var UpfIpN4 string = "10.180.8.47"

var UeIp string = "192.0.0.1"

//The Traffic Endpoint ID value shall be encoded as a binary integer value
// within the range of 0 to 255.
var TrafficEndpointIDValue uint8 = 1

var NetworkInstance types3gpp.Apn = types3gpp.Apn{}

var IPFilterRule string = "Action:permit,Direction:out,Protocol:ip,Source IP address:any, Destination IP address:assigned"

var ApplicationID string = "app1"

//The bit 8 of octet 5 is used to indicate if the Rule ID is dynamically allocated by the CP function or predefined in the UP function.
// If set to 0, it indicates that the Rule is dynamically provisioned by the CP Function. If set to 1,
// it indicates that the Rule is predefined in the UP Function.
var FARID uint32 = 1 // smf 动态分配，小于等于0x0FFFFFFF
var URRID uint32 = 0 // smf 动态分配，小于等于0x0FFFFFFF
var QERID uint32 = 0 // smf 动态分配，小于等于0x0FFFFFFF

var ANTunnelInfo types3gpp.Teid = 1
var SEID uint64 = 1

//func getCxtFromUeContext(smfUeCxt *smfcontext.UeContext, psi nas.PduSessID) {
//configure.SmfConf
//if len(smfUeCxt.GetPduSessCtxt(psi).PDR) < 2 {
//	rlogger.Trace(types.SmfN4Layer, rlogger.ERROR, smfUeCxt, "Input parameter PDR > 2 check failed")
//	return
//}
//PDRID = smfUeCxt.GetPduSessCtxt(psi).PDR[0].PDRID.RuleID
//PrecedenceValue = smfUeCxt.GetPduSessCtxt(psi).PDR[0].Precedence.PrecedenceValue
//LocalIp = smfUeCxt.GetPduSessCtxt(psi).UEIP.String()
//
//TrafficEndpointIDValue = smfUeCxt.GetPduSessCtxt(psi).PDR[0].PDI.TrafficEndpointID.Value
//NetworkInstance = smfUeCxt.GetPduSessCtxt(psi).DNN
//IPFilterRule = string(smfUeCxt.GetPduSessCtxt(psi).PDR[0].PDI.SDFFilters[0].FlowDescription)
//ApplicationID = string(smfUeCxt.GetPduSessCtxt(psi).PDR[1].PDI.ApplicationID.ApplicationIdentifier)

//}
