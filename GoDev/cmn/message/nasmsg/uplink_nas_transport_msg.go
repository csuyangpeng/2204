package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"

	"github.com/willf/bitset"
)

// 9.11.3.39
const (
	SizeofPayloadContainer = 2
)

// 24.501 Table 8.2.10.1.1((V15.1.0 (2018-09)))
type UplinkNasTransportMsg struct {
	// mandatory
	PayloadType      nasie.PayloadContainerType
	PayloadContainer nasie.PayloadContainerIE
	// 1、当PayloadType != MultiplePayload 时，PayloadContainer []byte
	// 2、当PayloadType == MultiplePayload 时，PayloadContainer nasie.PayloadContainerIE
	// 可以把情况1时Payload container的内容放到p.PayloadContainer.PayloadContainerEntry[0].ContainerContents([]byte)里

	// optional
	PduSessId       nas.PduSessID
	OldPduSessId    nas.PduSessID
	RequestType     nasie.PduSessRequestType
	SNssai          nasie.SNssai
	Dnn             types3gpp.Apn
	AdditionalInfor []byte

	// optional ie indicator
	OptIeBitSet bitset.BitSet
}

//type IeId uint
const (
	IeidUplinknastransPdusessid uint = iota
	IeidUplinknastransOldpdusessid
	IeidUplinknastransRequesttype
	IeidUplinknastransSnssai
	IeidUplinknastransDnn
	IeidUplinknastransAdditionalinfor
)

func (p *UplinkNasTransportMsg) Reset() {
	p.PayloadType = 0
	p.PayloadContainer.Reset()
	p.PduSessId = 0
	p.OldPduSessId = 0
	p.RequestType = 0
	p.SNssai.Reset()
	p.Dnn.Reset()
	p.AdditionalInfor = []byte{}
	p.OptIeBitSet = bitset.BitSet{}
}

func (p *UplinkNasTransportMsg) String() string {
	var msgStr string
	msgStr = "NAS Msg - Uplink Nas Transport Message ( "
	msgStr += fmt.Sprintf("PayloadType(%s) ", p.PayloadType)
	msgStr += fmt.Sprintf("PduSessId(%d) ", uint8(p.PduSessId))
	msgStr += fmt.Sprintf("OldPduSessId(%d) ", uint8(p.OldPduSessId))
	msgStr += fmt.Sprintf("RequestType(%s) ", p.RequestType)
	msgStr += fmt.Sprintf("SNssai(%s) ", p.SNssai)
	msgStr += fmt.Sprintf("Dnn(%s) ", p.Dnn)
	return msgStr
}

func (p *UplinkNasTransportMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	// mandatory IEs
	//Payload container type
	payloadTypeBytes, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to decode Payload container type")
	}
	payloadType, _ := utils.GetBitsValue(payloadTypeBytes, 1, 4)
	switch payloadType {
	case 1:
		p.PayloadType = nasie.N1SmInformation
	case 2:
		p.PayloadType = nasie.SMSCont
	case 3:
		p.PayloadType = nasie.LPPMsg
	case 4:
		p.PayloadType = nasie.SORTransCont
	case 5:
		p.PayloadType = nasie.UePolicyCont
	case 6:
		p.PayloadType = nasie.UeParaUpdateTransCont
	case 15:
		p.PayloadType = nasie.MultiplePayload
	default:
		return fmt.Errorf("failed to analysis Payload container type value")
	}

	// Payload container
	// Payload container
	if p.PayloadType == nasie.MultiplePayload {
		err = p.PayloadContainer.Decode(msgBuf)
		if err != nil {
			return fmt.Errorf("failed to decode Payload container, %s", err)
		}
	} else {
		// L
		plcOctet := make([]byte, 2)
		err = binary.Read(msgBuf, binary.BigEndian, &plcOctet)
		if err != nil {
			return fmt.Errorf("failed to decode length for Payload container")
		}
		payloadValueLen := binary.BigEndian.Uint16(plcOctet)
		// V
		payloadValue := make([]byte, payloadValueLen)
		err = binary.Read(msgBuf, binary.BigEndian, &payloadValue)
		if err != nil {
			return fmt.Errorf("failed to decode length for Payload container")
		}
		//container的内容放到p.PayloadContainer.PayloadContainerEntry[0].ContainerContents里面
		p.PayloadContainer.PayloadContainerEntry = make([]nasie.PayloadContainerENTRY, 1)
		p.PayloadContainer.PayloadContainerEntry[0].ContainerContents = payloadValue
	}

	// Optional IEs
	err = p.DecodeOptIes(msgBuf)
	if err != nil {
		return fmt.Errorf("decode failed, error(%s)", err)
	}

	return nil

}

