package nasmsg

import (
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 Table 8.2.18.1.1
type ServiceRejectMsg struct {
	//Mandatory IE
	//Extended protocol discriminator	Extended protocol discriminator 9.2
	//Security header type	Security header type 	9.3
	//Spare half octet	Spare half octet	9.5
	//Service reject message identity	Message type	9.7

	//5GMM cause	5GMM cause	9.11.3.2
	FiveGSRCause nas.Mm5gCause

	//Optional IE
	//PDU session status	PDU session status	9.11.3.44
	PDUSessStatus nasie.SessionStatus
	//T3346 value	GPRS timer 2	9.11.2.4
	T3346 nasie.GprsTimer2
	//EAP message	EAP message	9.11.2.2
	EAPMsg []byte
	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	IeidServicerejtPdusessstatus uint = iota
	IeidServicerejtT3346
	IeidServicerejtEapmsg
)

func (p ServiceRejectMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	// mandatory IEs
	// header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.ServiceReject
	encBuf = append(encBuf, header.Encode()...)

	// 5GMM cause
	encBuf = append(encBuf, byte(p.FiveGSRCause))

	//Optional IEs
	optIeOctet := p.EncodeOptIes()
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p ServiceRejectMsg) EncodeOptIes() []byte {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	// for optional IEs
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case IeidServicerejtPdusessstatus:
			typeOctet := byte(nasie.IeiPDUSessionStatus)
			PDUSessStatusByte := p.PDUSessStatus.Encode()
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			encBuf = append(encBuf, byte(len(PDUSessStatusByte)))
			// V
			encBuf = append(encBuf, PDUSessStatusByte...)
		case IeidServicerejtT3346:
			typeOctet := byte(nasie.IeiT3346Value)
			T3346Byte := p.T3346.Encode()
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			encBuf = append(encBuf, byte(len(T3346Byte)))
			// V
			encBuf = append(encBuf, T3346Byte...)
		case IeidServicerejtEapmsg:
			// T
			encBuf = append(encBuf, byte(nasie.IeiEAPMessage))
			// L
			encBuf = append(encBuf, byte(len(p.EAPMsg)))
			// V
			encBuf = append(encBuf, p.EAPMsg...)
		}
	}
	return encBuf
}
