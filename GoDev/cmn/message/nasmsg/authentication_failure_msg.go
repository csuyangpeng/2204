package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 8.2.4
type AuthenticationFailureMsg struct {
	//mandatory
	Cause nas.Mm5gCause

	IsAutsPrst bool
	Auts       [types.AutsSize]byte
}

func (p *AuthenticationFailureMsg) Reset() {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	p.Cause = 0
	p.IsAutsPrst = false
	for i := 0; i < int(types.AutsSize); i++ {
		p.Auts[i] = 0
	}
}

const (
	IeiAuthFailurePara uint = iota
)

// Print Registration Request Msg
func (p *AuthenticationFailureMsg) String() string {
	var msgStr string
	msgStr = "Authentication Failure Message ( "
	msgStr += fmt.Sprintf("Cause(%d) ", p.Cause)
	msgStr += fmt.Sprintf("IsAutsPrst(%v) ", p.IsAutsPrst)
	msgStr += fmt.Sprintf("Auts(%x) ", p.Auts)
	return msgStr
}

func (p *AuthenticationFailureMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	cause, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "failed to decode 5GMM Cause, error(%s)", err)
		return fmt.Errorf("failed to decode 5GMM Cause, error(%s)", err)
	}

	p.Cause = nas.Mm5gCause(cause)

	ieType, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more ies")
		return nil
	}

	if ieType == byte(IeiAuthFailurePara) {
		length, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "failed to decode the length of ie, error(%s)", err)
			return fmt.Errorf("failed to decode the length of ie, error(%s)", err)
		}

		if length != types.AutsSize {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "length of auts is 14, but get length(%d)", length)
			return fmt.Errorf("length of auts is 14, but get length(%d)", length)
		}

		auts := make([]byte, types.AutsSize)
		err = binary.Read(msgBuf, binary.BigEndian, &auts)
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "fail to decode the auts, error(%s)", err)
			return fmt.Errorf("fail to decode the auts, error(%s)", err)
		}

		for i := 0; i < int(types.AutsSize); i++ {
			p.Auts[i] = auts[i]
		}

		p.IsAutsPrst = true
	}

	return nil
}
