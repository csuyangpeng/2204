package n11msg

//3GPP TS 29.518 V15.1.0 (2018-09)
//6.1.6.2.19	Type: N1N2MessageTransferRspData
type N1N2MessageTransferRspData struct {
	// cause	N1N2MessageTransferCause	M
	// This IE shall provide the result of the N1/N2 message transfer processing at the AMF.
	Cause N1N2MessageTransferCause
	// SupportedFeatures	C
	// This IE shall be present if at least one optional feature defined in subclause 6.1.8 is supported.
	SupportedFeatures string
}

func (p N1N2MessageTransferRspData) N11MsgDataIf() {}

type N1N2MessageTransferCause byte

// 6.1.6.3.6	Enumeration: N1N2MessageTransferCause
const (
	// This cause code represents the case where the AMF has initiated paging to reach the UE in order to deliver the N1 message.
	ATTEMPTING_TO_REACH_UE N1N2MessageTransferCause = 1
	// This cause code represents the case where the AMF has initiated the N1/N2 message transfer towards the UE and/or the AN.
	N1_N2_TRANSFER_INITIATED N1N2MessageTransferCause = 2

	// This cause code represents the case where the AMF has stored the N1/N2 message due to Asynchronous Transfer.
	WAITING_FOR_ASYNCHRONOUS_TRANSFER N1N2MessageTransferCause = 3

	// This cause code represents the case that the AMF has initiated paging to reach the UE but the UE is not responding to the paging.
	UE_NOT_RESPONDING N1N2MessageTransferCause = 4

	// This cause code represents the case where the AMF has skipped sending N1 message to the UE,
	// when UE is in CM-IDLE and the "skipInd" is set to "true" in the request.
	N1_MSG_NOT_TRANSFERRED N1N2MessageTransferCause = 5

	// This cause code indicates that the UE is not reachable for the non-3GPP PDU session,
	// due to the UE being in CM-IDLE for non-3GPP access and the PDU session is not allowed to move to 3GPP access.
	UE_NOT_REACHABLE_FOR_SESSION N1N2MessageTransferCause = 6
)
