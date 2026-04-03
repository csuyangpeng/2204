package nasmsg

import (
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

//24501 f10 8.2.14  De-registration request (UE terminated de-registration)
type DeRegistrationRequestUeTMsg struct {
	//Extended protocol discriminator
	//Security header type
	//Spare half octet
	//De-registration request message identity
	DeRegistrationType nasie.DeRegistrationTypeIE
	//Spare half octet

	//optional
	Cause nas.Mm5gCause
	T3346 nasie.GprsTimer2

	//Indicates whether an IE is assigned or it is an empty value
	IeFlags bitset.BitSet
}

const (
	IeidDregUetCause uint = iota
	IeidDregUetT3546
)

func (p *DeRegistrationRequestUeTMsg) Reset() {
	p.DeRegistrationType.Reset()
}

//encode a deRegistration request msg from nas octet stream
func (p *DeRegistrationRequestUeTMsg) Encode() ([]byte, error) {

	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	//header
	var header nas.MmNasMessageHeader
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.DeregistrationRequestUeT
	encBuf = append(encBuf, byte(header.ExtendProtoDisc))
	encBuf = append(encBuf, byte(header.SecurityHeaderType)<<4)
	encBuf = append(encBuf, byte(header.MessageType))

	// De-registration type
	deRegistNgKSIOctet := byte(utils.BoolToByte(p.DeRegistrationType.SwithOff)) << 3
	deRegistNgKSIOctet |= byte(utils.BoolToByte(p.DeRegistrationType.IsReRegistrationRequired)) << 2
	deRegistNgKSIOctet |= byte(p.DeRegistrationType.AccessType)

	encBuf = append(encBuf, deRegistNgKSIOctet)
	// Optional IEs
	optIeOctet, err := p.EncodeOptIes()
	if err != nil {
		err = fmt.Errorf("failed to decode optional ies")
	}

	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}
func (p *DeRegistrationRequestUeTMsg) EncodeOptIes() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	var encBuf []byte
	//Cause, it's format is TV
	if p.IeFlags.Test(IeidDregUetCause) == true {
		// T
		encBuf = append(encBuf, byte(nasie.Iei5GMMCause))
		// V
		encBuf = append(encBuf, byte(p.Cause))
	}
	// T3346 TLV
	if p.IeFlags.Test(IeidDregUetT3546) == true {
		// T
		encBuf = append(encBuf, byte(nasie.IeiT3346Value))
		// L
		encBuf = append(encBuf, byte(len(p.T3346.Encode())))
		// V
		encBuf = append(encBuf, p.T3346.Encode()[:]...)
	}
	return encBuf, nil
}
