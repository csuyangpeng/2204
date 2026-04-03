package nas

type Mm5gCause byte

// 24.501 9.11.3.2
//Cause value (octet 2)
//Bits
//8	7	6	5	4	3	2	1
//0	0	0	0	0	0	1	1		Illegal UE
//0	0	0	0	0	1	0	1		PEI not accepted
//0	0	0	0	0	1	1	0		Illegal ME
//0	0	0	0	0	1	1	1		5GS services not allowed
//0	0	0	0	1	0	0	1		UE identity cannot be derived by the network
//0	0	0	0	1	0	1	0		Implicitly de-registered
//0	0	0	0	1	0	1	1		PLMN not allowed
//0	0	0	0	1	1	0	0		Tracking area not allowed
//0	0	0	0	1	1	0	1		Roaming not allowed in this tracking area
//0	0	0	0	1	1	1	1		No suitable cells in tracking area
//0	0	0	1	0	1	0	0		MAC failure
//0	0	0	1	0	1	0	1		Synch failure
//0	0	0	1	0	1	1	0		Congestion
//0	0	0	1	0	1	1	1		UE security capabilities mismatch
//0	0	0	1	1	0	0	0		Security mode rejected, unspecified
//0	0	0	1	1	0	1	0		Non-5G authentication unacceptable
//0	0	0	1	1	0	1	1		N1 mode not allowed
//0	0	0	1	1	1	0	0		Restricted service area
//0	0	1	0	1	0	1	1		LADN not available
//0	1	0	0	0	0	0	1		Maximum number of PDU sessions reached
//0	1	0	0	0	0	1	1		Insufficient resources for specific slice and DNN
//0	1	0	0	0	1	0	1		Insufficient resources for specific slice
//0	1	0	0	0	1	1	1		ngKSI already in use
//0	1	0	0	1	0	0	0		Non-3GPP access to 5GCN not allowed
//0	1	0	0	1	0	0	1		Serving network not authorized
//0	1	0	1	1	0	1	0		Payload was not forwarded
//0	1	0	1	1	0	1	1		DNN not supported or not subscribed in the slice
//0	1	0	1	1	1	0	0		Insufficient user-plane resources for the PDU session
//0	1	0	1	1	1	1	1		Semantically incorrect message
//0	1	1	0	0	0	0	0		Invalid mandatory information
//0	1	1	0	0	0	0	1		Message type non-existent or not implemented
//0	1	1	0	0	0	1	0		Message type not compatible with the protocol state
//0	1	1	0	0	0	1	1		Information element non-existent or not implemented
//0	1	1	0	0	1	0	0		Conditional IE error
//0	1	1	0	0	1	0	1		Message not compatible with the protocol state
//0	1	1	0	1	1	1	1		Protocol error, unspecified
//
//Any other value received by the mobile station shall be treated as 0110 1111,
// "protocol error, unspecified". Any other value received by the network shall be
// treated as 0110 1111, "protocol error, unspecified".

const (
	//Customize
	SuccessAccept Mm5gCause = 1
	SystemFailure Mm5gCause = 2
	IdRequestNeed Mm5gCause = 255

	//Standard definition
	IllegalUE             Mm5gCause = 3
	PeiNotAccept          Mm5gCause = 5
	IllegalME             Mm5gCause = 6
	ServiceNotAllowed     Mm5gCause = 7
	UeIdCannotDerivedbyNw Mm5gCause = 9
	ImplicityDeregisted   Mm5gCause = 0xA
	PlmnNotAllowed        Mm5gCause = 0xB
	TANotAllowed          Mm5gCause = 0xC
	RoamingNotAllowedinTA Mm5gCause = 0xD
	NoSuitableCellsinTA   Mm5gCause = 0xF
	MacFailure            Mm5gCause = 0x14
	SynchFailure          Mm5gCause = 0x15
	Congestion            Mm5gCause = 0x16
	UeSecCapMismatch      Mm5gCause = 0x17
	SecModRejectUnspec    Mm5gCause = 0x18
	Non5gAuthUnaccept     Mm5gCause = 0x1A
	N1ModeNotAllowed      Mm5gCause = 0x1B
	RestrictedSerArea     Mm5gCause = 0x1C
	LandNotAvailable      Mm5gCause = 0x2B
	MaxPduSessReached     Mm5gCause = 0x41
	InsufResforSliceDnn   Mm5gCause = 0x43
	InsufResforSlice      Mm5gCause = 0x45
	NgksiInUsed           Mm5gCause = 0x47
	Non3gppAccNotAllowed  Mm5gCause = 0x48
	SerNwNotAuthorized    Mm5gCause = 0x49
	PayloadNotFw          Mm5gCause = 0x5A
	DnnNotSupport         Mm5gCause = 0x5B
	InsufUserPlanRes4Pdu  Mm5gCause = 0x5C
	SemIncorrectMsg       Mm5gCause = 0x5F
	InvalidMandInfo       Mm5gCause = 0x60
	MsgTypeNonExist       Mm5gCause = 0x61
	MsgTypeNotCompatble   Mm5gCause = 0x62
	IEnotExitOrImplt      Mm5gCause = 0x63
	ConditionalError      Mm5gCause = 0x64
	MsgNotCompatble       Mm5gCause = 0x65
	ProtocalError         Mm5gCause = 0x6F
)
