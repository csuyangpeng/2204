package nasie

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

type FiveGUpdateType struct {
	NGRANRCU   bool
	SMSRequest bool
}

func (p *FiveGUpdateType) Encode() byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	var buf byte
	buf = utils.BoolToByte(p.NGRANRCU) << 1
	buf |= utils.BoolToByte(p.SMSRequest)
	return buf
}

func (p *FiveGUpdateType) Reset() {
	p.NGRANRCU = false
	p.SMSRequest = false
}
