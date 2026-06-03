package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

//24501 8.3.7 f40 2019-06
type PduSessionModifyRequestMsg struct {
	//mandatory
	Epd       nas.Epd
	MsgHeader nas.SmNasMessageHeader
	//optional
	SmCapability         nas.SMCapability
	SMCause              nas.Sm5gCause
	MaxNumOfSupPckFilter uint16 //9.11.4.9   range of 17 to 1024
	AlwaysOnPduSessReq   bool
	IntMaxDataRate       nas.IntergrityMaxDataRate
	RequestQosRules      nasie.QoSRules
	RequestQosFlowDesc   nasie.QoSFlowsDesc
	//Mapped EPS bearer contexts
	//Extended protocol configuration options
	// Ie flags
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	Ieid_PduSessionModReq_SmCapability uint = iota
	Ieid_PduSessionModReq_SMCause
	Ieid_PduSessionModReq_MaxNumOfSupPckFilter
	Ieid_PduSessionModReq_AlwaysOnPduSessReq
	Ieid_PduSessionModReq_IntMaxDataRate
	Ieid_PduSessionModReq_RequestQosRules
	Ieid_PduSessionModReq_RequestQosFlowDesc
)

func (p PduSessionModifyRequestMsg) String() string {
	var s string
	s = fmt.Sprintf("Session modification Msg info:")
	s += fmt.Sprintf("psi: %d ", p.MsgHeader.PduSessionID)
	s += fmt.Sprintf("pti: %d ", p.MsgHeader.PrcdTransactionID)
	s += fmt.Sprintf("smCapability: %v ", p.SmCapability)
	s += fmt.Sprintf("smCause: %x ", p.SMCause)
	s += fmt.Sprintf("MaxNumOfSupPckFilter: %d ", p.MaxNumOfSupPckFilter)
	s += fmt.Sprintf("AlwaysOnPduSessReq: %v ", p.AlwaysOnPduSessReq)
	s += fmt.Sprintf("IntMaxDataRate: %v ", p.IntMaxDataRate)
	s += fmt.Sprintf("RequestQosRules: %v ", p.RequestQosRules)
	s += fmt.Sprintf("RequestQosFlowDesc: %v ", p.RequestQosFlowDesc)
	return s
}

func (p *PduSessionModifyRequestMsg) Reset() {
	p.SmCapability.Reset()
	p.MaxNumOfSupPckFilter = 0
	p.AlwaysOnPduSessReq = false
	p.IntMaxDataRate.Reset()
}

