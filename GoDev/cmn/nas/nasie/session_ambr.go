package nasie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"strconv"
	"strings"
)

type BitRate struct {
	Value uint16
	Uint  BitRateUnit
}

func (p *BitRate) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)

	unit, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
		return fmt.Errorf("fail to read byte")
	}
	p.Uint = BitRateUnit(unit)

	valueBytes := make([]byte, 2)
	binary.Read(msgBuf, binary.BigEndian, valueBytes)
	value := binary.BigEndian.Uint16(valueBytes)
	p.Value = value

	return nil
}
func (p *BitRate) Encode() []byte {
	var encBuf []byte

	encBuf = append(encBuf, byte(p.Uint))

	val := make([]byte, 2)
	binary.BigEndian.PutUint16(val, p.Value)

	encBuf = append(encBuf, val...)

	return encBuf
}

func (p *BitRate) StoreWithString(val string) error {

	elems := strings.Split(val, " ")
	if len(elems) != 2 {
		return fmt.Errorf("invalid bit rate formate (%s)", val)
	}

	value, err := strconv.ParseInt(elems[0], 0, 16)
	if err != nil {
		return fmt.Errorf("invalid bit rate value error(%s)", err)
	}
	p.Value = uint16(value)

	err = p.Uint.storeWithString(elems[1])
	if err != nil {
		return fmt.Errorf("invalid bit rate unit. error(%s)", err)
	}

	return nil
}

func (p *BitRate) String() string {
	value := p.Value
	var unitStr string

	kbps := "Kbps"
	mbps := "Mbps"
	gbps := "Gbps"
	tbps := "Tbps"
	pbps := "Pbps"

	switch p.Uint {
	case Kbps1:
		unitStr = kbps
	case Kbps4:
		value *= 4
		unitStr = kbps
	case Kbps16:
		value *= 16
		unitStr = kbps
	case Kbps64:
		value *= 64
		unitStr = kbps
	case Kbps256:
		value *= 256
		unitStr = kbps
	case Mbps1:
		unitStr = mbps
	case Mbps4:
		value *= 4
		unitStr = mbps
	case Mbps16:
		value *= 16
		unitStr = mbps
	case Mbps64:
		value *= 64
		unitStr = mbps
	case Mbps256:
		value *= 256
		unitStr = mbps
	case Gbps1:
		unitStr = gbps
	case Gbps4:
		value *= 4
		unitStr = gbps
	case Gbps16:
		value *= 16
		unitStr = gbps
	case Gbps64:
		value *= 64
		unitStr = gbps
	case Gbps256:
		value *= 256
		unitStr = gbps
	case Tbps1:
		unitStr = tbps
	case Tbps4:
		value *= 4
		unitStr = tbps
	case Tbps16:
		value *= 16
		unitStr = tbps
	case Tbps64:
		value *= 64
		unitStr = tbps
	case Tbps256:
		value *= 256
		unitStr = tbps
	case Pbps1:
		unitStr = pbps
	case Pbps4:
		value *= 4
		unitStr = tbps
	case Pbps16:
		value *= 16
		unitStr = tbps
	case Pbps64:
		value *= 64
		unitStr = tbps
	case Pbps256:
		value *= 256
		unitStr = tbps
	default:
		value = 0
		unitStr = kbps
	}

	return fmt.Sprintf("%d %s", value, unitStr)
}

type SessionAmbr struct {
	UpLink   BitRate
	DownLink BitRate
}

func (p SessionAmbr) String() string {
	return fmt.Sprintf("Uplink(%s), Downlink(%s)", p.UpLink.String(), p.DownLink.String())
}

//encode SessionAmbr to nas octet stream
func (p *SessionAmbr) Encode() []byte {
	rlogger.FuncEntry(types.ModCmn, nil)

	var encBuf []byte

	encBuf = append(encBuf, p.UpLink.Encode()...)
	encBuf = append(encBuf, p.DownLink.Encode()...)

	return encBuf
}

