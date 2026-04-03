package nasmsg

import (
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"

	"github.com/willf/bitset"
)

//24501   8.2.25.1
type SecurityModeCommandMsg struct {
	// mandatory
	// Selected NAS security algorithms	9.11.3.32	M	V  1
	SelectNasSecAlg types3gpp.SelNasSecAlgo

	// NAS key set identifier	        9.11.3.32	M	V  1/2
	NgKSI nasie.NasKSI

	// Replayed UE security capabilities 9.11.3.54   M   LV 3-9
	UeSecCap types3gpp.SecurityCapability

	// optional
	// IMEISV request 9.11.3.28 O TV 1
	ImeiSvReq bool

	// Selected EPS NAS security algorithms 9.11.3.25 O TV 2
	//SelEpsNasSecAlg

	// Additional 5G security information  9.11.3.12 O TLV 3
	Add5GSecInfo nasie.Add5GSecInfo

	// ABBA  ABBA  9.11.3.10 O TLV 4-n
	Abba []byte

	// Replayed S1 UE security capabilities 9.11.3.48A O TLV 4-7
	//S1UeSecCap

	// EAP message 9.11.2.2 O TLV-E 7-1503
	// not supported currently

	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	IeidSecmodcmdImeireq uint = iota
	IeidSecmodcmdSelepsnassecalgo
	IeidSecmodcmdAdd5gsecinfo
	IeidSecmodcmdEapmsg
	IeidSecmodcmdAbba
	IeidSecmodcmdRes1ueseccap
)

func (p *SecurityModeCommandMsg) Reset() {
	p.SelectNasSecAlg.Reset()
	p.NgKSI.Reset()
	p.UeSecCap.Reset()
	p.ImeiSvReq = false
	p.Add5GSecInfo.Reset()
	p.Abba = []byte{}
	p.IeFlags = bitset.BitSet{}
}

// Print Registration Request Msg
func (p *SecurityModeCommandMsg) String() string {
	var msgStr string
	msgStr = "SecurityModeCommandMsg ( "
	msgStr += fmt.Sprintf("SelNasSecAlgo(%s) ", p.SelectNasSecAlg)
	msgStr += fmt.Sprintf("KSI(%s) ", p.NgKSI)
	msgStr += fmt.Sprintf("UeSecCapbility(%s) ", p.UeSecCap)
	msgStr += fmt.Sprintf("ImeiSvReq(%v) ", p.ImeiSvReq)
	msgStr += fmt.Sprintf("Add5GSecInfo(%s) ", p.Add5GSecInfo)
	msgStr += fmt.Sprintf("Abba(%x) ", p.Abba)
	msgStr += fmt.Sprintf("IeFlag(%v) ", p.IeFlags)
	return msgStr
}

func (p *SecurityModeCommandMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	// mandatory IEs
	// header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.SecurityModeCommand
	encBuf = append(encBuf, header.Encode()...)

	// select nas security alogrithms
	encBuf = append(encBuf, p.SelectNasSecAlg.Encode())

	// KSI
	//ngKSI && Spare half octet  V
	KSI := byte(utils.BoolToByte(p.NgKSI.Tsc)) << 3
	KSI |= p.NgKSI.Ksi
	encBuf = append(encBuf, KSI)

	// Replayed ue security capabilities  LV
	encBuf = append(encBuf, p.UeSecCap.Encode()...)

	//Optional IEs
	optIeOctet := p.EncodeOptIes()
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p *SecurityModeCommandMsg) EncodeOptIes() []byte {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	// for optional IEs
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case IeidSecmodcmdImeireq:
			//TV 1 byte
			//IMEISV request value (octet 1)
			//Bits
			//3	2	1
			//0	0	0		IMEISV not requested
			//0	0	1		IMEISV requested
			//
			//All other values are interpreted as IMEISV not requested by this version of the protocol.
			// T
			typeOctet := byte(nasie.IeiImsiSv)
			// V
			if p.ImeiSvReq {
				typeOctet = typeOctet | 0x01
			}
			encBuf = append(encBuf, typeOctet)
		case IeidSecmodcmdSelepsnassecalgo:
		case IeidSecmodcmdAdd5gsecinfo:
			encBuf = append(encBuf, p.Add5GSecInfo.Encode()...)
		case IeidSecmodcmdEapmsg:
		case IeidSecmodcmdAbba:
			// ABBA TLV 4-n 9.11.3.10
			//T
			typeOctet := byte(nasie.IeiABBA)
			encBuf = append(encBuf, typeOctet)
			//L
			abbalen := byte(len(p.Abba))
			if abbalen > 0 {
				encBuf = append(encBuf, byte(abbalen))
				encBuf = append(encBuf, p.Abba...)
			} else {
				return nil
			}
		case IeidSecmodcmdRes1ueseccap:
		}
	}

	return encBuf
}

func (p *SecurityModeCommandMsg) Decode() {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
}
