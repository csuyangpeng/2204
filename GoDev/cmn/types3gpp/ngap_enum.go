package types3gpp

import "fmt"

//const defs in ngap messages
const (
	MaxNoOfTACs            = 256
	MaxNoOfBPLMNs          = 12
	MaxPagingDrx           = 63  //(0..63)
	MaxLenRanNodeName      = 150 //(1..150)
	MaxGNBIdBitLen         = 32  //(22..32)
	MinGNBIdBitLen         = 22
	MaxNoOfSliceItems      = 1024
	MaxRelativeAMFCapacity = 255 //(0..255)

	SizeofTAC       = 3
	SizeofSST       = 1
	SizeofSD        = 3
	SizeofGNBID     = 4
	SizeofNrCellID  = 5
	SizeofTimeStamp = 4
	//SizeofTransportLayAdd = 20
)

type DataForwardingNotPossible uint8

const (
	DataFwdNotPossible DataForwardingNotPossible = 0
)


func (p DataForwardingNotPossible) String() string {
	switch p {
	case DataFwdNotPossible:
		return "data_forwarding_not_possible"
	default:
		return "unknown"
	}
}

type PduSessType byte

const (
	Ipv4         PduSessType = 1
	Ipv6         PduSessType = 2
	Ipv4v6       PduSessType = 3
	Ethernet     PduSessType = 4
	Unstructured PduSessType = 5
)

func (p *PduSessType) StoreWithString(val string) error {
	switch val {
	case "ipv4":
		*p = Ipv4
	case "ipv6":
		*p = Ipv6
	case "ethernet":
		*p = Ethernet
	case "unstructured":
		*p = Unstructured
	case "ipv4v6":
		*p = Ipv4v6
	default:
		return fmt.Errorf("invalid pdu session types (%s)", val)
	}
	return nil
}

func (p *PduSessType) Convert2NgApType() PduSessType {
	return PduSessType(uint8(*p) - 1)
}

func (p PduSessType) String() string {
	switch p {
	case Ipv4:
		return "ipv4"
	case Ipv6:
		return "ipv6"
	case Ipv4v6:
		return "ipv4v6"
	case Ethernet:
		return "ethernet"
	case Unstructured:
		return "unstructured"
	default:
		return "unknown"
	}
}

type IntegrityProtectionInd byte

const (
	IntPrctIndRequired  IntegrityProtectionInd = 0
	IntPrctIndPreferred IntegrityProtectionInd = 1
	IntPrctIndNotNeeded IntegrityProtectionInd = 2
)

func (p IntegrityProtectionInd) String() string {
	switch p {
	case IntPrctIndRequired:
		return "required"
	case IntPrctIndPreferred:
		return "preferred"
	case IntPrctIndNotNeeded:
		return "not_needed"
	default:
		return "unknown"
	}
}

type ConfdProtectionInd byte

const (
	ConfdPrctIndRequired  ConfdProtectionInd = 0
	ConfdPrctIndPreferred ConfdProtectionInd = 1
	ConfdPrctIndNotNeeded ConfdProtectionInd = 2
)

func (p ConfdProtectionInd) String() string {
	switch p {
	case ConfdPrctIndRequired:
		return "required"
	case ConfdPrctIndPreferred:
		return "preferred"
	case ConfdPrctIndNotNeeded:
		return "not_needed"
	default:
		return "unknow"
	}
}

type MaxIntProtectedDataRate byte

const (
	BitRate64kbs  MaxIntProtectedDataRate = 0
	MaximumUeRate MaxIntProtectedDataRate = 1
)

func (p MaxIntProtectedDataRate) String() string {
	switch p {
	case BitRate64kbs:
		return "bit_rate_64kbs"
	case MaximumUeRate:
		return "max_ue_rate"
	default:
		return "unknonw"
	}
}

type ReflectQosAtt byte

const (
	SubjectTo ReflectQosAtt = 0
)

func (p ReflectQosAtt) String() string {
	switch p {
	case SubjectTo:
		return "subject_to"
	default:
		return "unknonw"
	}
}

type AddQosFlowInfo byte

const (
	MoreLikely AddQosFlowInfo = 0
)

func (p AddQosFlowInfo) String() string {
	switch p {
	case MoreLikely:
		return "morelikely"
	default:
		return "unknonw"
	}
}

type DelayCritical byte

const (
	DelayCri DelayCritical = iota
	NoDelayCri
)

func (p DelayCritical) String() string {
	switch p {
	case DelayCri:
		return "delay_critical"
	case NoDelayCri:
		return "no_delay_critical"
	default:
		return "unknonw"
	}
}

type NotificationControl byte

const (
	NotificationEnabled NotificationControl = 0
)

func (p NotificationControl) String() string {
	switch p {
	case NotificationEnabled:
		return "enabled"
	default:
		return "unknonw"
	}
}

// security key 256 bit
type SecurityKey [32]byte

type RRCInactTransReportReq byte

const (
	SubsStateTransRep RRCInactTransReportReq = iota
	SingleRrcconnStateRep
	CancelRep
)

type RedirectionForVoiceEPSFallback byte

const (
	Possible RedirectionForVoiceEPSFallback = iota
	NotPossible
)