func (p *SessionAmbr) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	p.UpLink.Decode(msgBuf)
	p.DownLink.Decode(msgBuf)
	return nil
}
func (p *SessionAmbr) StoreWithString(uplkStr string, dwlkStr string) error {
	rlogger.FuncEntry(types.ModCmn, nil)

	err := p.UpLink.StoreWithString(uplkStr)
	if err != nil {
		return fmt.Errorf("invalid UpLink bit rate string(%s)", uplkStr)
	}
	err = p.DownLink.StoreWithString(dwlkStr)
	if err != nil {
		return fmt.Errorf("invalid DownLink bit rate string(%s)", dwlkStr)
	}
	return nil
}

//refer
//Unit for Session-AMBR for downlink (octet 3)
//
//0 0 0 0 0 0 0 0	value is not used
//0 0 0 0 0 0 0 1	value is incremented in multiples of 1 Kbps
//0 0 0 0 0 0 1 0	value is incremented in multiples of 4 Kbps
//0 0 0 0 0 0 1 1	value is incremented in multiples of 16 Kbps
//0 0 0 0 0 1 0 0	value is incremented in multiples of 64 Kbps
//0 0 0 0 0 1 0 1	value is incremented in multiples of 256 kbps
//0 0 0 0 0 1 1 0	value is incremented in multiples of 1 Mbps
//0 0 0 0 0 1 1 1	value is incremented in multiples of 4 Mbps
//0 0 0 0 1 0 0 0	value is incremented in multiples of 16 Mbps
//0 0 0 0 1 0 0 1	value is incremented in multiples of 64 Mbps
//0 0 0 0 1 0 1 0	value is incremented in multiples of 256 Mbps
//0 0 0 0 1 0 1 1	value is incremented in multiples of 1 Gbps
//0 0 0 0 1 1 0 0	value is incremented in multiples of 4 Gbps
//0 0 0 0 1 1 0 1	value is incremented in multiples of 16 Gbps
//0 0 0 0 1 1 1 0	value is incremented in multiples of 64 Gbps
//0 0 0 0 1 1 1 1	value is incremented in multiples of 256 Gbps
//0 0 0 1 0 0 0 0	value is incremented in multiples of 1 Tbps
//0 0 0 1 0 0 0 1	value is incremented in multiples of 4 Tbps
//0 0 0 1 0 0 1 0	value is incremented in multiples of 16 Tbps
//0 0 0 1 0 0 1 1	value is incremented in multiples of 64 Tbps
//0 0 0 1 0 1 0 0	value is incremented in multiples of 256 Tbps
//0 0 0 1 0 1 0 1	value is incremented in multiples of 1 Pbps
//0 0 0 1 0 1 1 0	value is incremented in multiples of 4 Pbps
//0 0 0 1 0 1 1 1	value is incremented in multiples of 16 Pbps
//0 0 0 1 1 0 0 0	value is incremented in multiples of 64 Pbps
//0 0 0 1 1 0 0 1	value is incremented in multiples of 256 Pbps
//
//Other values shall be interpreted as multiples of 256 Pbps in this version of the protocol.
type BitRateUnit byte

const (
	Kbps1   BitRateUnit = 1
	Kbps4   BitRateUnit = 2
	Kbps16  BitRateUnit = 3
	Kbps64  BitRateUnit = 4
	Kbps256 BitRateUnit = 5

	Mbps1   BitRateUnit = 6
	Mbps4   BitRateUnit = 7
	Mbps16  BitRateUnit = 8
	Mbps64  BitRateUnit = 9
	Mbps256 BitRateUnit = 10

	Gbps1   BitRateUnit = 11
	Gbps4   BitRateUnit = 12
	Gbps16  BitRateUnit = 13
	Gbps64  BitRateUnit = 14
	Gbps256 BitRateUnit = 15

	Tbps1   BitRateUnit = 16
	Tbps4   BitRateUnit = 17
	Tbps16  BitRateUnit = 18
	Tbps64  BitRateUnit = 19
	Tbps256 BitRateUnit = 20

	Pbps1   BitRateUnit = 21
	Pbps4   BitRateUnit = 22
	Pbps16  BitRateUnit = 23
	Pbps64  BitRateUnit = 24
	Pbps256 BitRateUnit = 25
)

