package nasie

type PduSessRequestType uint8

// 24.501 9.11.3.47
//Request type value (octet 1, bit 1 to bit 4)
//
//Bits
//3	2	1
//0	0	1		initial request
//0	1	0		existing PDU session
//0	1	1		initial emergency request
//1	0	0		existing emergency PDU session
//1	0	1		modification request
//1	1	1		reserved
//
//All other values are unused and shall be interpreted as "initial request", if received by the network.
const (
	InitialRequest    PduSessRequestType = 1
	ExistingPduSess   PduSessRequestType = 2
	InitialEmergency  PduSessRequestType = 3
	ExistingEmergency PduSessRequestType = 4
	ModifyRequest     PduSessRequestType = 5
	ReservedReqType   PduSessRequestType = 7
)

func (p PduSessRequestType) String() string {
	switch p {
	case InitialRequest:
		return "Initial Request"
	case ExistingPduSess:
		return "Existing Pdu Sess"
	case InitialEmergency:
		return "Initial Emergency"
	case ExistingEmergency:
		return "Existing Emergency"
	case ModifyRequest:
		return "Modify Reqeust"
	case ReservedReqType:
		return "reserveded Req Type"
	default:
		return ""
	}
}
