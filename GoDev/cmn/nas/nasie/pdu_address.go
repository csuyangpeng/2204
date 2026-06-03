package nasie

import (
	"fmt"
	"lite5gc/cmn/types3gpp"
	"net"
)

type PDUAddress struct {
	PDUSessionType  types3gpp.PduSessType
	PDUAddressInfor net.IP //todo
}

func (p PDUAddress) String() string {
	return fmt.Sprintf("PduSessType(%s), IP(%s)", p.PDUSessionType, p.PDUAddressInfor)
}
func (p *PDUAddress) GetPDUAddrInfor() net.IP {
	switch p.PDUSessionType {
	case types3gpp.Ipv4:
		return p.PDUAddressInfor.To4()
	case types3gpp.Ipv6:
	case types3gpp.Ipv4v6:
	default:
	}
	return net.IP{}
}

//refer to 9.11.4.10
//PDU address information (octet 4 to 15)
//If the PDU session type value indicates IPv4, the PDU address information in octet 4 to octet 7 contains an IPv4 address.
//If the PDU session type value indicates IPv6, the PDU address information in octet 4 to octet 11 contains an interface identifier for the IPv6 link local address. todo
//If the PDU session type value indicates IPv4v6, the PDU address information in octet 4 to octet 11 contains an interface identifier for the IPv6 link local address and in octet 12 to octet 15 contains an IPv4 address. //todo

//encode PDUAddress to nas octet stream
func (p *PDUAddress) Encode() ([]byte, error) {
	var encBuf []byte
	// PDUSessionType
	if p.PDUSessionType > 3 {
		return encBuf, fmt.Errorf("pdu session type value out of range")
	}
	encBuf = append(encBuf, byte(p.PDUSessionType))

	// PDUAddressInfor
	if p.PDUSessionType == types3gpp.Ipv4 {
		encBuf = append(encBuf, byte(p.PDUAddressInfor[12]))
		encBuf = append(encBuf, byte(p.PDUAddressInfor[13]))
		encBuf = append(encBuf, byte(p.PDUAddressInfor[14]))
		encBuf = append(encBuf, byte(p.PDUAddressInfor[15]))
	} else {
		return encBuf, fmt.Errorf("currently only support Ipv4 format")
	}

	return encBuf, nil
}
