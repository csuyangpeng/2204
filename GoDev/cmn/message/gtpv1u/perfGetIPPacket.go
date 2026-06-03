package gtpv1u

import (
	"encoding/binary"
	"fmt"
)

func GetRawIPPacket(msg []byte) ([]byte, error) {
	//安全检查
	if len(msg) < 8 {
		return nil, fmt.Errorf("invalid message")
	}
	GtpMsg := GTPPDU{Gtpv1uHeader: Gtpv1uHeader{}}
	//make([]byte, len(m.Msgbuf)),

	//r := bytes.NewReader(msg)
	var topOctet uint8
	// 1 Octet
	//binary.Read(r, binary.BigEndian, &topOctet)
	topOctet = msg[0]
	GtpMsg.Version = topOctet >> 5
	//top 3 bits
	GtpMsg.PT = topOctet >> 4 & 1
	// 1 bit
	GtpMsg.Spare = topOctet >> 3 & 1
	// 1 bit
	GtpMsg.EFlag = topOctet >> 2 & 1
	// 1 bit
	GtpMsg.SFlag = topOctet >> 1 & 1
	// 1 bit
	GtpMsg.PNFlag = topOctet & 1
	// 1 bit
	// Message Type 2 Octet
	GtpMsg.MessageType = msg[1]
	//Length 3-4 Octet
	GtpMsg.Length = binary.BigEndian.Uint16(msg[2:4])

	//Tunnel Endpoint Identifier 5-8 Octet
	GtpMsg.TEID = binary.BigEndian.Uint32(msg[4:8])
	// 长度在1500范围内
	mFieldLen := 8
	oFieldLen := 0
	if GtpMsg.EFlag == 1 {
		oFieldLen = 8
	}

	offset := mFieldLen + oFieldLen
	if len(msg) <= (offset + 1) {
		return nil, fmt.Errorf("invalid message")
	}
	return msg[offset:], nil
}
