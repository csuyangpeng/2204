package nasmsg

import (
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

// 24.501  8.3.3
type PduSessionEstbRejectMsg struct {
	//Mandatory
	MsgHeader nas.SmNasMessageHeader
	SMCause   nas.Sm5gCause

	//Optional
	BackOffTimerValue nasie.GprsTimer3  // procedure retry
	AllowedSSCMode    AllowedSSCModeSet // when cause #68 "not supported SSC mode
	//EAPMessage
	FiveGSMCongestionReAttemptIndicator byte
	//Extended protocol configuration options

	//Indicates whether an IE is assigned or it is an empty value
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	Ieid_PduSessionEstbReject_BackOffTimerValue uint = iota
	Ieid_PduSessionEstbReject_AllowedSSCMode
	Ieid_PduSessionEstbReject_EAPMessage
	Ieid_PduSessionEstbReject_EPCO
)

// Print Session Reject Msg
func (srMsg PduSessionEstbRejectMsg) Print() {
	fmt.Println("-------------------------")
	fmt.Println("Session Reject Msg info:")
	fmt.Println("	MsgHeader: ", srMsg.MsgHeader)
	fmt.Println("	SMCause: ", srMsg.SMCause)
	fmt.Println("	BackOffTimerValue: ", srMsg.BackOffTimerValue)
	fmt.Println("	AllowedSSCMode: ", srMsg.AllowedSSCMode)
	fmt.Println("-------------------------")
}

type AllowedSSCModeSet struct {
	Ssc1 bool
	Ssc2 bool
	Ssc3 bool
}

//encode sessionRejectMsg from nas octet stream
func (p *PduSessionEstbRejectMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	// mandatory IEs
	//epd
	encBuf = append(encBuf, byte(nas.Epd5gsSessMgntMsg))
	// MsgHeader,V
	msgHeaderOctet, _ := p.MsgHeader.Encode()
	encBuf = append(encBuf, msgHeaderOctet[:]...)

	// SMCause,V
	encBuf = append(encBuf, byte(p.SMCause))

	// Optional IEs
	optIeOctet, err := p.EncodeOptIes()
	if err != nil {
		err = fmt.Errorf("failed to decode optional ies")
	}
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p *PduSessionEstbRejectMsg) EncodeOptIes() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte

	//Optional
	//for other optional IEs:
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextClear(i + 1) {
		switch i {
		case Ieid_PduSessionEstbReject_BackOffTimerValue:
			// T
			encBuf = append(encBuf, byte(nasie.IeiBackOffTimerValue))
			timerOctet := p.BackOffTimerValue.Encode()
			// L
			encBuf = append(encBuf, byte(len(timerOctet)))
			// V
			encBuf = append(encBuf, timerOctet[:]...)
		case Ieid_PduSessionEstbReject_AllowedSSCMode:
			sscOctet := byte(nasie.IeiAllowedSSCMode)
			sscOctet |= utils.BoolToByte(p.AllowedSSCMode.Ssc3) << 2
			sscOctet |= utils.BoolToByte(p.AllowedSSCMode.Ssc2) << 1
			sscOctet |= utils.BoolToByte(p.AllowedSSCMode.Ssc1)
			encBuf = append(encBuf, sscOctet)
		}
	}
	return encBuf, nil
}
