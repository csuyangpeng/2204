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

//24501 8.3.9 f40 2019-06

type PDUSessionModifyCommandMsg struct {
	EPD       nas.Epd
	MsgHeader nas.SmNasMessageHeader
	//optional
	SMCause            nas.Sm5gCause
	Ambr               nasie.SessionAmbr
	RQTimer            nasie.GprsTimer
	AlwaysOnPduSessReq bool
	RequestQosRules    nasie.QoSRules
	//Mapped EPS bearer contexts
	RequestQosFlowDesc nasie.QoSFlowsDesc
	//Extended protocol configuration options

	// Ie flags
	IeFlags bitset.BitSet
}

const (
	Ieid_PduSessionModCmd_SMCause uint = iota
	Ieid_PduSessionModCmd_Ambr
	Ieid_PduSessionModCmd_AlwaysOnPduSessReq
	Ieid_PduSessionModCmd_RQTimer
	Ieid_PduSessionModCmd_RequestQosRules
	Ieid_PduSessionModCmd_RequestQosFlowDesc
)

//encode a session release request msg from nas octet stream
func (p *PDUSessionModifyCommandMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte
	//Mandatory
	encBuf = append(encBuf, byte(p.EPD))
	//MsgHeader V
	msgHeaderValue, _ := p.MsgHeader.Encode()
	encBuf = append(encBuf, msgHeaderValue[:]...)
	// Optional IEs
	//NonMandatory IE
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case Ieid_PduSessionModCmd_SMCause:
			encBuf = append(encBuf, byte(nasie.IeiSMCause))
			encBuf = append(encBuf, byte(p.SMCause))
		case Ieid_PduSessionModCmd_Ambr:
			//fmt.Println(p.Ambr)
			sessionAMBRValue := p.Ambr.Encode()
			encBuf = append(encBuf, byte(nasie.IeiAmbr))
			encBuf = append(encBuf, byte(len(sessionAMBRValue)))
			encBuf = append(encBuf, sessionAMBRValue[:]...)
			//fmt.Println(byte(len(sessionAMBRValue)),sessionAMBRValue)
		case Ieid_PduSessionModCmd_RQTimer:
			// T
			encBuf = append(encBuf, byte(nasie.IeiRQTimerValue))
			// V
			encBuf = append(encBuf, p.RQTimer.Encode()[:]...)
		case Ieid_PduSessionModCmd_AlwaysOnPduSessReq:
			//only one byte
			bytes := byte(nasie.IeiAlwaysOnPDUSessionRequested) | utils.BoolToByte(p.AlwaysOnPduSessReq)
			encBuf = append(encBuf, bytes)
		case Ieid_PduSessionModCmd_RequestQosRules:
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
		case Ieid_PduSessionModCmd_RequestQosFlowDesc:
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

func (p *PDUSessionModifyCommandMsg) Decode(msgBuf *bytes.Reader) error {
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
			p.IeFlags.Set(Ieid_PduSessionModCmd_AlwaysOnPduSessReq)
			//fmt.Println("p.AlwaysOnPduSessReq",p.AlwaysOnPduSessReq)
		}
		// 第一个字节就是IE的标志，直接识别即可
		switch nasie.Iei(ieType) {
		case nasie.IeiAmbr:
			len, _ := msgBuf.ReadByte()
			if len != 6 {
				rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "ambr len wrong")
				return fmt.Errorf("ambr len wrong")
			}
			err = p.Ambr.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "ambr decode wrong")
				return fmt.Errorf("ambr decode wrong")
			}
			p.IeFlags.Set(Ieid_PduSessionModCmd_Ambr)
			//fmt.Println(p.Ambr)
		case nasie.IeiRQTimerValue:
			err = p.RQTimer.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "RQTimer decode wrong")
				return fmt.Errorf("RQTimer decode wrong")
			}
			p.IeFlags.Set(Ieid_PduSessionModCmd_RQTimer)
			//fmt.Println(p.RQTimer)
		case nasie.IeiSMCause: //TV
			cause, _ := msgBuf.ReadByte()
			p.SMCause = nas.Sm5gCause(cause)
			p.IeFlags.Set(Ieid_PduSessionModCmd_SMCause)
			//fmt.Println("p.SMCause",p.SMCause)
		case nasie.IeiQosRules: //TLV
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
			p.IeFlags.Set(Ieid_PduSessionModCmd_RequestQosRules)
		case nasie.IeiAuthorizedQoSFlowDescriptions: //TLV
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
			//fmt.Println("p.RequestQosFlowDesc",p.RequestQosFlowDesc)
			p.IeFlags.Set(Ieid_PduSessionModCmd_RequestQosFlowDesc)
		default:
			rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "not support yet")
		}
	}
	return nil
}
