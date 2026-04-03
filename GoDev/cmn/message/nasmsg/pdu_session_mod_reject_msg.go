package nasmsg

import (
	"bytes"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

//24501 8.3.8 f40 2019-06
type PduSessionModifyRejectMsg struct {
	//mandatory
	EPD       nas.Epd
	MsgHeader nas.SmNasMessageHeader
	SMCause   nas.Sm5gCause
	//optional
	BackOffTimer            nasie.GprsTimer3
	SMCongestionReAttempInd bool
	//Extended protocol configuration options
	// Ie flags
	IeFlags bitset.BitSet
}

const (
	Ieid_PduSessionModReject_BackOffTimer uint = iota
	Ieid_PduSessionModReject_SMCongestionReAttempInd
)

func (p *PduSessionModifyRejectMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	//Mandatory
	//MsgHeader V
	msgHeaderValue, _ := p.MsgHeader.Encode()
	encBuf = append(encBuf, msgHeaderValue[:]...)
	encBuf = append(encBuf, byte(p.SMCause))
	// Optional IEs
	//NonMandatory IE
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case Ieid_PduSessionModReject_BackOffTimer:
			// T
			encBuf = append(encBuf, byte(nasie.IeiBackOffTimerValue))
			timerOctet := p.BackOffTimer.Encode()
			// L
			encBuf = append(encBuf, byte(len(timerOctet)))
			// V
			encBuf = append(encBuf, timerOctet[:]...)
		case Ieid_PduSessionModReject_SMCongestionReAttempInd:
			// T
			encBuf = append(encBuf, byte(nasie.IeiSMCongestionReAttempInd))
			// L
			encBuf = append(encBuf, 1)
			// V
			encBuf = append(encBuf, utils.BoolToByte(p.SMCongestionReAttempInd))
		}
	}
	return encBuf, nil
}

func (p *PduSessionModifyRejectMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	// mandatory IEs
	// the header have already decoded
	cause, _ := msgBuf.ReadByte()
	p.SMCause = nas.Sm5gCause(cause)
	//fmt.Println(p.SMCause)
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
			msgBuf.ReadByte()
			err := p.BackOffTimer.Decode(msgBuf)
			if err != nil {
				return fmt.Errorf("failed to decode len of BackOffTimer")
			}
			p.IeFlags.Set(Ieid_PduSessionModReject_BackOffTimer)
			//fmt.Println(p.BackOffTimer)
		case nasie.IeiSMCongestionReAttempInd:
			msgBuf.ReadByte()
			value, _ := msgBuf.ReadByte()
			p.SMCongestionReAttempInd, _ = utils.GetBitValue(value, 1)
			p.IeFlags.Set(Ieid_PduSessionModReject_SMCongestionReAttempInd)
			//fmt.Println(p.SMCongestionReAttempInd)
		}
	}
	return nil
}
