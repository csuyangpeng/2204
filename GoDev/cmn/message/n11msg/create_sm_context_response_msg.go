package n11msg

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
)

// 29.502  6.1.6.2.3	Type:SMContextCreatedData(响应)

type SMContextCreatedData struct {
	// hsmfUri		C
	// This IE shall be present in HR roaming scenarios if the additionalHsmfUri
	// IE was received in the request and the V-SMF established the PDU session
	// towards an alternative SMF listed in the additionalHsmfUri IE. When present,
	// it shall contain the URI of the H-SMF towards which the PDU session was established.
	HsmfUri string
	// PduSessionId	C
	// This IE shall be present, during an EPS to 5GS Idle mode mobility or handover using the N26 interface.
	// When present, it shall be set to the PDU Session ID.
	PduSessionId nas.PduSessID
	// Snssai	C	This IE shall be present during an EPS to 5GS Idle mode mobility or handover using the N26 interface.
	// When present, it shall contain the S-NSSAI assigned to the PDU session.
	SNssai nasie.SNssai
	// UpCnxState	C
	// This IE shall be present if the SMF was requested to activate
	// the user plane connection of the PDU session in the corresponding request.
	// When present, it shall be set as specified in subclause 5.2.2.2.2.
	// if the PDU session was requested to be re-activated,
	// i.e. if the PDU Session ID was present in the pduSessionsActivateList:
	//-	upCnxState attribute set to ACTIVATING;
	UpCnxState UpCnxState
	// n2SmInfo	RefToBinaryData	C
	// This IE shall be present if N2 SM Information needs to be sent to the AN.
	N2SmInfo string
	// N2SmInfoType	C
	// This IE shall be present if "n2SmInfo" attribute is present.
	// When present, this IE shall indicate the NGAP IE type
	// for the NG AP SMF related IE container carried in "n2SmInfo" attribute.
	N2SmInfoType N2SmInfoType
	// allocatedEbiList	C  1..N
	// This IE shall be present if the consumer
	// NF is an AMF and Inter-system mobility happens. When present, it shall contain
	// an array of EBI to ARP mappings currently allocated to the PDU session.
	AllocatedEbiList []int
	// HoState	C
	// This IE shall be present if the SMF was requested
	// to prepare an EPS to 5GS handover for the PDU session in the corresponding request.
	// When present, it shall be set as specified in subclause 5.2.2.2.3.
	HoState HoState
	// SupportedFeatures	C
	// This IE shall be present if at least one optional feature defined in subclause 6.1.8 is supported.
	SupportedFeatures string

	//smfServiceInstanceId	string	O	0..1
	//When present, this IE shall contain the serviceInstanceId of
	//the SMF service instance serving the PDU session Context.
	//This IE may be used by the AMF to identify PDU session contexts affected
	//by a failure or restart of the SMF service instance (see subclause 6.2 of 3GPP TS 23.527 [24]).
	SmfServiceInstanceId string

	//recoveryTime	DateTime	O	0..1
	//Timestamp when the SMF service instance serving the PDU session was (re)started
	//(see subclause 6.3 of 3GPP TS 23.527 [24]).
	recoveryTime string
}

//type N2SmInfoType struct {
//	EpsBearerId int
//	Arp         Arp
//}

type Arp struct {
	PriorityLevel int
	PreemptCap    string
	PreemptVuln   string
}

func (p SMContextCreatedData) N11MsgDataIf() {}
