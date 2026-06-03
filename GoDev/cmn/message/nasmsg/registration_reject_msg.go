package nasmsg

import (
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// 5GMM cause              9.11.3.2
// T3346   GPRS timer 2    9.11.2.4
// T3502   GPRS timer 2    9.11.2.4
// EAP message	           9.11.2.2

type RegistrationRejectMsg struct {
	MMCause nas.Mm5gCause
	T3346   nasie.GprsTimer2
	T3502   nasie.GprsTimer2
	//EapMessage

	//Indicates whether an IE is assigned or it is an empty value
	IeFlags bitset.BitSet
}

const (
	IeidRegRejectT3346 uint = iota
	IeidRegRejectT3502
)

//encode a registration accept msg from nas octet stream
func (p *RegistrationRejectMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.RegistrationReject
	encBuf = header.Encode()

	// mandatory IEs
	encBuf = append(encBuf, byte(p.MMCause))

	// Optional IEs
	optIeOctet, err := p.EncodeOptIes()
	if err != nil {
		err = fmt.Errorf("failed to decode optional ies")
	}

	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}
func (p *RegistrationRejectMsg) EncodeOptIes() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	var encBuf []byte

	if p.IeFlags.Test(IeidRegRejectT3346) == true {
		// T
		encBuf = append(encBuf, byte(nasie.IeiT3346Value))
		// L
		encBuf = append(encBuf, byte(len(p.T3346.Encode())))
		// V
		encBuf = append(encBuf, p.T3346.Encode()[:]...)
	}

	if p.IeFlags.Test(IeidRegRejectT3502) == true {
		// T
		encBuf = append(encBuf, byte(nasie.IeiT3502Value))
		// L
		encBuf = append(encBuf, byte(len(p.T3502.Encode())))
		// V
		encBuf = append(encBuf, p.T3502.Encode()[:]...)
	}
	return encBuf, nil
}
