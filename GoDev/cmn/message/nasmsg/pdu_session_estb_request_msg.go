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

// refer to 24.501 8.3(V15.1.0 (2018-09))
type PduSessionEstbRequestMsg struct {
	//Mandatory
	Psi              nas.PduSessID
	Pti              nas.PrcdTransID
	InterMaxDataRate nas.IntergrityMaxDataRate //Integrity protection

	//optional
	SessionType    types3gpp.PduSessType
	SscMode        nas.SSCMode
	SmCapability   nas.SMCapability
	MaxNumberOfSPF uint16 // 24.501 9.11.4.9
	AlwaysOn       bool
	// SM PDU DN request container //TBD, for the PDU Session authorization by the external DN
	ExtendProtocolConfigOpt []byte

	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	Ieid_PduSessionEstbReq_SessionType uint = iota
	Ieid_PduSessionEstbReq_SscMode
	Ieid_PduSessionEstbReq_SmCapability
	Ieid_PduSessionEstbReq_MaxNumberOfSPF
	Ieid_PduSessionEstbReq_AlwaysOn
	Ieid_PduSessionEstbReq_ExtendProtocolConfigOptReq
)

func (p *PduSessionEstbRequestMsg) Reset() {
	//todo
}

// Print Registration Session Msg
func (srMsg PduSessionEstbRequestMsg) String() string {
	var s string
	//s = fmt.Sprintf("Session Request Msg info:")
	//s += fmt.Sprintf("psi: %s", srMsg.Psi)
	//s += fmt.Sprintf("pti: %s", srMsg.Pti)
	//s += fmt.Sprintf("MaxDataRate: %s", srMsg.InterMaxDataRate)
	//s += fmt.Sprintf("SessionType: %s", srMsg.SessionType)
	//s += fmt.Sprintf("SscMode: %s", srMsg.SscMode)
	//s += fmt.Sprintf("SmCapability: %s", srMsg.SmCapability)
	//s += fmt.Sprintf("MaxNumberOfSPF: %s", srMsg.MaxNumberOfSPF)
	//s += fmt.Sprintf("AlwaysOn: %s", srMsg.AlwaysOn)
	//s += fmt.Sprintf("ExtendProtocolConfigOpt: %s", srMsg.ExtendProtocolConfigOpt)
	return s
}

//encode a session request msg from nas octet stream
func (p *PduSessionEstbRequestMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte

	//Mandatory
	//EPD
	encBuf = append(encBuf, byte(nas.Epd5gsSessMgntMsg))

	//MsgHeader V
	header := nas.SmNasMessageHeader{}
	header.PduSessionID = p.Psi
	header.PrcdTransactionID = p.Pti
	header.MessageType = nas.PduSessEstabishRequest
	msgHeaderValue, _ := header.Encode()

	encBuf = append(encBuf, msgHeaderValue[:]...)

	//InterMaxDataRate
	encBuf = append(encBuf, byte(p.InterMaxDataRate.MaxRateUpLink))
	encBuf = append(encBuf, byte(p.InterMaxDataRate.MaxRateDownLink))

	// Optional IEs
	optIeOctet, err := p.EncodeOptIes()
	if err != nil {
		err = fmt.Errorf("failed to decode optional ies")
	}
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p *PduSessionEstbRequestMsg) EncodeOptIes() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte

	// for other optional IEs:
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case Ieid_PduSessionEstbReq_SessionType:
			typeOctet := byte(nasie.IeiPDUSessionType)
			typeOctet |= byte(p.SessionType)
			encBuf = append(encBuf, typeOctet)
		case Ieid_PduSessionEstbReq_SscMode:
			sscOctet := byte(nasie.IeiSSCMode)
			sscOctet |= byte(p.SscMode)
			encBuf = append(encBuf, sscOctet)
		case Ieid_PduSessionEstbReq_SmCapability:
			var smCapabilityValueOctet []byte
			tmpOctet := utils.BoolToByte(p.SmCapability.MPTCP) << 4
			tmpOctet |= utils.BoolToByte(p.SmCapability.ATSLL) << 3
			tmpOctet |= utils.BoolToByte(p.SmCapability.EPTS1) << 2
			tmpOctet |= utils.BoolToByte(p.SmCapability.MH6PDU) << 1
			tmpOctet |= utils.BoolToByte(p.SmCapability.RqoS)
			smCapabilityValueOctet = append(smCapabilityValueOctet, tmpOctet)
			smCapabilityValueOctet = append(smCapabilityValueOctet, p.SmCapability.Spare[:]...)
			//T
			encBuf = append(encBuf, byte(nasie.Iei5GSMCapability))
			//L
			encBuf = append(encBuf, byte(len(smCapabilityValueOctet)))
			//V
			encBuf = append(encBuf, smCapabilityValueOctet[:]...)
		case Ieid_PduSessionEstbReq_MaxNumberOfSPF:
			//T
			encBuf = append(encBuf, byte(nasie.IeiMaximumNumberOfSupportedPacketFilters))
			//V
			if p.MaxNumberOfSPF < nas.MinNumOfSPF || p.MaxNumberOfSPF > nas.MaxNumOfSPF {
				return encBuf, fmt.Errorf("maximum number of supported packet filters out of range")
			} else {
				// refer to 24.501 9.11.4.9
				//octet 2
				octet2 := p.MaxNumberOfSPF >> 2
				octet2 &= 0x00FF
				encBuf = append(encBuf, byte(octet2))
				//octet 3
				octet3 := p.MaxNumberOfSPF << 6
				octet3 &= 0x00C0
				encBuf = append(encBuf, byte(octet3))
			}
		case Ieid_PduSessionEstbReq_AlwaysOn:
			alwaysOctet := byte(nasie.IeiAlwaysOnPDUSessionRequested)
			alwaysOctet |= byte(utils.BoolToByte(p.AlwaysOn))
			encBuf = append(encBuf, alwaysOctet)
		case Ieid_PduSessionEstbReq_ExtendProtocolConfigOptReq:
			//ExtendProtocolConfigOpt TLV-E 4~65538
			//if p.IeiMark[nasie.IeiExtendedProtocolConfigurationOptions] == true {
			//	//T
			//	encBuf = append(encBuf, byte(nasie.IeiExtendedProtocolConfigurationOptions))
			//	//L
			//	//todo
			//	lenNum := uint16(len(p.ExtendProtocolConfigOpt)) // uint16
			//	lenNumBytes := utils.Uint16ToBytes(lenNum)
			//	encBuf = append(encBuf,lenNumBytes[:]...)
			//	//V
			//	encBuf = append(encBuf,p.ExtendProtocolConfigOpt[:]...)
			//}
		}
	}

	return encBuf, nil
}

