package nasmsg

import (
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"

	"github.com/willf/bitset"
)

//24501   8.2.5
type AuthenticationRejectMsg struct {
	// optional
	// EAP message 9.11.2.2 O TLV-E 7-1503

	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	Ieid_AuthReject_EAPMessage uint = iota
)

func (p *AuthenticationRejectMsg) Reset() {
	p.IeFlags = bitset.BitSet{}
}

// Print Registration Request Msg
func (p *AuthenticationRejectMsg) String() string {
	var msgStr string
	msgStr = "Authentication Request Message ( "
	msgStr += fmt.Sprintf("IeFlag(%v) ", p.IeFlags)
	return msgStr
}
func (p *AuthenticationRejectMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	// mandatory IEs
	// header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.AuthenticationReject
	encBuf = append(encBuf, header.Encode()...)

	return encBuf, nil
}
