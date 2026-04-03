package nasie

import (
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/types3gpp"
	"net"
)

type DNNConfiguration struct {
	//mandatory
	Snssai *SNssai

	//Data Network Name
	DNN types3gpp.Apn

	//Default/Allowed session types
	DefaultPDUSessionType types3gpp.PduSessType
	AllowedSessionType    []types3gpp.PduSessType

	//Default/Allowed SSC modes
	DefaultSSCMode nas.SSCMode
	AllowedSSCMode []nas.SSCMode

	//optional
	//Indicates whether interworking with EPS is subscribed:
	//true: Subscribed; false: Not susubscribed;
	//If this attribute is absent it means not subscribed.
	IwkEpsInd bool

	//Indicates whether the DNN is a local area data network
	LadnIndicator bool

	//5G QoS parameters associated to the session for a data network
	QosProf QosProfile

	//The maximum aggregated uplink and downlink bit rates to be shared across
	// all Non-GBR QoS Flows in each PDU Session.
	SessionAmbr SessionAmbr

	//Subscribed charging characteristics data associated to the session for a data network
	ChargingCharacteristics string

	//Subscribed static IP address(es) of the IPv4 and/or IPv6 type
	StaticIpAddress net.IP

	//When present, this IE shall indicate the security policy for integrity protection and
	// encryption for the user plane.
	UpSecurityIntegr UpSecurityValue

	UpSecurityConfid UpSecurityValue
}

//func (p *DNNConfiguration) String() string {
//	return fmt.Sprintf("Snssai(%s), Dnn(%s), DefaultPDUSessionType(%s),"+
//		" DefaultSSCMode(%s),SessionAmbr(%s),StaticIpAddress(%s)",
//		p.Snssai,
//		p.DNN.String(),
//		p.DefaultPDUSessionType.String(),
//		p.DefaultSSCMode.String(),
//		p.SessionAmbr.String(),
//		p.StaticIpAddress.String())
//}

type UpSecurityValue uint8

const (
	Required  UpSecurityValue = 1
	Prefered  UpSecurityValue = 2
	NotNeeded UpSecurityValue = 3
)

func (p *UpSecurityValue) StoreWithString(val string) error {
	switch val {
	case "REQUIRED":
		*p = Required
	case "PREFERED":
		*p = Prefered
	case "NOT_NEEDED":
		*p = NotNeeded
	default:
		return fmt.Errorf("invalid up security value(%s)", val)
	}
	return nil
}
