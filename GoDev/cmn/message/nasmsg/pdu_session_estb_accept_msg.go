package nasmsg

import (
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

// refer to 24.501 8.3.2(V15.1.0 (2018-09))
type PduSessionEstbAcceptMsg struct {
	//Mandatory
	ExtendProtoDisc    nas.Epd
	MsgHeader          nas.SmNasMessageHeader
	SessionType        types3gpp.PduSessType
	SscMode            nas.SSCMode
	AuthorizedQoSRules nasie.QoSRules
	SessionAMBR        nasie.SessionAmbr

	//Optional
	SMCause                       nas.Sm5gCause
	PDUaddress                    nasie.PDUAddress
	RQTimerValue                  nasie.GprsTimer
	SNSSAI                        nasie.SNssai
	AlwaysOn                      bool
	AuthorizedQoSFlowDescriptions nasie.QoSFlowsDesc
	ExtendProtocolConfigOpt       []byte
	DNN                           types3gpp.Apn
	MappedEPSBearerContexts       uint16 // 17~1024, not develop yet

	//Indicates whether an IE is assigned or it is an empty value
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	Ieid_PduSessionEstbAcpt_SMCause uint = iota
	Ieid_PduSessionEstbAcpt_PDUaddress
	Ieid_PduSessionEstbAcpt_RQTimerValue
	Ieid_PduSessionEstbAcpt_SNSSAI
	Ieid_PduSessionEstbAcpt_AlwaysOn
	Ieid_PduSessionEstbAcpt_AuthorizedQoSFlowDescriptions
	Ieid_PduSessionEstbAcpt_ExtendProtocolConfigOptResp
	Ieid_PduSessionEstbAcpt_DNN
)

// Print Session Accept Msg
func (p PduSessionEstbAcceptMsg) String() string {
	var s string
	s = fmt.Sprintf("PduSessEstbAcceptMsg:")
	s += fmt.Sprintf("MsgHeader(%s),", &(p.MsgHeader))
	s += fmt.Sprintf("SessionType(%s),", p.SessionType)
	s += fmt.Sprintf("SscMode(%s),", p.SscMode)
	s += fmt.Sprintf("DNN(%s)", p.DNN)
	s += fmt.Sprintf("AuthorizedQoSRules(%s),", p.AuthorizedQoSRules)
	s += fmt.Sprintf("SessionAMBR(%s),", p.SessionAMBR)
	s += fmt.Sprintf("PDUAddress(%s),", p.PDUaddress)
	//s += fmt.Sprintf("AuthorizedQoSFlowDescriptions: ", saMsg.AuthorizedQoSFlowDescriptions)
	return s
}

//encode SessionAcceptMsg to nas octet stream
func (p *PduSessionEstbAcceptMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)

	var encBuf []byte
	//mandatory
	//EPD
	encBuf = append(encBuf, byte(p.ExtendProtoDisc))

	//MsgHeader V
	msgHeaderValue, _ := p.MsgHeader.Encode()
	encBuf = append(encBuf, msgHeaderValue[:]...)

	//SessionType V & SscMode  V (total: 1 octet)
	stSsc := byte(p.SscMode) << 4
	stSsc |= byte(p.SessionType)
	encBuf = append(encBuf, stSsc)

	//AuthorizedQoSRules  LV
	var rulesByte []byte
	for i := 0; i < len(p.AuthorizedQoSRules.QoSRules); i++ {
		//octet 4
		//rulesByte = append(rulesByte, p.AuthorizedQoSRules.QoSRules[i].QoSFlowIdentifier)
		rulesByte = append(rulesByte, p.AuthorizedQoSRules.QoSRules[i].QoSRuleID)

		authorizedQoSRulesValue, _ := p.AuthorizedQoSRules.QoSRules[i].Encode()
		authorizedQoSRulesLen := len(authorizedQoSRulesValue)
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "QosRule[%d]-(%x)", i, authorizedQoSRulesValue)
		// octet 5 ~ 6
		lengthBuf := make([]byte, 2)
		binary.BigEndian.PutUint16(lengthBuf, uint16(authorizedQoSRulesLen))
		rulesByte = append(rulesByte, lengthBuf[:]...)

		// octet 7 ~ m + 2*
		rulesByte = append(rulesByte, authorizedQoSRulesValue[:]...)
	}
	// L
	lengthBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(lengthBuf, uint16(len(rulesByte)))
	encBuf = append(encBuf, lengthBuf...)
	//V
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "QosRules Buffer (%x)", rulesByte)
	encBuf = append(encBuf, rulesByte...)

	//SessionAMBR  LV
	sessionAMBRValue := p.SessionAMBR.Encode()
	encBuf = append(encBuf, byte(len(sessionAMBRValue)))
	encBuf = append(encBuf, sessionAMBRValue[:]...)

	// Optional IEs
	optIeOctet, err := p.EncodeOptIes()
	if err != nil {
		err = fmt.Errorf("failed to decode optional ies")
	}
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p *PduSessionEstbAcceptMsg) EncodeOptIes() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)
	var encBuf []byte

	//Optional
	// for other optional IEs:
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case Ieid_PduSessionEstbAcpt_SMCause:
			// T
			encBuf = append(encBuf, byte(nasie.IeiSMCause))
			// V
			encBuf = append(encBuf, byte(p.SMCause))
		case Ieid_PduSessionEstbAcpt_PDUaddress:
			// T
			encBuf = append(encBuf, byte(nasie.IeiPDUAddress))
			pduAddrValue, _ := p.PDUaddress.Encode()
			// L
			encBuf = append(encBuf, byte(len(pduAddrValue)))
			// V
			encBuf = append(encBuf, pduAddrValue[:]...)
		case Ieid_PduSessionEstbAcpt_RQTimerValue:
			// T
			encBuf = append(encBuf, byte(nasie.IeiRQTimerValue))
			// V
			encBuf = append(encBuf, p.RQTimerValue.Encode()[:]...)
		case Ieid_PduSessionEstbAcpt_SNSSAI:
			// T
			encBuf = append(encBuf, byte(nasie.IeiSNSSAI))
			snssaiValue := p.SNSSAI.Encode()
			// L
			encBuf = append(encBuf, byte(len(snssaiValue)))
			// V
			encBuf = append(encBuf, snssaiValue[:]...)
		case Ieid_PduSessionEstbAcpt_AlwaysOn:
			// T
			alwaysOnValue := byte(nasie.IeiAlwaysOnPDUSessionIndication)
			alwaysOnValue |= utils.BoolToByte(p.AlwaysOn)
			// V
			encBuf = append(encBuf, alwaysOnValue)
		case Ieid_PduSessionEstbAcpt_AuthorizedQoSFlowDescriptions:
			// T
			encBuf = append(encBuf, byte(nasie.IeiAuthorizedQoSFlowDescriptions))

			var qosFlowDesEncBuf []byte
			for i := 0; i < len(p.AuthorizedQoSFlowDescriptions.Descr); i++ {
				AuthValue, _ := p.AuthorizedQoSFlowDescriptions.Descr[i].Encode()
				qosFlowDesEncBuf = append(qosFlowDesEncBuf, AuthValue...)
			}

			// L
			lengthBuf := make([]byte, 2)
			//binary.BigEndian.PutUint16(lengthBuf, uint16(len(p.AuthorizedQoSFlowDescriptions.Descr)))
			binary.BigEndian.PutUint16(lengthBuf, uint16(len(qosFlowDesEncBuf)))
			encBuf = append(encBuf, lengthBuf...)

			// V
			encBuf = append(encBuf, qosFlowDesEncBuf...)

		case Ieid_PduSessionEstbAcpt_ExtendProtocolConfigOptResp:
		case Ieid_PduSessionEstbAcpt_DNN:
			//DNN LV
			dnnBytes := p.DNN.Encode()
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "DNN(%s),before encode dnn(%x)", p.DNN, encBuf)
			encBuf = append(encBuf, byte(nasie.IeiDnn))
			encBuf = append(encBuf, byte(len(dnnBytes)))
			encBuf = append(encBuf, dnnBytes[:]...)
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "after encode dnn(%x)", encBuf)
		}
	}
	return encBuf, nil
}
