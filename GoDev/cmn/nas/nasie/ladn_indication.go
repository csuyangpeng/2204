package nasie

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/types3gpp"
)

//24.501 9.11.3.30 ladn DNN Indication
const (
	MaxLandDnnNuminIndIe = 8
)

type LadnIndication struct {
	Dnn []types3gpp.Apn
}

func (p *LadnIndication) Decode(msgBuf *bytes.Reader) error {
	var length uint16
	err := binary.Read(msgBuf, binary.LittleEndian, &length)
	if err != nil {
		return fmt.Errorf("failed to decode length for ladnIndication")
	}

	length -= 3
	for length > 0 {
		dnnLen, _ := msgBuf.ReadByte()
		msgBuf.UnreadByte()

		apn := types3gpp.Apn{}
		err = apn.Decode(msgBuf)
		if err != nil {
			return fmt.Errorf("failed to decode length for dnn in ladnIndication")
		}

		p.Dnn = append(p.Dnn, apn)
		length -= uint16(dnnLen)
	}

	return nil
}

func (p *LadnIndication) Reset() {
	for i := 0; i < len(p.Dnn); i++ {
		p.Dnn[i].Reset()
	}
}
