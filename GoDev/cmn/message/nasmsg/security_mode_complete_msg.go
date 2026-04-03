package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

//24501   8.2.26.1
type SecurityModeCompleteMsg struct {
	// optional
	// IMEISV request 9.11.3.28 O TLV 11
	ImeiSv [types3gpp.LenImeiSv]byte

	// NAS message container 9.11.3.33 O TLV 4-n
	//NasMsgContainer TODO support container msg here, register request or service request message

	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

func (p SecurityModeCompleteMsg) String() string {
	return fmt.Sprintf("security mod complete, imsisv(%x)", p.ImeiSv)
}

const (
	IeidSecmodcmpImeisv uint = iota
)

func (p *SecurityModeCompleteMsg) Reset() {
	p.IeFlags.ClearAll()
}

func (p *SecurityModeCompleteMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	// Optional IEs
	err := p.DecodeOptIes(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode optional ies")
	}

	return nil
}

func (p *SecurityModeCompleteMsg) DecodeOptIes(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	for {
		// IE的标识被编码进了第一个字节，所以要单独拎出来
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}

		switch nasie.Iei(ieType) {
		case nasie.IeiImsiSv:
			length := make([]byte, 2)
			err := binary.Read(msgBuf, binary.BigEndian, &length)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "fail to read len of ImsiSv")
				return fmt.Errorf("fail to read len of ImsiSv")
			}
			imeisvLen := binary.BigEndian.Uint16(length)
			if imeisvLen != uint16(types3gpp.LenImeiSv/2+1) {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "invalid imsisv length")
				return fmt.Errorf("invalid imsisv length")
			}

			byteOctec, err := msgBuf.ReadByte()
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "fail to read ImsiSv first byte")
				return fmt.Errorf("fail to read ImsiSv first byte")
			}
			p.ImeiSv[0] = byteOctec >> 4

			for i := 1; i < int(imeisvLen); i++ {
				digit, err := msgBuf.ReadByte()
				if err != nil {
					rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "fail to read digit of ImsiSv")
					return fmt.Errorf("fail to read digit of ImsiSv")
				}

				p.ImeiSv[i*2-1] = digit & 0x0F
				if 2*i != 16 { //skip the last digit
					p.ImeiSv[i*2] = (digit & 0xF0) >> 4
				}
			}

			p.IeFlags.Set(IeidSecmodcmpImeisv)
		case nasie.IeiNasMsgContainer:
			//length
			length := make([]byte, 2)
			err := binary.Read(msgBuf, binary.BigEndian, &length)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "fail to read len of ImsiSv")
				return fmt.Errorf("fail to read len of ImsiSv")
			}
			//value not used yet
			valueLen := binary.BigEndian.Uint16(length)
			value := make([]byte, valueLen)
			err = binary.Read(msgBuf, binary.BigEndian, &value)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "fail to read len of NasMsgContainer")
				return fmt.Errorf("fail to read len of NasMsgContainer")
			}
			//todo value:register request or service request msg
			rlogger.Trace(types.ModuleCmnMsg, rlogger.INFO, nil, "NasMsgContainer")
		default:
			//TODO decode the registration request message
		}
	}

	return nil
}
