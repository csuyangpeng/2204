package nasie

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

type PacketFilterLists struct {
	PFList []PacketFilterList
}

func (p PacketFilterLists) String() string {
	var rt string
	for i, v := range p.PFList {
		rt = rt + fmt.Sprintf("QosRule{%d : %s}", i+1, v)
	}

	return rt
}

type PacketFilterList struct {
	PktFilterDirection   PacketFilterDirection
	PktFilterIdentifier  byte // 0~15
	PacketFilterContents PacketFilterContent
}

func (p PacketFilterList) String() string {
	return fmt.Sprintf("Filter Direction(%s), Filter id(%d), Filter Contents(%s),",
		p.PktFilterDirection.String(), p.PktFilterIdentifier, p.PacketFilterContents)
}

type PacketFilterDirection byte

func (p *PacketFilterDirection) StoreWithString(val string) error {
	switch val {
	case "uplink":
		*p = UplinkOnly
	case "downlink":
		*p = DownlinkOnly
	case "bidirectional":
		*p = Bidirectional
	default:
		return fmt.Errorf("invalid packet filter direction(%s)", val)
	}
	return nil
}
func (p PacketFilterDirection) String() string {
	switch p {
	case Reserved:
		return fmt.Sprintf("Reserved")
	case DownlinkOnly:
		return fmt.Sprintf("DownlinkOnly")
	case UplinkOnly:
		return fmt.Sprintf("UplinkOnly")
	case Bidirectional:
		return fmt.Sprintf("Bidirectional")
	default:
		return fmt.Sprintf("invalid value")
	}
	return fmt.Sprintf("invalid value")
}

const (
	Reserved      = iota
	DownlinkOnly
	UplinkOnly
	Bidirectional
)

func (p *PacketFilterList) Decode(msgBuf *bytes.Reader) error {
	//octet 8
	octet8, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
		return fmt.Errorf("fail to read byte")
	}
	p.PktFilterIdentifier, _ = utils.GetBitsValue(octet8, 1, 4)
	direction, _ := utils.GetBitsValue(octet8, 5, 6)
	//fmt.Println("direction", direction >> 4)
	p.PktFilterDirection = PacketFilterDirection(direction >> 4)

	//fmt.Println("p.PktFilterDirection", p.PktFilterDirection)
	//fmt.Println("p.PktFilterIdentifier,", p.PktFilterIdentifier)

	p.PacketFilterContents.Decode(msgBuf)

	return nil
}

//encode PacketFilterList to nas octet stream
func (p *PacketFilterList) Encode() []byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte

	//octet 8
	octet8 := byte(p.PktFilterDirection) << 4
	octet8 |= byte(p.PktFilterIdentifier)
	encBuf = append(encBuf, octet8)

	contentValue := p.PacketFilterContents.Encode()

	//octet 9
	octet9 := byte(len(contentValue))
	encBuf = append(encBuf, octet9)

	//octet 10 ~ m
	encBuf = append(encBuf, contentValue[:]...)

	return encBuf
}
