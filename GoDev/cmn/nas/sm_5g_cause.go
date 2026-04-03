package nas

type Sm5gCause byte

//24.501 Table 9.11.4.2.1: 5GSM cause information element
//Cause value (octet 2)
//Bits
//8	7	6	5	4	3	2	1
//0	0	0	1	1	0	1	0		Insufficient resources
//0	0	0	1	1	0	1	1		Missing or unknown DNN
//0	0	0	1	1	1	0	0		Unknown PDU session type
//0	0	0	1	1	1	0	1		User authentication or authorization failed
//0	0	0	1	1	1	1	1		Request rejected, unspecified
//0	0	1	0	0	0	1	0		Service option temporarily out of order
//0	0	1	0	0	0	1	1		PTI already in use
//0	0	1	0	0	1	0	0		Regular deactivation
//0	0	1	0	0	1	1	1		Reactivation requested
//0	0	1	0	1	0	1	1		Invalid PDU session identity
//0	0	1	0	1	1	0	0		Semantic errors in packet filter(s)
//0	0	1	0	1	1	0	1		Syntactical error in packet filter(s)
//0	0	1	0	0	1	1	0		Out of LADN service area
//0	0	1	0	1	1	1	1		PTI mismatch
//0	0	1	1	0	0	1	0		PDU session type IPv4 only allowed
//0	0	1	1	0	0	1	1		PDU session type IPv6 only allowed
//0	0	1	1	0	1	1	0		PDU session does not exist
//0	1	0	0	0	0	1	1		Insufficient resources for specific slice and DNN
//0	1	0	0	0	1	0	0		Not supported SSC mode
//0	1	0	0	0	1	0	1		Insufficient resources for specific slice
//0	1	0	0	0	1	1	0		Missing or unknown DNN in a slice
//0	1	0	1	0	0	0	1		Invalid PTI value
//0	1	0	1	0	0	1	0		Maximum data rate per UE for user-plane integrity protection is too low
//0	1	0	1	0	0	1	1		Semantic error in the QoS operation
//0	1	0	1	0	1	0	0		Syntactical error in the QoS operation
//0	1	0	1	1	1	1	1		Semantically incorrect message
//0	1	1	0	0	0	0	0		Invalid mandatory information
//0	1	1	0	0	0	0	1		Message type non-existent or not implemented
//0	1	1	0	0	0	1	0		Message type not compatible with the protocol state
//0	1	1	0	0	0	1	1		Information element non-existent or not implemented
//0	1	1	0	0	1	0	0		Conditional IE error
//0	1	1	0	0	1	0	1		Message not compatible with the protocol state
//0	1	1	0	1	1	1	1		Protocol error, unspecified
//
// Any other value received by the UE shall be treated as 0010 0010, "service option temporarily out of order".
// Any other value received by the network shall be treated as 0110 1111, "protocol error, unspecified".

const (
	SuccessNoReason                                             Sm5gCause = 1
	OperatorDeterminedBarring                                   Sm5gCause = 0x08
	InsufficientResources                                       Sm5gCause = 0x1A
	MissingOrUnknownDNN                                         Sm5gCause = 0x1B
	UnknownPDUSessionType                                       Sm5gCause = 0x1C
	UserAuthenticationOrAuthorizationFailed                     Sm5gCause = 0x1D
	RequestRejectedUnspecified                                  Sm5gCause = 0x1F
	ServiceOptionNotSupported                                   Sm5gCause = 0x20
	RequestedServiceOptNotSubscribed                            Sm5gCause = 0x21
	ServiceOptionTemporarilyOutOfOrder                          Sm5gCause = 0x22
	PTIAlreadyInUse                                             Sm5gCause = 0x23
	RegularDeactivation                                         Sm5gCause = 0x24
	NetworkFailure                                              Sm5gCause = 0x26
	ReactivationRequested                                       Sm5gCause = 0x27
	InvalidPDUSessionIdentity                                   Sm5gCause = 0x2B
	SemanticErrorsInPacketFilters                               Sm5gCause = 0x2C
	SyntacticalErrorInPacketFilters                             Sm5gCause = 0x2D
	OutOfLADNServiceArea                                        Sm5gCause = 0x2E
	PTIMismatch                                                 Sm5gCause = 0x2F
	PDUSessionTypeIPv4OnlyAllowed                               Sm5gCause = 0x32
	PDUSessionTypeIPv6OnlyAllowed                               Sm5gCause = 0x33
	PDUSessionDoesNotExist                                      Sm5gCause = 0x36
	InsufficientResourcesForSpecificSliceAndDNN                 Sm5gCause = 0x43
	NotSupportedSSCMode                                         Sm5gCause = 0x44
	InsufficientResourcesForSpecificSlice                       Sm5gCause = 0x45
	MissingOrUnknownDNNInASlice                                 Sm5gCause = 0x46
	InvalidPTIValue                                             Sm5gCause = 0x51
	MaximumDataRatePerUEForUserPlaneIntegrityProtectionIsTooLow Sm5gCause = 0x52
	SemanticRrrorInTheQoSOperation                              Sm5gCause = 0x53
	SyntacticalRrrorInTheQoSOperation                           Sm5gCause = 0x54
	InvalidMappedEPSBearerIdentity                              Sm5gCause = 0x55
	SemanticallyIncorrectMessage                                Sm5gCause = 0x5F
	InvalidMandatoryInformation                                 Sm5gCause = 0x60
	MessageTypeNonExistentOrNotImplemented                      Sm5gCause = 0x61
	MessageTypeNotCompatibleWithTheProtocolState                Sm5gCause = 0x62
	InformationElementNonExistentOrNotImplemented               Sm5gCause = 0x63
	ConditionalIEError                                          Sm5gCause = 0x64
	MessageNotCompatibleWithTheProtocolState                    Sm5gCause = 0x65
	ProtocolErrorUnspecified                                    Sm5gCause = 0x6F
	OtherValue                                                  Sm5gCause = 0x6F
)
