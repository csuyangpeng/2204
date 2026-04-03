package nasie

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

type T3UintType byte

const (
	TenMin          T3UintType = 0
	OneHour         T3UintType = 1
	TenHour         T3UintType = 2
	TwoSecs         T3UintType = 3
	ThirtySecs      T3UintType = 4
	OneMin          T3UintType = 5
	ThreeTwentyHour T3UintType = 6
)

type GprsTimer3 struct {
	Uint       T3UintType
	TimerValue uint8
}

// refer to TS24.008  10.5.7.4a
func (p *GprsTimer3) Encode() []byte {
	var octet []byte
	//fmt.Println("p",p)
	gprs3Octet := byte(p.Uint) << 5
	gprs3Octet |= byte(p.TimerValue)

	octet = append(octet, gprs3Octet)
	//fmt.Println("gprs3Octet",gprs3Octet)
	return octet
}

func (p *GprsTimer3) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	gprsTimer3Value, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read GprsTimer3")
	}
	//fmt.Println("gprsTimer3Value",gprsTimer3Value)
	unit, _ := utils.GetBitsValue(gprsTimer3Value, 6, 8)
	p.Uint = T3UintType(unit >> 5)
	value, _ := utils.GetBitsValue(gprsTimer3Value, 1, 5)
	p.TimerValue = value
	//fmt.Println("p",p)
	return nil
}