//encode a session release request msg from nas octet stream
func (p *PduSessionModifyRequestMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	//Mandatory
	//MsgHeader V
	msgHeaderValue, _ := p.MsgHeader.Encode()
	encBuf = append(encBuf, msgHeaderValue[:]...)
	// Optional IEs
	//NonMandatory IE
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case Ieid_PduSessionModReq_SmCapability:
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
		case Ieid_PduSessionModReq_SMCause:
			encBuf = append(encBuf, byte(nasie.IeiSMCause))
			encBuf = append(encBuf, byte(p.SMCause))
		case Ieid_PduSessionModReq_MaxNumOfSupPckFilter:
			//T
			encBuf = append(encBuf, byte(nasie.IeiMaximumNumberOfSupportedPacketFilters))
			//V
			if p.MaxNumOfSupPckFilter < nas.MinNumOfSPF || p.MaxNumOfSupPckFilter > nas.MaxNumOfSPF {
				return encBuf, fmt.Errorf("maximum number of supported packet filters out of range")
			} else {
				// refer to 24.501 9.11.4.9
				//octet 2
				octet2 := p.MaxNumOfSupPckFilter >> 2
				octet2 &= 0x00FF
				encBuf = append(encBuf, byte(octet2))
				//octet 3
				octet3 := p.MaxNumOfSupPckFilter << 6
				octet3 &= 0x00C0
				encBuf = append(encBuf, byte(octet3))
			}
		case Ieid_PduSessionModReq_AlwaysOnPduSessReq:
			//only one byte
			bytes := byte(nasie.IeiAlwaysOnPDUSessionRequested) | utils.BoolToByte(p.AlwaysOnPduSessReq)
			encBuf = append(encBuf, bytes)
		case Ieid_PduSessionModReq_IntMaxDataRate:
			encBuf = append(encBuf, byte(nasie.IeiInterProctMaxDataRate))
			encBuf = append(encBuf, byte(p.IntMaxDataRate.MaxRateUpLink))
			encBuf = append(encBuf, byte(p.IntMaxDataRate.MaxRateDownLink))
		case Ieid_PduSessionModReq_RequestQosRules:
			//AuthorizedQoSRules  TLV
			var rulesByte []byte
			for i := 0; i < len(p.RequestQosRules.QoSRules); i++ {
				//octet 4
				//rulesByte = append(rulesByte, p.AuthorizedQoSRules.QoSRules[i].QoSFlowIdentifier)
				rulesByte = append(rulesByte, p.RequestQosRules.QoSRules[i].QoSRuleID)

				authorizedQoSRulesValue, _ := p.RequestQosRules.QoSRules[i].Encode()

				authorizedQoSRulesLen := len(authorizedQoSRulesValue)
				rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "QosRule[%d]-(%x)", i, authorizedQoSRulesValue)
				// octet 5 ~ 6
				lengthBuf := make([]byte, 2)
				binary.BigEndian.PutUint16(lengthBuf, uint16(authorizedQoSRulesLen))
				rulesByte = append(rulesByte, lengthBuf[:]...)

				// octet 7 ~ m + 2*
				rulesByte = append(rulesByte, authorizedQoSRulesValue[:]...)
			}
			// T
			encBuf = append(encBuf, byte(nasie.IeiQosRules))
			// L
			lengthBuf := make([]byte, 2)
			binary.BigEndian.PutUint16(lengthBuf, uint16(len(rulesByte)))
			encBuf = append(encBuf, lengthBuf...)
			//V
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "QosRules Buffer (%x)", rulesByte)
			encBuf = append(encBuf, rulesByte...)
		case Ieid_PduSessionModReq_RequestQosFlowDesc:
			// T
			encBuf = append(encBuf, byte(nasie.IeiAuthorizedQoSFlowDescriptions))

			var qosFlowDesEncBuf []byte
			for i := 0; i < len(p.RequestQosFlowDesc.Descr); i++ {
				AuthValue, _ := p.RequestQosFlowDesc.Descr[i].Encode()
				qosFlowDesEncBuf = append(qosFlowDesEncBuf, AuthValue...)
			}

			// L
			lengthBuf := make([]byte, 2)
			//binary.BigEndian.PutUint16(lengthBuf, uint16(len(p.AuthorizedQoSFlowDescriptions.Descr)))
			binary.BigEndian.PutUint16(lengthBuf, uint16(len(qosFlowDesEncBuf)))
			encBuf = append(encBuf, lengthBuf...)

			// V
			encBuf = append(encBuf, qosFlowDesEncBuf...)
		}
	}
	return encBuf, nil
}

