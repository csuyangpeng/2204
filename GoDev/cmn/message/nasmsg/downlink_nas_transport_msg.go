package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"

	"github.com/willf/bitset"
)

// 24501 Table 8.2.11.1.1
type DownLinkNasTransportMsg struct {
	// mandatory
	//Extended protocol discriminator	Extended protocol discriminator 9.2
	//Security header type	Security header type 9.3
	//DLNASMessageID   nas.MmMsgType
	PayloadType      nasie.PayloadContainerType //V
	PayloadContainer nasie.PayloadContainerIE   //LV-E
	// 1、当PayloadType != MultiplePayload 时，PayloadContainer []byte
	// 2、当PayloadType == MultiplePayload 时，PayloadContainer nasie.PayloadContainerIE
	// 可以把情况1时Payload container的内容放到p.PayloadContainer.PayloadContainerEntry[0].ContainerContents([]byte)里

	// optional
	PduSessId       nas.PduSessID    //TV
	AdditionalInfor []byte           //TLV
	FiveGSMcause    nas.Sm5gCause    //TV
	BackOffTimer    nasie.GprsTimer3 //TLV

	// optional ie indicator
	OptIeBitSet bitset.BitSet
}

//type IeId uint
const (
	IeidDownlinknastransPdusessid uint = iota
	IeidDownlinknastransAdditionalinfor
	IeidDownlinknastransFivegmmcause
	IeidDownlinknastransBackofftimer
)

func (p *DownLinkNasTransportMsg) Decode(msgBuf *bytes.Reader) error {
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
		p.PayloadContainer.PayloadContainerEntry[0].ContainerContents = payloadValue
	}

	// Optional IEs
	err = p.DecodeOptIes(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode optional ies")
	}

	return nil
}

func (p *DownLinkNasTransportMsg) DecodeOptIes(msgBuf *bytes.Reader) error {
	// Optional IEs
	for {
		// T
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		// 第一个字节就是IE的标志，直接识别即可
		switch nasie.Iei(ieType) {
		case nasie.IeiPduSessId:
			// 9.11.3.41
			// V
			pduSessIdByte, _ := msgBuf.ReadByte()
			p.PduSessId = nas.PduSessID(pduSessIdByte)
		case nasie.IeiAdditionalInformation:
			// 9.11.2.1
			// LV
			addInforLen, err := msgBuf.ReadByte()
			if err != nil {
				return fmt.Errorf("failed to decode len of AdditionalInformation")
			}
			// V
			p.AdditionalInfor = make([]byte, addInforLen)
			binary.Read(msgBuf, binary.BigEndian, &p.AdditionalInfor)
		case nasie.IeiFiveGMMcause:
			// 9.11.3.2
			// V
			FiveGMMcauseByte, err := msgBuf.ReadByte()
			if err != nil {
				return fmt.Errorf("failed to decode FiveGMMcause")
			}
			p.FiveGSMcause = nas.Sm5gCause(FiveGMMcauseByte)
		case nasie.IeiBackOffTimer:
			// 9.11.2.5
			// LV
			err := p.BackOffTimer.Decode(msgBuf)
			if err != nil {
				return fmt.Errorf("failed to decode len of BackOffTimer")
			}
		}
	}
	return nil

}

func (p *DownLinkNasTransportMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.DLNasTransport
	encBuf = header.Encode()

	// mandatory

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

func (p *DownLinkNasTransportMsg) EncodeOptIes() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	// for other optional IEs:
	for i, e := p.OptIeBitSet.NextSet(0); e; i, e = p.OptIeBitSet.NextSet(i + 1) {
		switch i {
		case IeidDownlinknastransPdusessid:
			// T
			encBuf = append(encBuf, byte(nasie.IeiPduSessId))
			// V
			encBuf = append(encBuf, byte(p.PduSessId))
		case IeidDownlinknastransAdditionalinfor:
			// T
			encBuf = append(encBuf, byte(nasie.IeiAdditionalInformation))
			// L
			encBuf = append(encBuf, byte(len(p.AdditionalInfor)))
			// V
			encBuf = append(encBuf, p.AdditionalInfor[:]...)
		case IeidDownlinknastransFivegmmcause:
			// T
			encBuf = append(encBuf, byte(IeidDownlinknastransFivegmmcause))
			// V
			encBuf = append(encBuf, byte(p.FiveGSMcause))
		case IeidDownlinknastransBackofftimer:
			backOffTimerBytes := p.BackOffTimer.Encode()
			// T
			encBuf = append(encBuf, byte(nasie.IeiSNSSAI))
			// L
			encBuf = append(encBuf, byte(len(backOffTimerBytes)))
			// V
			encBuf = append(encBuf, backOffTimerBytes[:]...)
		}
	}
	return encBuf, nil

}
