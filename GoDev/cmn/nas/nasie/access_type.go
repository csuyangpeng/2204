package nasie

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type AccessType uint8

//Access type value (octet 1, bit 1 to bit 2)
//
//Bits
//2	1
//0	1			3GPP access
//1	0			Non-3GPP access
//
//All other values are reserved.
const (
	Access3gpp    AccessType = 1
	AccessNon3gpp AccessType = 2
)

func (p *AccessType) String() string {
	switch *p {
	case Access3gpp:
		return "3gppAccess"
	case AccessNon3gpp:
		return "Non3gppAccess"
	default:
		return "invalidType"
	}
}
func (p *AccessType) Encode() byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	var octet byte

	//todo 24.501  9.11.3.11
	return octet
}
