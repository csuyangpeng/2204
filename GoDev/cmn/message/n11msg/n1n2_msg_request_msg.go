package n11msg

import (
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/types3gpp"
)

// 3GPP TS 29.518 V15.1.0 (2018-09)
// 6.1.6.2.18	Type: N1N2MessageTransferReqData
type N1N2MessageTransferReqData struct {
	// N1MessageContainer	C
	// This IE shall be included if a N1 message needs to be transferred.
	N1MessageContainer *N1MessageContainerIE

	// N2InfoContainer	C
	// This IE shall be included if a N2 information needs to be transferred.
	N2InfoContainer *N2InfoContainerIE

	// skipInd		C
	// This IE shall be present and set to "true" if the service consumer (e.g. SMF)
	// requires the N1 message to be sent to the UE only when UE is in CM-CONNECTED,
	// e.g. during SMF initiated PDU session release procedure (see subclause 4.3.4.2 of 3GPP TS 23.502 [2]).
	// When present, this IE shall be set as following:
	// -	true: AMF should skip sending N1 message to UE, when the UE is in CM-IDLE.
	// -	false (default): the AMF shall send the N1 message to the UE.
	SkipInd bool

	// lastMsgIndication		O
	// This flag when present shall indicate that the message transferred is the last message.
	// (See subclause 4.13.3.3 of 3GPP TS 23.502 [2].
	LastMsgIndication bool

	// pduSessionId	PduSessionId	O	0..1
	// PDU Session ID for which the N1 / N2 message is sent, if the N1 / N2 message class is SM.
	SessionId nas.PduSessID

	// lcsCorrelationId	CorrelationID	O	0..1
	// LCS Correlation ID, for which the N1 message is sent, if the N1 message class is LPP
	// (see subclause 4.13.5.4 of 3GPP TS 23.502 [2]).
	LcsCorrelationId string

	// Ppi	O
	// This IE when present shall indicate the Paging policy to be applied.
	// The paging policies are configured at the AMF.
	Ppi int

	// arp	Arp	O
	// This IE when present shall indicate the Allocation and
	// Retention Priority of the PDU session for which the N1/N2 message transfer is initiated.
	// This IE shall not be present when the N1/N2 message class is not SM.
	Arp nasie.ARP

	// 5Qi	O
	// This IE when present shall indicate the 5QI associated with the PDU session
	// for which the N1 / N2 message transfer is initiated.
	// This IE shall not be present when the N1/N2 message class is not SM.
	Qi5 nasie.QI5Contents

	// n1n2FailureTxfNotifURI		O
	// If included, this IE represents the callback URI on which the AMF shall notify the N1/N2 message transfer failure.
	N1n2FailureTxfNotifURI string

	// SupportedFeatures	C
	// This IE shall be present if at least one optional feature defined in subclause 6.1.8 is supported.
	SupportedFeatures string

	// smfReallocationInd	boolean	O	0..1
	// This IE shall indicate that the SMF is requested to be reallocated
	// (see subclause 4.3.5.2 of 3GPP TS 23.502 [2]).
	//When present, this IE shall be set as follows:
	//-	true: the SMF is requested to be reallocated.
	//-	false (default): the SMF is not requested to be reallocated.
	SmfReallocationInd bool

	//areaOfValidity	AreaOfValidity	O	0..1
	// This IE represents the list of TAs where the provided N2 information is valid.
	// See subclause 5.2.2.2.7 and 4.2.3.3 of 3GPP TS 23.502 [3].
	AreaOfValidity AreaOfValidity

	// Ie flags
	IeFlags bitset.BitSet
}

type AreaOfValidity struct {
	// taiList	array(Tai)	M	0..N
	// An array of TAI representing the area of validity of the associated N2 information provided.
	TaiList []types3gpp.TAI
}

func (p *N1N2MessageTransferReqData) SetN1MsgContainer(data *N1MessageContainerIE) {
	p.IeFlags.Set(Ieid_n1MessageContainer)
	p.N1MessageContainer = data
}

func (p *N1N2MessageTransferReqData) SetN2InfoContainer(data *N2InfoContainerIE) {
	p.IeFlags.Set(Ieid_n2InfoContainer)
	p.N2InfoContainer = data
}

func (p *N1N2MessageTransferReqData) SetSkipInd(data bool) {
	p.IeFlags.Set(Ieid_skipInd)
	p.SkipInd = data
}

func (p *N1N2MessageTransferReqData) SetSessionId(data nas.PduSessID) {
	p.IeFlags.Set(Ieid_pdusessionId)
	p.SessionId = data
}

// todo add the rest ies
func (p N1N2MessageTransferReqData) N11MsgDataIf() {}

func (n11Data N1N2MessageTransferReqData) String() string {
	return fmt.Sprintf("",
		"N1MessageContainer.N1MsgClass: ", n11Data.N1MessageContainer.N1MsgClass,
		"N1MessageContainer.N1MessageContent: ", n11Data.N1MessageContainer.N1MessageContent,
		"SessionId: ", n11Data.SessionId,
		"N2InfoContainer.N2InforClass: ", n11Data.N2InfoContainer.N2InforClass,
		"N2InfoContainer.SmInfo.N2InfoCont.NgapData: ", n11Data.N2InfoContainer.SmInfo.N2InfoCont.NgapData,
		"N2InfoContainer.SmInfo.N2InfoCont.NgapIeType: ", n11Data.N2InfoContainer.SmInfo.N2InfoCont.NgapIeType)
}
