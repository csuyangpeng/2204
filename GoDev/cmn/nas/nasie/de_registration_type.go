package nasie

//24501 9.11.3.20
//Switch off (octet 1, bit 4)
//In the UE to network direction:
//Bit
//4
//0				Normal de-registration
//1				Switch off
//In the network to UE direction bit 4 is spare. The network shall set this bit to zero.
//
//Re-registration required (octet 1, bit 3)
//In the network to UE direction:
//Bit
//3
//0				re-registration not required
//1				re-registration required
//In the UE to network direction bit 3 is spare. The UE shall set this bit to zero.
//
//Access type (octet 1,bit 2, bit 1)
//Bit
//2	1
//0	1			3GPP access
//1	0			Non-3GPP access
//1	1			3GPP access and non-3GPP access
//All other values are reserved.

type DeRegistrationTypeIE struct {
	SwithOff                 bool
	IsReRegistrationRequired bool
	AccessType               AccessTypes
}

type AccessTypes byte

const (
	ThreeGppAccess       AccessTypes = 1
	Non3gppAccess        AccessTypes = 2
	Access3gppAndNon3gpp AccessTypes = 3
)

func (p *DeRegistrationTypeIE) Reset() {
	p.SwithOff = false
	p.IsReRegistrationRequired = false
	p.AccessType = 0
}
