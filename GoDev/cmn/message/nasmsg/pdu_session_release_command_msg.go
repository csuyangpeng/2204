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
type PduSessionReleaseCommandMsg struct {
	//Mandatory
	EPD       nas.Epd
	MsgHeader nas.SmNasMessageHeader
	SMCause   nas.Sm5gCause

	//optional
	BackOffTimer                        nasie.GprsTimer3
	FiveGSMCongestionReAttemptIndicator byte
	//EAPMessage todo
	//ExtendProtocolConfigOpt []byte

	// Ie flags
	IeFlags bitset.BitSet
}

//encode a session release command msg from nas octet stream
func (p *PduSessionReleaseCommandMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte

	encBuf = append(encBuf, byte(p.EPD))
	//MsgHeader V
	msgHeaderValue, _ := p.MsgHeader.Encode()
	encBuf = append(encBuf, msgHeaderValue[:]...)

	//SMCause V
	// V
	encBuf = append(encBuf, byte(p.SMCause))

	// Optional IEs
	//NonMandatory IE
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case nas.Ieid_BackOffTimer:
			// T
			encBuf = append(encBuf, byte(nasie.IeiBackOffTimerValue))
			valueBuf := p.BackOffTimer.Encode()
			//L
			encBuf = append(encBuf, byte(len(valueBuf)))
			// V
			encBuf = append(encBuf, valueBuf[:]...)

			if p.SMCause == nas.InsufficientResourcesForSpecificSliceAndDNN ||
				p.SMCause == nas.InsufficientResourcesForSpecificSlice {
				// T
				encBuf = append(encBuf, byte(nasie.IeiFIVEGSMCongestionReAttemptInt))
				// L
				encBuf = append(encBuf, 1)
				// V
				encBuf = append(encBuf, p.FiveGSMCongestionReAttemptIndicator)
			}
		}
	}

	return encBuf, nil
}

// decode a session release command msg from nas octet stream
func (p *PduSessionReleaseCommandMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	// mandatory IEs
	epd, _ := msgBuf.ReadByte()
	if epd != byte(nas.Epd5gsSessMgntMsg) {
		return fmt.Errorf("epd error")
	}
	p.EPD = nas.Epd(epd)
	// the header have already decoded
	err := p.MsgHeader.Decode(msgBuf)
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to decode header, err:", err)
		return fmt.Errorf("failed to decode header")
	}
	cause, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read sm nas msg header psi, "+
			"error:%s", err)
	}
	p.SMCause = nas.Sm5gCause(cause)
	//optional
	//optional
	for {
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		// 第一个字节就是IE的标志，直接识别即可
		switch nasie.Iei(ieType) {
		case nasie.IeiBackOffTimerValue:
			//todo
		}
	}
	return nil
}
