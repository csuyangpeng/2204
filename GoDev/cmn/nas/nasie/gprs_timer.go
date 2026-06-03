package nasie

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

type GprsTimer struct {
	Uint       UintType
	TimerValue uint8
}

type UintType byte

const (
	TwoSeconds  UintType = 0
	OneMinute   UintType = 1
	Decihours   UintType = 2
	Deactivated UintType = 7
)

// refer to TS24.008  10.5.7.4a
func (p *GprsTimer) Encode() []byte {
	var octet []byte
	gprs2Octet := byte(p.Uint) << 5
	gprs2Octet |= byte(p.TimerValue)
	octet = append(octet, gprs2Octet)
	return octet
}

func (p *GprsTimer) Decode(msgBuf *bytes.Reader) error {
	octet, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "gprsTimer decode wrong")
		return fmt.Errorf("gprsTimer decode wrong")
	}
	unit, _ := utils.GetBitsValue(octet, 6, 8)
	p.Uint = UintType(unit)
	value, _ := utils.GetBitsValue(octet, 1, 5)
	p.TimerValue = uint8(value)
	return nil
}
