package nasmsg

import (
	"bytes"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 8.3.11.1 f40 2019-06
type PduSessionCommandRejectMsg struct {
	EPD      nas.Epd
	SmHeader nas.SmNasMessageHeader
	SmCause  nas.Sm5gCause
	//Extended protocol configuration options
}

func (p *PduSessionCommandRejectMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	// mandatory IEs
	// the header have already decoded
	cause, _ := msgBuf.ReadByte()
	p.SmCause = nas.Sm5gCause(cause)
	return nil
}
