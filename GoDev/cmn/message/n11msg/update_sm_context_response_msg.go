package n11msg

import (
	"github.com/willf/bitset"
)

//3GPP TS 29.502 V15.1.0(2018-09)
// 6.1.6.2.5	Type: SMContextUpdatedData
type UpdateSMContextResponseData struct {
	// UpCnxState	C
	// This IE shall be present if the SMF was requested
	// to activate or deactivate the user plane connection of the PDU session in the corresponding request.
	// When present, it shall be set as specified in subclause 5.2.2.3.2.
	UpCnxState UpCnxState
	// hoState	HoState	C
	// This IE shall be present if the SMF was requested to prepare,
	// execute or cancel a handover for the PDU session in the corresponding request.
	// When present, it shall be set as specified in subclause 5.2.2.3.4.
	HoState HoState
	// releaseEbiList	array(EpsBearerId)	C	1..N
	// This IE shall be present if the SMF determines that some EBIs are not needed. When present,
	// it shall contain the EBIs to be released.
	ReleaseEbiList []int
	// allocatedEbiList	array(EbiArpMapping)	C	1..N
	// This IE shall be present if the consumer NF is an AMF and Inter-system mobility happens.
	// When present, it shall contain an array of EBI to ARP mappings currently allocated to the PDU session.
	AllocatedEbiList []N2SmInfoType
	// modifiedEbiList	array(EbiArpMapping)	C	0..N
	// This IE shall be present if a PDU session modification procedure resulted
	// in the change of ARP for a QoS flow that was already allocated an EBI.
	ModifiedEbiList []N2SmInfoType
	// n1SmMsg	RefToBinaryData	C
	// This IE shall be present if N1 SM Information needs to be sent to the UE.
	// When present, this IE shall reference the N1 SM Message binary data (see subclause 6.1.6.4.2).
	N1SmMsg string
	// n2SmInfo	RefToBinaryData	C
	// This IE shall be present if N2 SM Information needs to be sent to the AN.
	// When present, this IE shall reference the N2 SM Information binary data (see subclause 6.1.6.4.3).
	N2SmInfo string
	// n2SmInfoType	N2SmInfoType	C
	// This IE shall be present if "n2SmInfo" attribute is present.
	// When present, this IE shall indicate the NGAP IE type
	// for the NG AP SMF related IE container carried in "n2SmInfo" attribute.
	N2SmInfoType N2SmInfoType
	// epsBearerSetup	array(EpsBearerContainer)	C	1..N
	// This IE shall be present during an EPS to 5GS handover using the N26 interface.
	// When present, it shall include the EPS bearer context(s) successfully handed over to 5GS.
	EpsBearerSetup []string
	// dataForwarding	boolean	C
	// This IE shall be present if it was present in the corresponding request.
	// When present, it shall be set as specified in subclause 5.2.2.3.9.
	DataForwarding bool

	// Ie flags
	IeFlags bitset.BitSet
}

func (p UpdateSMContextResponseData) N11MsgDataIf() {}
