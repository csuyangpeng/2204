package nasie

import (
	"bytes"
	"encoding/binary"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type SessionReactCause struct {
	PduSessionId uint8
	Cause        nas.Mm5gCause
}

//24501 9.11.3.43
type SessionReactErrCause struct {
	SessReactErr []SessionReactCause
}

func (p SessionReactErrCause) Encode() []byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	if len(p.SessReactErr) > 1 {
		for i := 0; i < len(p.SessReactErr); i++ {
			encBuf = append(encBuf, p.SessReactErr[i].PduSessionId)
			encBuf = append(encBuf, byte(p.SessReactErr[i].Cause))
		}
	}
	return encBuf
}

func (p *SessionReactErrCause) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	//L
	lenBytes := make([]byte, 2)
	err := binary.Read(msgBuf, binary.BigEndian, &lenBytes)
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "Decode err len failed")
		return nas.ErrDecodeNasIeFailed
	}
	errLen := binary.BigEndian.Uint16(lenBytes)

	if errLen > 1 {
		for i := 0; i < len(p.SessReactErr); i++ {
			p.SessReactErr[i].PduSessionId, _ = msgBuf.ReadByte()
			cause, _ := msgBuf.ReadByte()
			p.SessReactErr[i].Cause = nas.Mm5gCause(cause)
		}
	}

	return nil
}
