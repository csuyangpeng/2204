package nasmsg

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501   8.2.26.1
type SecurityModeRejectMsg struct {
	//Extended protocol discriminator
	//Security header type
	//Spare half octet
	//Security mode reject message identity
	Cause nas.Mm5gCause
}

func (p SecurityModeRejectMsg) String() string {
	return fmt.Sprintf("security mod REJECT, imsisv(%x)", p.Cause)
}

func (p *SecurityModeRejectMsg) Reset() {
	p.Cause = 1
}

func (p *SecurityModeRejectMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	ieType, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
		return nil
	}
	p.Cause = nas.Mm5gCause(ieType)
	return nil
}