// decode a session request msg from nas octet stream
func (p *PduSessionEstbRequestMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	// mandatory IEs
	//already decode epd and header before
	// the header have already decoded in : nas layer / ingress / smfNasIngress
	//ExtendProtoDisc nas.Epd
	//epd, err := msgBuf.ReadByte()
	//if err != nil {
	//	rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to decode epd, err:",err)
	//	return fmt.Errorf("failed to read bytes when decode epd")
	//}
	//p.ExtendProtoDisc = nas.Epd(epd)
	//MsgHeader       nas.SmNasMessageHeader
	//err = p.MsgHeader.Decode(msgBuf)
	//if err != nil {
	//	rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to decode header, err:",err)
	//	return fmt.Errorf("failed to decode header")
	//}

	//InterMaxDataRate
	maxRateByteUp, _ := msgBuf.ReadByte()
	p.InterMaxDataRate.MaxRateUpLink = nas.MAXDataRate(maxRateByteUp)
	maxRateByteDown, _ := msgBuf.ReadByte()
	p.InterMaxDataRate.MaxRateDownLink = nas.MAXDataRate(maxRateByteDown)

	// Optional IEs
	for {
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		// IE的标识被编码进了第一个字节，所以要单独拎出来
		switch nasie.Iei(ieType & 0xF0) {
		case nasie.IeiPDUSessionType:
			//SessionType TV 1
			sessionType, _ := utils.GetBitsValue(ieType, 1, 3)
			p.SessionType = types3gpp.PduSessType(sessionType)
		case nasie.IeiSSCMode:
			//SscMode TV 1
			//SmCapability TLV 3~15
			sscMode, _ := utils.GetBitsValue(ieType, 1, 3)
			p.SscMode = nas.SSCMode(sscMode)
		case nasie.IeiAlwaysOnPDUSessionRequested:
			//AlwaysOn TV 1
			sscMode, _ := utils.GetBitValue(ieType, 1)
			p.AlwaysOn = sscMode
		}

		// 第一个字节就是IE的标志，直接识别即可
		switch nasie.Iei(ieType) {
		case nasie.Iei5GSMCapability:
			//SmCapability TLV 3~15
			lenSmCapab, _ := msgBuf.ReadByte()
			octet3, _ := msgBuf.ReadByte()
			p.SmCapability.RqoS, _ = utils.GetBitValue(octet3, 1)
			p.SmCapability.MH6PDU, _ = utils.GetBitValue(octet3, 2)
			p.SmCapability.EPTS1, _ = utils.GetBitValue(octet3, 3)
			p.SmCapability.ATSLL, _ = utils.GetBitValue(octet3, 4)
			p.SmCapability.MPTCP, _ = utils.GetBitValue(octet3, 5)
			if lenSmCapab > 1 {
				spare := make([]byte, lenSmCapab-1)
				binary.Read(msgBuf, binary.BigEndian, spare)
				p.SmCapability.Spare = spare
			}
			p.IeFlags.Set(Ieid_PduSessionModReq_SmCapability)
			fmt.Println("p.SmCapability", p.SmCapability)
		case nasie.IeiMaximumNumberOfSupportedPacketFilters:
			//MaxNumberOfSPF TV 3
			// refer to 24.501 9.11.4.9
			octet2, _ := msgBuf.ReadByte()
			octet3, _ := msgBuf.ReadByte()

			octet3Bit12 := octet3 >> 6
			octet3Bit12Uint16 := uint16(octet3Bit12)

			octet2Uint16 := uint16(octet2)
			octet2Uint16Bit3To10 := octet2Uint16 << 2

			p.MaxNumberOfSPF = octet2Uint16Bit3To10 | octet3Bit12Uint16
		case nasie.IeiSMPDUDNRequestContainer:
			utils.SkipIe(msgBuf)
		case nasie.IeiExtendedProtocolConfigurationOptions:
			utils.SkipIe(msgBuf)
		}
	}
	return nil
}
