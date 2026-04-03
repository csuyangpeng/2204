package nasie

type PayloadContainerType byte

// 24.501 9.11.3.40
const (
	N1SmInformation       PayloadContainerType = 1
	SMSCont               PayloadContainerType = 2
	LPPMsg                PayloadContainerType = 3
	SORTransCont          PayloadContainerType = 4 //see subclauses   9.11.3.51
	UePolicyCont          PayloadContainerType = 5 //see subclause 9.11.3.53A
	UeParaUpdateTransCont PayloadContainerType = 6
	MultiplePayload       PayloadContainerType = 15 //multiple payloads
)

func (p PayloadContainerType) String() string {
	switch p {
	case N1SmInformation:
		return "N1 Sm Information"
	case SMSCont:
		return "SMS Container"
	case LPPMsg:
		return "LPP Message"
	case SORTransCont:
		return "SOR Trans Container"
	case UePolicyCont:
		return "Ue Policy Container"
	case UeParaUpdateTransCont:
		return "Ue Para Update Trans Container"
	case MultiplePayload:
		return "Multiple Payload"
	default:
		return ""
	}
}
