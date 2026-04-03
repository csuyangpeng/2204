package nasmsg

import (
	"bytes"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 8.3.10.1 f40 2019-06
type PduSessionModifyCompleteMsg struct {
	EPD      nas.Epd
	SmHeader nas.SmNasMessageHeader
	//Extended protocol configuration options
}

func (p *PduSessionModifyCompleteMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "nothing to decode")
	return nil
}
