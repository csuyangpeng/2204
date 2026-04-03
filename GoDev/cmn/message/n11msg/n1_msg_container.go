package n11msg

//29.518 6.1.6.2.17
type N1MessageContainerIE struct {
	//mandatory IE
	//This IE shall contain the N1 message class for the message content specified in n1MessageContent.
	N1MsgClass N1MessageClass

	//This IE shall reference the N1 message binary data corresponding to the n1MessageClass.
	//This IE shall contain the value of the Content-ID header of the referenced binary body part.
	// See 3GPP TS 24.501 [11]. See subclause 6.1.6.4.2.
	N1MessageContent []byte

	//optional IE
	//This IE shall be present when the n1MessageClass IE is set to "LPP".
	//When present, this IE shall carry the identifier of the LMF sending or receiving the LPP data.
	LmfId string
}

type N1MessageClass byte

const (
	//The whole NAS message as received
	// (for e.g. used in forwarding the Registration message to target AMF
	// during Registration procedure with AMF redirection).
	FiveGMM N1MessageClass = 1

	//The N1 message of SM type
	SM_N1Info N1MessageClass = 2

	//The N1 message of LPP type.
	LPP N1MessageClass = 3

	//The N1 message of SMS type.
	SMS N1MessageClass = 4
)