func (p *BitRateUnit) storeWithString(val string) error {
	switch val {
	case "Kbps":
		*p = Kbps1
	case "Mbps":
		*p = Mbps1
	case "Gbps":
		*p = Gbps1
	case "Tbps":
		*p = Tbps1
	case "Pbps":
		*p = Pbps1
	default:
		return fmt.Errorf("inavlid ambr bitrate unit(%s)", val)
	}
	return nil
}

func (p *BitRate) Tokbps() uint64 {
	value := p.Value
	var kbps uint64

	kbpsMultiple := 1
	mbpsMultiple := kbpsMultiple * 1000
	gbpsMultiple := mbpsMultiple * 1000
	tbpsMultiple := gbpsMultiple * 1000
	pbpsMultiple := tbpsMultiple * 1000

	switch p.Uint {
	case Kbps1:
		kbps = uint64(value) * uint64(kbpsMultiple)
	case Kbps4:
		value *= 4
		kbps = uint64(value) * uint64(kbpsMultiple)
	case Kbps16:
		value *= 16
		kbps = uint64(value) * uint64(kbpsMultiple)
	case Kbps64:
		value *= 64
		kbps = uint64(value) * uint64(kbpsMultiple)
	case Kbps256:
		value *= 256
		kbps = uint64(value) * uint64(kbpsMultiple)

	case Mbps1:
		kbps = uint64(value) * uint64(mbpsMultiple)
	case Mbps4:
		value *= 4
		kbps = uint64(value) * uint64(mbpsMultiple)
	case Mbps16:
		value *= 16
		kbps = uint64(value) * uint64(mbpsMultiple)
	case Mbps64:
		value *= 64
		kbps = uint64(value) * uint64(mbpsMultiple)
	case Mbps256:
		value *= 256
		kbps = uint64(value) * uint64(mbpsMultiple)

	case Gbps1:
		kbps = uint64(value) * uint64(gbpsMultiple)
	case Gbps4:
		value *= 4
		kbps = uint64(value) * uint64(gbpsMultiple)
	case Gbps16:
		value *= 16
		kbps = uint64(value) * uint64(gbpsMultiple)
	case Gbps64:
		value *= 64
		kbps = uint64(value) * uint64(gbpsMultiple)
	case Gbps256:
		value *= 256
		kbps = uint64(value) * uint64(gbpsMultiple)

	case Tbps1:
		kbps = uint64(value) * uint64(tbpsMultiple)
	case Tbps4:
		value *= 4
		kbps = uint64(value) * uint64(tbpsMultiple)
	case Tbps16:
		value *= 16
		kbps = uint64(value) * uint64(tbpsMultiple)
	case Tbps64:
		value *= 64
		kbps = uint64(value) * uint64(tbpsMultiple)
	case Tbps256:
		value *= 256
		kbps = uint64(value) * uint64(tbpsMultiple)

	case Pbps1:
		kbps = uint64(value) * uint64(pbpsMultiple)
	case Pbps4:
		value *= 4
		kbps = uint64(value) * uint64(pbpsMultiple)
	case Pbps16:
		value *= 16
		kbps = uint64(value) * uint64(pbpsMultiple)
	case Pbps64:
		value *= 64
		kbps = uint64(value) * uint64(pbpsMultiple)
	case Pbps256:
		value *= 256
		kbps = uint64(value) * uint64(pbpsMultiple)
	default:
		value = 0
		kbps = uint64(value) * uint64(pbpsMultiple)
	}

	return kbps
}
