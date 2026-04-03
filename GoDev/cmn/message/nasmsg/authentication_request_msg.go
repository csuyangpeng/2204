package nasmsg

import (
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"

	"github.com/willf/bitset"
)

//24501   8.2.1.1
type AuthenticationRequestMsg struct {
	//mandatory
	//ngKSI 	NAS key set identifier	9.11.3.32	M	V  1/2
	NgKSI nasie.NasKSI

	//ABBA  ABBA  9.11.3.10 M LV 3-n
	Abba []byte

	// Authentication parameter RAND  9.11.3.16 O TV 17
	Rand [types.RandSize]byte

	// Authentication parameter AUTN  9.11.3.15 O TLV 18
	Autn [types.RandSize]byte

	// EAP message 9.11.2.2 O TLV-E 7-1503
	// not supported currently

	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	IeidAuthreqRand uint = iota
	IeidAuthreqAutn
	IeidAuthreqEap
)

func (p *AuthenticationRequestMsg) Reset() {
	p.NgKSI.Reset()
	p.Abba = []byte{}
	p.NgKSI.Reset()
	p.IeFlags = bitset.BitSet{}
}

// Print Registration Request Msg
func (p *AuthenticationRequestMsg) String() string {
	var msgStr string
	msgStr = "Authentication Request Message ( "
	msgStr += fmt.Sprintf("KSI(%s) ", p.NgKSI)
	msgStr += fmt.Sprintf("Abba(%x) ", p.Abba)
	msgStr += fmt.Sprintf("Rand(%x) ", p.Rand)
	msgStr += fmt.Sprintf("Autn(%x) ", p.Autn)
	msgStr += fmt.Sprintf("IeFlag(%v) ", p.IeFlags)
	return msgStr
}

func (p *AuthenticationRequestMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	// mandatory IEs
	// header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.AuthenticationRequest
	encBuf = append(encBuf, header.Encode()...)

	// KSI
	//ngKSI && Spare half octet
	KSI := byte(utils.BoolToByte(p.NgKSI.Tsc)) << 3
	KSI |= p.NgKSI.Ksi
	encBuf = append(encBuf, KSI)

	// ABBA
	abbalen := byte(len(p.Abba))
	if abbalen > 0 {
		encBuf = append(encBuf, byte(abbalen))
		encBuf = append(encBuf, p.Abba...)
	} else {
		return nil, fmt.Errorf("failed to encode Authentication Reqeust Message, invalid ABBA, len=%d", abbalen)
	}

	//Optional IEs
	optIeOctet := p.EncodeOptIes()
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p *AuthenticationRequestMsg) EncodeOptIes() []byte {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	// for optional IEs
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case IeidAuthreqRand:
			typeOctet := byte(nasie.IeiAuthRand)
			// T
			encBuf = append(encBuf, typeOctet)
			// V
			encBuf = append(encBuf, p.Rand[:]...)
		case IeidAuthreqAutn:
			typeOctet := byte(nasie.IeiAuthAutn)
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			encBuf = append(encBuf, byte(len(p.Autn)))
			// V
			encBuf = append(encBuf, p.Autn[:]...)
		case IeidAuthreqEap:
		}
	}

	return encBuf
}

func (p *AuthenticationRequestMsg) Decode() {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
}
