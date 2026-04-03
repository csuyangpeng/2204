package nas

import (
	"fmt"
	"lite5gc/cmn/types3gpp"
)

// registration type definition
type RegistrationType byte

const (
	InitRegist          RegistrationType = 1
	MobRegistUpdating   RegistrationType = 2
	PeriodicRegUpdating RegistrationType = 3
	EmergencyRegist     RegistrationType = 4
)

func (p RegistrationType) String() string {
	var regType string
	switch p {
	case InitRegist:
		regType = fmt.Sprintf("InitRegist")
	case MobRegistUpdating:
		regType = fmt.Sprintf("MobRegistUpdating")
	case PeriodicRegUpdating:
		regType = fmt.Sprintf("PeriodicRegUpdating")
	case EmergencyRegist:
		regType = fmt.Sprintf("EmergencyRegist")
	default:
		regType = fmt.Sprintf("Unknown")
	}
	return regType
}

// registration result definition
type RegistrationResult byte

const (
	Access3gpp        RegistrationResult = 1
	AccessNon3gpp     RegistrationResult = 2
	Access3gppNon3gpp RegistrationResult = 3
)

const (
	MaxRestransNumT3550 uint8 = 4 //3gpp 24.501 5.5.1.2.8 c)
	MaxRestransNumT3592 uint8 = 4 //3gpp 24.501 6.3.3.5 a)
	MaxRestransNumT3570 uint8 = 4
)

// PDU Session Type Definition
//type PDUSessionType byte
//
//const (
//	IPv4         PDUSessionType = 0x01
//	IPv6         PDUSessionType = 0x02
//	IPv4v6       PDUSessionType = 0x03
//	Unstructured PDUSessionType = 0x04
//	Ethernet     PDUSessionType = 0x05
//	Reserved     PDUSessionType = 0x06
//)

//func (p PDUSessionType) String() string {
//	var psType string
//	switch p {
//	case IPv4:
//		psType = fmt.Sprintf("IPV4")
//	case IPv6:
//		psType = fmt.Sprintf("IPV6")
//	case IPv4v6:
//		psType = fmt.Sprintf("IPV4V6")
//	case Unstructured:
//		psType = fmt.Sprintf("UNSTR")
//	case Ethernet:
//		psType = fmt.Sprintf("ETHER")
//	default:
//		psType = fmt.Sprintf("Unknown")
//	}
//	return psType
//}

//func (p *PDUSessionType) StoreWithString(val string) error {
//	switch val {
//	case "IPV4":
//		*p = IPv4
//	case "IPV6":
//		*p = IPv6
//	case "ETHER":
//		*p = Ethernet
//	case "UNSTR":
//		*p = Unstructured
//	case "IPV4V6":
//		*p = IPv4v6
//	default:
//		return fmt.Errorf("invalid pdu session types (%s)", val)
//	}
//	return nil
//}

// PduSessionTypes
type PduSessionTypes struct {
	DefaultType     types3gpp.PduSessType
	AddAllowedTypes []types3gpp.PduSessType
}

// Ssc Mode definition
type SSCMode byte

const (
	SSCMode1   SSCMode = 0x01
	SSCMode2   SSCMode = 0x02
	SSCMode3   SSCMode = 0x03
	AsSSCMode1 SSCMode = 0x04
	AsSSCMode2 SSCMode = 0x05
	AsSSCMode3 SSCMode = 0x06
)

func (p SSCMode) String() string {
	var sscMode string
	switch p {
	case SSCMode1:
		sscMode = fmt.Sprintf("SSC_Mode_1")
	case SSCMode2:
		sscMode = fmt.Sprintf("SSC_Mode_2")
	case SSCMode3:
		sscMode = fmt.Sprintf("SSC_Mode_3")
	case AsSSCMode1:
		sscMode = fmt.Sprintf("As_SSC_Mode_1")
	case AsSSCMode2:
		sscMode = fmt.Sprintf("As_SSC_Mode_2")
	case AsSSCMode3:
		sscMode = fmt.Sprintf("As_SSC_Mode_3")
	default:
		sscMode = fmt.Sprintf("Unknown")
	}
	return sscMode
}
func (p *SSCMode) StoreWithString(val string) error {
	switch val {
	case "ssc mode 1":
		*p = SSCMode1
	case "ssc mode 2":
		*p = SSCMode2
	case "ssc mode 3":
		*p = SSCMode3
	default:
		return fmt.Errorf("invalid ssc mode(%s)", val)
	}
	return nil
}

type SSCModes struct {
	DefSscMode        SSCMode
	AddAllowedSscMode []SSCMode
}

// Max Data Rate definition
type IntergrityMaxDataRate struct {
	MaxRateUpLink   MAXDataRate
	MaxRateDownLink MAXDataRate
}

func (i IntergrityMaxDataRate) Reset() {
	i.MaxRateDownLink = KBPS_64
	i.MaxRateUpLink = KBPS_64
}

type MAXDataRate byte

const (
	KBPS_64      MAXDataRate = 0x00
	FullDataRate MAXDataRate = 0xFF
)

type SMCapability struct {
	MPTCP  bool // MPTCP functionality supported or not
	ATSLL  bool // ATSSS Low-Layer functionality supported or not
	EPTS1  bool // Ethernet PDN type in S1 mode supported or not
	MH6PDU bool // for IPV6
	RqoS   bool // if the UE supports reflective QoS, the UE should set the RQoS bit to "Reflective QoS supported" true
	Spare  []byte
}

func (s SMCapability) Reset() {
	s.RqoS = false
	s.MH6PDU = false
	s.Spare = []byte{}
}

const (
	MinNumOfSPF                  = 17
	MaxNumOfSPF                  = 1024
	MinNumOfSpareInSMCapability  = 1
	MaxNumOfSpareInSMCapability  = 12
	MinOfExtendProtocolConfigOpt = 4
	MaxOfExtendProtocolConfigOpt = 65538
)

type ServiceTYPE byte

//0	0	0	0	signalling
//0	0	0	1	data
//0	0	1	0	mobile terminated services
//0	0	1	1	emergency services
//0	1	0	0	emergency services fallback
//0	1	0	1	high priority access
//0	1	1	0	unused; shall be interpreted as "signalling", if received by the network
//0	1	1	1	unused; shall be interpreted as "signalling", if received by the network
//1	0	0	0	unused; shall be interpreted as "signalling", if received by the network
//1	0	0	1	unused; shall be interpreted as "data", if received by the network
//1	0	1	0	unused; shall be interpreted as "data", if received by the network
//1	0	1	1	unused; shall be interpreted as "data", if received by the network

const (
	Signalling                ServiceTYPE = 0
	Data                      ServiceTYPE = 1
	MobileTerminatedServices  ServiceTYPE = 2
	EmergencyServices         ServiceTYPE = 3
	EmergencyServicesFallback ServiceTYPE = 4
	HighPriorityAccess        ServiceTYPE = 5
)
