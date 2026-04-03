package nasmsg

import (
	"bytes"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 8.3.14
type PduSessionReleaseCompleteMsg struct {
	//Mandatory
	MsgHeader nas.SmNasMessageHeader
	//optional
	SMCause nas.Sm5gCause
	//ExtendProtocolConfigOpt []byte
	//Ie flags
	IeFlags bitset.BitSet
}

//encode a session release complete msg from nas octet stream
func (p *PduSessionReleaseCompleteMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	//Mandatory
	//MsgHeader V
	msgHeaderValue, _ := p.MsgHeader.Encode()
	encBuf = append(encBuf, msgHeaderValue[:]...)
	// Optional IEs
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case nas.Ieid_SMCause:
			// T
			encBuf = append(encBuf, byte(nasie.IeiSMCause))
			// V
			encBuf = append(encBuf, byte(p.SMCause))
		}
	}
	return encBuf, nil
}

// decode a session release complete msg from nas octet stream
func (p *PduSessionReleaseCompleteMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	// mandatory IEs
	// the header have already decoded
	//optional
	for {
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			return fmt.Errorf("failed to read bytes when decode optional ies")
		}
		// 第一个字节就是IE的标志，直接识别即可
		switch nasie.Iei(ieType) {
		case nasie.IeiSMCause:
			cause, _ := msgBuf.ReadByte()
			p.SMCause = nas.Sm5gCause(cause)
		}
	}
	return nil
}
