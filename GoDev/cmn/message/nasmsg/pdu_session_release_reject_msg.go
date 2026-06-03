package nasmsg

import (
	"bytes"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 8.3.13
type PduSessionReleaseRejectMsg struct {
	//Mandatory
	MsgHeader nas.SmNasMessageHeader
	SMCause   nas.Sm5gCause
	//optional
	//ExtendProtocolConfigOpt []byte
}

//encode a session release reject msg from nas octet stream
func (p *PduSessionReleaseRejectMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	//MsgHeader V
	msgHeaderValue, _ := p.MsgHeader.Encode()
	encBuf = append(encBuf, msgHeaderValue[:]...)
	//SMCause V
	encBuf = append(encBuf, byte(p.SMCause))
	return encBuf, nil
}

func (p *PduSessionReleaseRejectMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	// mandatory IEs
	// the header have already decoded
	cause, _ := msgBuf.ReadByte()
	p.SMCause = nas.Sm5gCause(cause)
	return nil
}
