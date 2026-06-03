package n11msg

import (
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
)

type N2InformationClass byte

const (
	SM_N2Info N2InformationClass = iota + 1 // N2 SM information.
	NRPPa                                   // N2 NRPPa information.
	PWS                                     // N2 PWS information of PWS type.
	PWS_BCAL                                // N2 Broadcast Completed Area List or the Broadcast Cancelled Area List.
	PWS_RF                                  // N2 Restart Indication or Failure Indication
	RAN                                     // N2 RAN related information.
)

//29518 6.1.6.2.15
type N2InfoContainerIE struct {
	//N2InformationClass  M
	//This IE represents the class of N2 information to be transferred.
	N2InforClass N2InformationClass

	//N2SmInformation C
	//This IE shall be present if session management N2 information is to be transferred.
	// When present, it represents a session management SMF related N2 information data part.
	SmInfo *N2SmInformation

	// ranInfo	N2RanInformation	C	0..1
	// This IE shall be present if RAN related N2 information is to be transferred
	// (i.e. n2InformationClass is "RAN"). When present, it shall contain the RAN related N2 information data part.
	RanInfo N2RanInformation

	//NrppaInformation C
	//This IE shall be present if location service related N2 information is to be transferred.
	//When present, it represents a NRPPa related N2 information data part.
	NrppaInfo NrppaInformation

	//PwsInformation C
	//This IE shall be present if PWS related N2 information is to be transferred.
	pwsInfor PwsInformation

	// Ie flags
	IeFlags bitset.BitSet
}

func (p *N2InfoContainerIE) SetN2SmInfo(data *N2SmInformation) {
	p.IeFlags.Set(Ieid_n2InfoContainer)
	p.SmInfo = data
}

func (p *N2InfoContainerIE) SetRanInfo(data N2RanInformation) {
	p.IeFlags.Set(Ieid_RanInfo)
	p.RanInfo = data
}

type N2SmInformation struct {
	//PduSessionId	M	1	Indicates the PDU Session Identity
	PduSessionId nas.PduSessID

	//N2InfoContent	C	0..1
	// This IE shall be present if a SMF related IE should be transferred.
	// When present, the IE contains one of NGAP SMF related IEs specified in
	// subclause 9.3.4 of 3GPP TS 38.413 [12].
	N2InfoCont *N2InfoContent

	//N1MessageContainer	C	0..1
	// This IE shall be present if NAS messages should be transfer between UE and SMF for session management function.
	// When present, it represents the NAS PDU to be carried over N2 between SMF and AN/UE.

	// nasPdu N1MessageContainer // removed in 29.518 f20 2018/12 version

	//Snssai	C	0..1
	// This IE shall be present if network slice information to be transferred for session management.
	// When present, the IE indicates the network slice the PDU session belongs to.
	Snssai nasie.SNssai

	//boolean	C	0..1
	// This IE shall be present if n2InfoContent carries a " Handover Required Transfer" IE.
	// When present, it Indicates whether the PDU session shall be subject to handover to the target node.
	SubjectToHo bool

	// Ie flags
	IeFlags bitset.BitSet
}

func (p *N2SmInformation) SetN2InfoContent(data *N2InfoContent) {
	p.IeFlags.Set(Ieid_n2InfoContent)
	p.N2InfoCont = data
}

type NrppaInformation struct {
	//NfInstanceId	M	1
	// This IE shall carry the identifier of the LMF that is sending or receiving the NRPPa data..
	lmfId string

	//N2InfoContent	M	1
	// This IE represents the encoded NGAP NRPPa-PDU IE, which is transparent to AMF.
	nrppaPdu N2InfoContent
}

type PwsInformation struct {
	//N2InfoContent	M	1
	// This IE represents the encoded NGAP NRPPa-PDU IE, which is transparent to AMF.
	messageIdentifier uint16

	//Uint16	M	1
	// identifies a particular message from the source and type indicated by the Message Identifier
	serialNumber uint16

	//N2InfoContent	M	1	This IE represents the PWS N2 information data part to be relayed between CBCF and AN.
	pwsContainer N2InfoContent

	//boolean	O	0..1
	// This IE shall be present if the AMF is expected to send the n2Information
	// it has received from the RAN nodes to the Service Consumer.
	// When present,it represents the value of Send-Broadcast-Completed-Area-List IE or
	// the Send-Broadcast-Cancelled-Area-List IE.
	sendN2Container bool

	//OmcIdentifier	O	0..1	This IE shall be present if the AMF is required to write the n2Information
	// it has received from the RAN nodes into trace records on the OMC.
	// When present, it indicates the identifier of OMC.
	omcId string
}

type N2InfoContent struct {
	//Uinteger	C
	//This IE shall be present if PWS related N2 information is to be transferred.
	//When present, it shall indicate the NGAP Message type of the ngapData.
	// Its value is numeric code of the NGAP Message Type IE type defined in ASN.1
	NgapMessageType uint

	//Uinteger	C
	//This IE shall be present if SM or NRPPa related N2 information is to be transferred.
	//When present, it shall indicate the NGAP IE type of the ngapData.
	// Its value is numeric code of the NGAP IE type defined in ASN.1.
	NgapIeType NgapIeType

	//RefToBinaryData	M	1
	// This IE reference the N2 information binary data corresponding to the N2 information class.
	// See subclause 6.1.6.4.3.
	NgapData []byte

	// Ie flags
	IeFlags bitset.BitSet
}

func (p *N2InfoContent) SetNgapIeType(data NgapIeType) {
	p.IeFlags.Set(Ieid_ngApIeType)
	p.NgapIeType = data
}

type N2RanInformation struct {
	// n2InfoContent	N2InfoContent	M	1
	// This IE shall contain the N2 RAN information to transfer.
	Content N2InfoContent
}
