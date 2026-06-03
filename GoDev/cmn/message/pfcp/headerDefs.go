package pfcp

import (
	"bytes"
	"encoding/binary"
	. "lite5gc/cmn/message/pfcp/utils"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)

/*Figure 7.2.2.1-1: General format of PFCP Header
Bits
Octets		8	7	6	5	4	   3	2	1
1		Version	      Spare	Spare Spare	MP	S
2		Message Type
3		Message Length (1st Octet)
4		Message Length (2nd Octet)
m to k(m+7)		If S flag is set to 1, then SEID shall be placed into octets 5-12. Otherwise, SEID field is not present at all.
n to (n+2)		Sequence Number
(n+3)		Spare*/

type PfcpHeader struct {
	Version     uint8 //which represent the PFCP version, shall be set to decimal 1 ("001").
	MPFlag      uint8 //the "MP" flag
	SFlag       uint8 //the SEID flag
	MessageType uint8 //Message type values are specified in Table 7.3-1 "Message types".
	//This field shall indicate the length of the message in octets excluding the mandatory part of the PFCP header (the first 4 octets).
	Length uint16 //header length	// The SEID (if present) and the Sequence Number shall be included in the length count.

	// If S flag is set to 1, then SEID shall be placed into octets 5-12.
	// Otherwise, SEID field is not present at all.
	SEID            uint64
	SequenceNumber  uint32 // 3字节，Spare 1字节
	MessagePriority uint8
}

// 7.2.2.2	PFCP Header for Node Related Messages
type PfcpHeaderforNode struct {
	Version     uint8 //shall be set to decimal 1 ("001").
	MessageType uint8

	Length uint16 //header length

	SequenceNumber uint32 // 3字节，Spare 1字节
}

// 7.2.2.3	PFCP Header for Session Related Messages
/*Bits
Octets		8	7	6	5	   4	  3	    2	1
1		Version	        Spare Spare Spare	MP	S=1
2		Message Type
3		Message Length (1st Octet)
4		Message Length (2nd Octet)
5		Session Endpoint Identifier (1st Octet)
6		Session Endpoint Identifier (2nd Octet)
7		Session Endpoint Identifier (3rd Octet)
8		Session Endpoint Identifier (4th Octet)
9		Session Endpoint Identifier (5th Octet)
10		Session Endpoint Identifier (6th Octet)
11		Session Endpoint Identifier (7th Octet)
12		Session Endpoint Identifier (8th Octet)
13		Sequence Number (1st Octet)
14		Sequence Number (2nd Octet)
15		Sequence Number (3rd Octet)
16		Message Priority	Spare*/

type PfcpHeaderforSession struct {
	Version     uint8 //shall be set to decimal 1 ("001").
	MPFlag      uint8 //the "MP" flag
	SFlag       uint8 //the SEID flag, S=1
	MessageType uint8

	Length uint16 //header length

	// If S flag is set to 1, then SEID shall be placed into octets 5-12.
	// Otherwise, SEID field is not present at all.
	SEID            uint64 //Session Endpoint Identifier
	SequenceNumber  uint32 // 3字节
	MessagePriority uint8
}

const (
	Version uint8 = 1
	Flag    uint8 = 1
)

func (p *PfcpHeaderforSession) Set() {
	p.Version = Version
	p.SFlag = Flag
	p.MessageType = PFCP_Session_Establishment_Request
}
func (p *PfcpHeader) Set() {
	p.Version = Version
	p.MessageType = PFCP_Session_Establishment_Request
}

func (h *PfcpHeader) MarshalBinary() (data []byte, err error) {
	//Octets		8	7	6	5	   4	  3	    2	1
	// 1		Version	        Spare Spare Spare	MP	S=1
	var buf [1024]byte
	encBuf := bytes.NewBuffer(buf[:0])
	//var tmpByte byte
	tmpByte := (h.Version << 5) | (h.MPFlag << 1) | (h.SFlag)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	//2		Message Type
	err = encBuf.WriteByte(h.MessageType)
	if err != nil {
		return
	}
	//binary.Write(encBuf,binary.BigEndian,h.MessageType)

	//3		Message Length (1st Octet)
	//4		Message Length (2nd Octet)
	h.Length += 4 //Sequence Number + Message Priority or Spare
	if h.SFlag == 1 {
		h.Length += 8
	}
	err = binary.Write(encBuf, binary.BigEndian, h.Length)
	if err != nil {
		return
	}

	// 7.2.2.3	PFCP Header for Session Related Messages
	if h.SFlag == 1 {
		//5	- 12		Session Endpoint Identifier (8 Octets)
		err = binary.Write(encBuf, binary.BigEndian, h.SEID)
		if err != nil {
			return
		}
	}
	// 7.2.2.2	PFCP Header for Node Related Messages
	if h.SFlag == 0 {
	}
	//13		Sequence Number (1st Octet)
	//14		Sequence Number (2nd Octet)
	//15		Sequence Number (3rd Octet)
	byteThree := [3]byte{}
	byteThree[0] = byte(h.SequenceNumber >> 16)
	byteThree[1] = byte(h.SequenceNumber >> 8)
	byteThree[2] = byte(h.SequenceNumber)
	_, err = encBuf.Write(byteThree[:])
	if err != nil {
		return
	}

	//16		Message Priority	Spare*/
	if h.MPFlag == 1 {
		err = encBuf.WriteByte(h.MessagePriority << 4)
		if err != nil {
			return
		}
	} else {
		tmpByte := byte(0)
		err = encBuf.WriteByte(tmpByte)
		if err != nil {
			return
		}
	}
	return encBuf.Bytes(), nil
}

func (h *PfcpHeader) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)
	//	Octets		8	7	6	    5	   4	  3	    2	1
	//     1		Version	        Spare Spare Spare	MP	S
	var byteTmp byte
	err := binary.Read(r, binary.BigEndian, &byteTmp)
	if err != nil {
		return err
	}
	//decimal 1 ("001").
	v := ByteOne(byteTmp)
	//v.GetBits(6, 8)
	h.Version, _ = v.RightShift(6)
	h.MPFlag, _ = v.GetBit(2)
	h.SFlag, _ = v.GetBit(1)
	//2		Message Type
	err = binary.Read(r, binary.BigEndian, &byteTmp)
	if err != nil {
		return err
	}
	h.MessageType = byteTmp

	//3		Message Length (1st Octet)
	//4		Message Length (2nd Octet)
	err = binary.Read(r, binary.BigEndian, &h.Length)
	if err != nil {
		return err
	}

	// 7.2.2.3	PFCP Header for Session Related Messages
	if h.SFlag == 1 {
		//5	- 12		Session Endpoint Identifier (8 Octets)
		err = binary.Read(r, binary.BigEndian, &h.SEID)
		if err != nil {
			return err
		}

	}
	// 7.2.2.2	PFCP Header for Node Related Messages
	//if h.SFlag == 0 {
	//}
	//13		Sequence Number (1st Octet)
	//14		Sequence Number (2nd Octet)
	//15		Sequence Number (3rd Octet)

	byteThree := [3]byte{}
	err = binary.Read(r, binary.BigEndian, byteThree[:])
	if err != nil {
		return err
	}
	h.SequenceNumber = (uint32(byteThree[0]) << 16) +
		(uint32(byteThree[1]) << 8) +
		(uint32(byteThree[2]))

	//16		Message Priority	Spare
	if h.MPFlag == 1 {
		err = binary.Read(r, binary.BigEndian, &h.MessagePriority)
		if err != nil {
			return err
		}
	} else {
		r.ReadByte()
	}

	return nil
}

func (h *PfcpHeader) Len() int { return 0 }