func (p *PduSessionModifyRequestMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	// mandatory IEs
	// the header have already decoded

	//optional
	for {
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		// IE的标识被编码进了第一个字节，所以要单独拎出来
		switch nasie.Iei(ieType & 0xF0) {
		case nasie.IeiAlwaysOnPDUSessionRequested:
			//AlwaysOn TV 1
			apsr, _ := utils.GetBitValue(ieType, 1)
			p.AlwaysOnPduSessReq = apsr
			p.IeFlags.Set(Ieid_PduSessionModReq_AlwaysOnPduSessReq)
			//fmt.Println("p.AlwaysOnPduSessReq",p.AlwaysOnPduSessReq)
		}
		// 第一个字节就是IE的标志，直接识别即可
		switch nasie.Iei(ieType) {
		case nasie.Iei5GSMCapability:
			//TLV
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
			//fmt.Println("p.SmCapability",p.SmCapability)
		case nasie.IeiSMCause:
			//TV
			cause, _ := msgBuf.ReadByte()
			p.SMCause = nas.Sm5gCause(cause)
			p.IeFlags.Set(Ieid_PduSessionModReq_SMCause)
			//fmt.Println("p.SMCause",p.SMCause)
		case nasie.IeiMaximumNumberOfSupportedPacketFilters: //TV
			//MaxNumberOfSPF TV 3
			// refer to 24.501 9.11.4.9
			octet2, _ := msgBuf.ReadByte()
			octet3, _ := msgBuf.ReadByte()

			octet3Bit12 := octet3 >> 6
			octet3Bit12Uint16 := uint16(octet3Bit12)

			octet2Uint16 := uint16(octet2)
			octet2Uint16Bit3To10 := octet2Uint16 << 2

			p.MaxNumOfSupPckFilter = octet2Uint16Bit3To10 | octet3Bit12Uint16
			p.IeFlags.Set(Ieid_PduSessionModReq_MaxNumOfSupPckFilter)
			//fmt.Println("p.MaxNumOfSupPckFilter",p.MaxNumOfSupPckFilter)
		case nasie.IeiInterProctMaxDataRate: //TV
			maxRateByteUp, _ := msgBuf.ReadByte()
			p.IntMaxDataRate.MaxRateUpLink = nas.MAXDataRate(maxRateByteUp)
			maxRateByteDown, _ := msgBuf.ReadByte()
			p.IntMaxDataRate.MaxRateDownLink = nas.MAXDataRate(maxRateByteDown)
			p.IeFlags.Set(Ieid_PduSessionModReq_IntMaxDataRate)
			//fmt.Println("p.IntMaxDataRate",p.IntMaxDataRate)
		case nasie.IeiQosRules:
			//TLV
			//L    octet 2-3
			lenBytes := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, lenBytes)
			length := binary.BigEndian.Uint16(lenBytes)
			//fmt.Println("length",length)
			//V
			p.RequestQosRules.QoSRules = []nasie.QoSRule{}
			for i := 0; length > 0; i++ {
				qosrule := nasie.QoSRule{}

				//octet 4
				ruleID, err := msgBuf.ReadByte()
				if err != nil {
					rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "no more bytes")
					return fmt.Errorf("fail to read byte")
				}
				qosrule.QoSRuleID = ruleID
				//fmt.Println("qosrule.QoSRuleID ",qosrule.QoSRuleID )

				//octet 5-6   rule  length
				ruleLenBytes := make([]byte, 2)
				binary.Read(msgBuf, binary.BigEndian, ruleLenBytes)
				ruleLength := binary.BigEndian.Uint16(ruleLenBytes)
				//fmt.Println("ruleLenBytes ",ruleLenBytes )
				//fmt.Println("ruleLength ",ruleLength )

				length -= ruleLength + 3

				qosrule.Decode(msgBuf)

				p.RequestQosRules.QoSRules = append(p.RequestQosRules.QoSRules, qosrule)
			}
			//fmt.Println("p.RequestQosRules",p.RequestQosRules)
			p.IeFlags.Set(Ieid_PduSessionModReq_RequestQosRules)
		case nasie.IeiAuthorizedQoSFlowDescriptions:
			//TLV
			//L    octet 2-3
			lenBytes := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, lenBytes)
			length := binary.BigEndian.Uint16(lenBytes)
			//fmt.Println("length",length)
			//V
			p.RequestQosFlowDesc.Descr = []nasie.QoSFlowDescription{}
			for i := 0; length > 0; i++ {
				qosflow := nasie.QoSFlowDescription{}
				err, len := qosflow.Decode(msgBuf)
				//fmt.Println("~",err)
				if err != nil {
					rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "fail to decode Descr")
					return fmt.Errorf("fail to decode Descr")
				}
				//fmt.Println("len",len)
				length -= uint16(len)
				//fmt.Println("length",length)
				p.RequestQosFlowDesc.Descr = append(p.RequestQosFlowDesc.Descr, qosflow)
			}
			p.IeFlags.Set(Ieid_PduSessionModReq_RequestQosFlowDesc)
		default:
			rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "not support yet")
		}
	}
	return nil
}
