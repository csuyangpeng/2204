package nasie

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type SNssaiInd byte

const (
	SstOnly          SNssaiInd = 1
	SstMapSst        SNssaiInd = 2
	SstSd            SNssaiInd = 4
	SstSdMapSst      SNssaiInd = 5
	SstSdMapSstMapSd SNssaiInd = 8
)

const (
	SizeofSd = 3
)

//23.501 Table 5.15.2.2-1 - Standardised SST values
//Slice/Service type	SST value	Characteristics.
//eMBB  					1		Slice suitable for the handling of 5G enhanced Mobile Broadband.
//URLLC						2		Slice suitable for the handling of ultra- reliable low latency communications.
//MIoT						3		Slice suitable for the handling of massive IoT.

// When Snssai needs to be converted to string (e.g. when used in maps as key),
// the string shall be composed of one to three digits "sst" optionally followed by "-" and 6 hexadecimal digits "sd",
// and shall match the following pattern: ^([0-9]|[1-9][0-9]|1[0-9][0-9]|2([0-4][0-9]|5[0-5]))(-[A-Fa-f0-9]{6})?$
// Example 1:	"255-19CDE0"
// Example 2:	"29"
type SNssai struct {
	Ind       SNssaiInd
	Sst       byte
	Sd        [SizeofSd]byte
	MappedSst byte
	MappedSd  [SizeofSd]byte
}

func (p *SNssai) Reset() {
	p.Ind = 0
	p.Sst = 0
	p.Sd = [SizeofSd]byte{}
	p.MappedSst = 0
	p.MappedSd = [SizeofSd]byte{}
}

func (p SNssai) String() string {
	var out string
	switch p.Ind {
	case SstOnly:
		out = fmt.Sprintf("%d", p.Sst)
	case SstSd:
		out = fmt.Sprintf("%d-%s", p.Sst, hex.EncodeToString(p.Sd[:]))
	default:
	}
	return out
	//return fmt.Sprintf("%d.%d.0x%s.%d.0x%s",
	//	p.Ind,
	//	p.Sst,
	//	hex.EncodeToString(p.Sd[:]),
	//	p.MappedSst,
	//	hex.EncodeToString(p.MappedSd[:]))
}

func (p *SNssai) IsValid() bool {
	invalidSNssai := SNssai{}
	if p.String() == invalidSNssai.String() {
		return false
	}
	return true
}

func (p *SNssai) GetLen() (rt uint8) {
	switch p.Ind {
	case SstOnly:
		rt = 2
	case SstMapSst:
		rt = 3
	case SstSd:
		rt = 5
	case SstSdMapSst:
		rt = 6
	case SstSdMapSstMapSd:
		rt = 9
	default:
		rt = 0
	}
	return
}

func (p *SNssai) Encode() []byte {
	var sNssaiOctet []byte
	switch p.Ind {
	case SstOnly:
		sNssaiOctet = append(sNssaiOctet, p.Sst)
	case SstMapSst:
		sNssaiOctet = append(sNssaiOctet, p.Sst)
		sNssaiOctet = append(sNssaiOctet, p.MappedSst)
	case SstSd:
		sNssaiOctet = append(sNssaiOctet, p.Sst)
		sNssaiOctet = append(sNssaiOctet, p.Sd[:]...)
	case SstSdMapSst:
		sNssaiOctet = append(sNssaiOctet, p.Sst)
		sNssaiOctet = append(sNssaiOctet, p.Sd[:]...)
		sNssaiOctet = append(sNssaiOctet, p.MappedSst)
	case SstSdMapSstMapSd:
		sNssaiOctet = append(sNssaiOctet, p.Sst)
		sNssaiOctet = append(sNssaiOctet, p.Sd[:]...)
		sNssaiOctet = append(sNssaiOctet, p.MappedSst)
		sNssaiOctet = append(sNssaiOctet, p.MappedSd[:]...)
	}
	return sNssaiOctet
}

func (p *SNssai) Decode(msgBuf *bytes.Reader) error {
	ind, _ := msgBuf.ReadByte()
	// 读到的第一个字节是包含长度值的值，而p.Ind是不含长度值的，所以要减去1，才是value部分真正的长度
	p.Ind = SNssaiInd(ind)
	switch p.Ind {
	case SstOnly:
		p.Sst, _ = msgBuf.ReadByte()
	case SstMapSst:
		p.Sst, _ = msgBuf.ReadByte()
		p.MappedSst, _ = msgBuf.ReadByte()
	case SstSd:
		p.Sst, _ = msgBuf.ReadByte()
		binary.Read(msgBuf, binary.BigEndian, &(p.Sd))
	case SstSdMapSst:
		p.Sst, _ = msgBuf.ReadByte()
		binary.Read(msgBuf, binary.BigEndian, &(p.Sd))
		p.MappedSst, _ = msgBuf.ReadByte()
	case SstSdMapSstMapSd:
		p.Sst, _ = msgBuf.ReadByte()
		binary.Read(msgBuf, binary.BigEndian, &(p.Sd))
		p.MappedSst, _ = msgBuf.ReadByte()
		binary.Read(msgBuf, binary.BigEndian, &(p.MappedSd))
	default:
		return fmt.Errorf("invalid length for SNSSAI IE")
	}
	return nil
}