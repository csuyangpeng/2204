package nasmsg

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"

	"github.com/willf/bitset"
)

//24501   8.2.2
type AuthenticationResponseMsg struct {
	//Optional
	//Auth Resp parameter 	Res* 24.301 9.9.3.4 TLV
	ResStar [types.XressSize]byte

	// EAP message 9.11.2.2 O TLV-E 7-1503
	// not supported currently

	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	IeidAuthrespResstar uint = iota
	Ieid_AuthResp_Eap
)

func (p *AuthenticationResponseMsg) Reset() {
	p.IeFlags = bitset.BitSet{}
}

// Print Authentication Response
func (p *AuthenticationResponseMsg) String() string {
	var msgStr string
	msgStr = "Authentication Response Message ( "
	msgStr += fmt.Sprintf("RES*(%x) ", p.ResStar)
	msgStr += fmt.Sprintf("IeFlag(%v) ", p.IeFlags)
	return msgStr
}

func (p *AuthenticationResponseMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	// Optional IEs
	err := p.DecodeOptIes(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode optional ies")
	}

	return nil
}

func (p *AuthenticationResponseMsg) DecodeOptIes(msgBuf *bytes.Reader) error {
	for {
		// IE的标识被编码进了第一个字节，所以要单独拎出来
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}

		switch nasie.Iei(ieType) {
		case nasie.IeiAuthResStar:
			length, _ := msgBuf.ReadByte()
			if length != types.XressSize {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "failed to decode Res* from "+
					"AuthenticationResponse Msg .length(%d) should be 16 in 5G AKA", length)
				continue
			}
			for i := 0; i < int(types.XressSize); i++ {
				val, _ := msgBuf.ReadByte()
				p.ResStar[i] = val
			}
			p.IeFlags.Set(IeidAuthrespResstar)
		default:
		}
	}

	return nil
}