func (p *UplinkNasTransportMsg) DecodeOptIes(msgBuf *bytes.Reader) error {

	// Optional IEs
	for {
		// T
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		// 第一个字节只有前四位是IE的标志，需要额外判断
		switch nasie.Iei(ieType & 0xF0) {
		case nasie.IeiRequestType:
			// 9.11.3.47
			// V
			requestTypeByte, _ := utils.GetBitsValue(ieType, 1, 3)
			p.RequestType = nasie.PduSessRequestType(requestTypeByte)
			p.OptIeBitSet.Set(IeidUplinknastransRequesttype)
		}

		// 第一个字节就是IE的标志，直接识别即可
		switch nasie.Iei(ieType) {
		case nasie.IeiPduSessId:
			// 9.11.3.41
			// V
			pduSessIdByte, _ := msgBuf.ReadByte()
			p.PduSessId = nas.PduSessID(pduSessIdByte)
			p.OptIeBitSet.Set(IeidUplinknastransPdusessid)
		case nasie.IeiOldPDUSession:
			// 9.11.3.41
			// V
			oldPduSessIdByte, _ := msgBuf.ReadByte()
			p.OldPduSessId = nas.PduSessID(oldPduSessIdByte)
			p.OptIeBitSet.Set(IeidUplinknastransOldpdusessid)
		case nasie.IeiSNSSAI:
			// 9.11.2.8
			// LV
			err := p.SNssai.Decode(msgBuf)
			if err != nil {
				return fmt.Errorf("failed to decode snssai")
			}
			p.OptIeBitSet.Set(IeidUplinknastransSnssai)
		case nasie.IeiDNN:
			// 9.11.3.21
			// LV
			err := p.Dnn.Decode(msgBuf)
			if err != nil {
				return fmt.Errorf("failed to decode dnn")
			}
			p.OptIeBitSet.Set(IeidUplinknastransDnn)
		case nasie.IeiAdditionalInformation:
			// 9.11.2.1
			// LV
			// L
			addInforLen, err := msgBuf.ReadByte()
			if err != nil {
				return fmt.Errorf("failed to decode len of AdditionalInformation")
			}
			// V
			p.AdditionalInfor = make([]byte, addInforLen)
			binary.Read(msgBuf, binary.BigEndian, &p.AdditionalInfor)
			p.OptIeBitSet.Set(IeidUplinknastransAdditionalinfor)
		}
	}
	return nil

}

func (p *UplinkNasTransportMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.ULNasTransport
	encBuf = header.Encode()
	rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "header: %v", encBuf)

	// mandatory IEs
	//PayloadType      nasie.PayloadContainerType
	encBuf = append(encBuf, byte(p.PayloadType))

	//PayloadContainer []byte // payload container
	if p.PayloadType == nasie.MultiplePayload {
		PayloadContainerByte := p.PayloadContainer.Encode()
		//L
		plcBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(plcBytes, uint16(len(PayloadContainerByte)))
		encBuf = append(encBuf, plcBytes...)
		// v
		encBuf = append(encBuf, PayloadContainerByte...)
	} else {
		//container的内容放到p.PayloadContainer.PayloadContainerEntry[0].ContainerContents里面
		containerCONTENT := p.PayloadContainer.PayloadContainerEntry[0].ContainerContents
		//L
		plcBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(plcBytes, uint16(len(containerCONTENT)))
		encBuf = append(encBuf, plcBytes...)
		// v
		encBuf = append(encBuf, containerCONTENT...)
	}

	// optional
	optIeOctet, err := p.EncodeOptIes()
	if err != nil {
		err = fmt.Errorf("failed to decode optional ies")
	}
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p *UplinkNasTransportMsg) EncodeOptIes() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	// for other optional IEs:
	for i, e := p.OptIeBitSet.NextSet(0); e; i, e = p.OptIeBitSet.NextSet(i + 1) {
		switch i {
		case IeidUplinknastransPdusessid:
			// T
			encBuf = append(encBuf, byte(nasie.IeiPduSessId))
			// V
			encBuf = append(encBuf, byte(p.PduSessId))
		case IeidUplinknastransOldpdusessid:
			// T
			encBuf = append(encBuf, byte(nasie.IeiOldPDUSession))
			// V
			encBuf = append(encBuf, byte(p.OldPduSessId))
		case IeidUplinknastransRequesttype:
			// T
			requestTypeByte := byte(nasie.IeiRequestType) & 0xF0
			// V
			requestTypeByte |= byte(p.RequestType)

			encBuf = append(encBuf, requestTypeByte)
		case IeidUplinknastransSnssai:
			snssaiBytes := p.SNssai.Encode()
			// T
			encBuf = append(encBuf, byte(nasie.IeiSNSSAI))
			// L
			encBuf = append(encBuf, byte(len(snssaiBytes)))
			// V
			encBuf = append(encBuf, snssaiBytes[:]...)
		case IeidUplinknastransDnn:
			// T
			encBuf = append(encBuf, byte(nasie.IeiDNN))
			// L
			encBuf = append(encBuf, byte(len(p.Dnn.Encode())))
			// V
			encBuf = append(encBuf, p.Dnn.Encode()[:]...)
		}
	}
	return encBuf, nil
}
