package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	T "lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
)

type RegistrationResult byte

const (
	Access3gpp        RegistrationResult = 1
	AccessNon3gpp     RegistrationResult = 2
	Access3gppNon3gpp RegistrationResult = 3
)

// refer to TS24.501 Table 8.2.7.1.1
type RegistrationAcceptMsg struct {
	RegResult       RegistrationResult
	AllowSmsOverNAS bool //false - not Allow

	Guti5g                     nasie.MobileIdentity
	EquivalentPlmns            T.PlmnList
	TaiList                    T.TAIList
	AllowedNssai               nasie.Nssai
	RejectNssai                nasie.Nssai
	ConfigNssai                nasie.Nssai
	NwFeatuerSupport           nasie.NetworkFeatureSupport
	PduSessStatus              nasie.SessionStatus
	PduSessReactResult         nasie.SessionStatus
	PduSessReactResultErrCause nasie.SessionReactErrCause
	LandInfo                   nasie.LadnInformation
	MicoIndication             bool                         //false - all plmn registartion area not allocated
	NwSlicingSubsChanged       nasie.NetSliceIndicationType // Network slicing subscription not changed
	ServiceAreaList            nasie.ServiceAreaList
	T3512                      nasie.GprsTimer3
	T3502                      nasie.GprsTimer2
	//TODO
	//Non3gppDeregitTimer        nasie.GprsTimer2
	//EmergencyNumList
	//ExtEmergencyNumList
	//SorTransContainer
	//EapMessage
	//NSSAIInclusionMode nasie.NSSAIInclusionModeType
	//OperatorDefinedAccessCategoryDefinitions
	//NegotiatedDRXParameters

	//Indicates whether an IE is assigned or it is an empty value
	IeiMark map[nasie.Iei]bool
}

//encode a registration accept msg from nas octet stream
func (p *RegistrationAcceptMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte

	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.RegistrationAccept
	encBuf = header.Encode()

	// mandatory IEs

	// 5GS registration result, refer to TS24.501 Figure 9.11.3.6.1
	var (
		regResultOctet []byte
		regValueLength byte = 0
	)
	//5GS registration result's format is LV , no T
	// L
	regResultOctet = append(regResultOctet, regValueLength)

	var ROctet3 byte
	ROctet3 |= byte(p.RegResult) & 0x03
	ROctet3 |= utils.BoolToByte(p.AllowSmsOverNAS) << 3

	// V
	regResultOctet = append(regResultOctet, ROctet3)

	// 5GS registration result's value contains only one byte, so length = 1
	regValueLength = 1
	regResultOctet[0] = regValueLength

	encBuf = append(encBuf, regResultOctet...)

	// Optional IEs
	optIeOctet, err := p.EncodeOptIes()
	if err != nil {
		err = fmt.Errorf("failed to decode optional ies. %s", err)
	}

	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p *RegistrationAcceptMsg) EncodeOptIes() ([]byte, error) {
	var encBuf []byte

	// Guti5g, refer to TS24.501 Figure 9.11.3.4.1
	if p.IeiMark[nasie.IeiGuti5G] == true {
		var (
			gutiOctet       []byte
			gutiValueLength byte = 0
		)
		//Guti5g's format is TLV
		gutiValue, _ := p.Guti5g.Encode()
		gutiValueLength = byte(len(gutiValue))

		// T
		gutiOctet = append(gutiOctet, byte(nasie.IeiGuti5G))
		// L
		lenID := make([]byte, 2)
		binary.BigEndian.PutUint16(lenID, uint16(gutiValueLength))
		gutiOctet = append(gutiOctet, lenID...)
		// V，value encode
		gutiOctet = append(gutiOctet, gutiValue[:]...)

		encBuf = append(encBuf, gutiOctet...)
	}

	// Equivalent PLMNs(list), refer to TS24.501 Figure 9.11.3.46.1
	if p.IeiMark[nasie.IeiEquivalentPLMNs] == true {
		var (
			plmnListOctet  []byte
			plmnListLength byte = 0
		)
		// T
		plmnListOctet = append(plmnListOctet, byte(nasie.IeiEquivalentPLMNs))
		// L
		plmnListOctet = append(plmnListOctet, plmnListLength)
		// V，value encode
		gutiValueOctet, err := p.EquivalentPlmns.Encode()
		if err != nil {
			return encBuf, fmt.Errorf("failed to encode EquivalentPlmns")
		}
		plmnListOctet = append(plmnListOctet, gutiValueOctet[:]...)

		plmnListLength = byte(len(gutiValueOctet))
		plmnListOctet[1] = plmnListLength

		encBuf = append(encBuf, plmnListOctet...)
	}

	//TAI list, refer to TS24.501 Figure 9.11.3.9.1
	if p.IeiMark[nasie.IeiTAIList] == true {
		var (
			taiListOctet  []byte
			taiListLength byte = 0
		)
		// T
		taiListOctet = append(taiListOctet, byte(nasie.IeiTAIList))
		// L
		taiListOctet = append(taiListOctet, taiListLength)
		// V, value encode
		taiListValueOctet, err := p.TaiList.Encode()
		if err != nil {
			return encBuf, fmt.Errorf("failed to encode TAI list")
		}
		taiListOctet = append(taiListOctet, taiListValueOctet...)

		taiListLength = byte(len(taiListValueOctet))
		taiListOctet[1] = taiListLength

		encBuf = append(encBuf, taiListOctet...)
	}

	//Allowed NSSAI , refer to TS24.501 Figure 9.11.3.37.1
	if p.IeiMark[nasie.IeiAllowedNSSAI] == true {
		var (
			nssaiOctet  []byte
			nssaiLength byte = 0
		)
		//T
		nssaiOctet = append(nssaiOctet, byte(nasie.IeiAllowedNSSAI))
		//L
		nssaiOctet = append(nssaiOctet, nssaiLength)
		//V
		nssaiOctet = append(nssaiOctet, p.AllowedNssai.Encode()[:]...)

		nssaiLength = byte(len(p.AllowedNssai.Encode()))
		nssaiOctet[1] = nssaiLength

		encBuf = append(encBuf, nssaiOctet...)
	}
	//MICO indication , refer to TS24.501 Figure 9.11.3.31.1
	if p.IeiMark[nasie.IeiMICOIndication] == true {
		// octet 1
		micoOctet := byte(nasie.IeiMICOIndication) //0xB0
		micoOctet |= byte(utils.BoolToByte(p.MicoIndication))

		encBuf = append(encBuf, micoOctet)
	}
	//Network slicing indication , refer to TS24.501 Figure 9.11.3.36
	if p.IeiMark[nasie.IeiNetworkSlicingIndication] == true {
		// octet 1
		sliceOctet := byte(nasie.IeiNetworkSlicingIndication)
		sliceOctet |= byte(utils.BoolToByte(p.NwSlicingSubsChanged.NSSCI))

		encBuf = append(encBuf, sliceOctet)
	}
	//Service area list ,refer to TS24.501 Figure 9.11.3.49.1
	if p.IeiMark[nasie.IeiServiceAreaList] == true {
		var (
			serAreaOctet  []byte
			serAreaLength byte = 0
		)
		//T
		serAreaOctet = append(serAreaOctet, byte(nasie.IeiServiceAreaList))
		//L
		serAreaOctet = append(serAreaOctet, serAreaLength)
		//V
		serAreaValueOctet, _ := p.ServiceAreaList.Encode()
		serAreaOctet = append(serAreaOctet, serAreaValueOctet[:]...)

		serAreaLength = byte(len(serAreaValueOctet))
		serAreaOctet[1] = serAreaLength

		encBuf = append(encBuf, serAreaOctet...)
	}
	//T3502 value *
	if p.IeiMark[nasie.IeiT3502Value] == true {
		var (
			t3502Octet   []byte
			t3502OLength byte = 0
		)
		//T
		t3502Octet = append(t3502Octet, byte(nasie.IeiT3502Value))
		//L
		t3502Octet = append(t3502Octet, t3502OLength)
		//V
		t3502ValueOctet := p.T3502.Encode()
		t3502Octet = append(t3502Octet, t3502ValueOctet[:]...)

		t3502OLength = byte(len(t3502ValueOctet))
		t3502Octet[1] = t3502OLength

		encBuf = append(encBuf, t3502Octet...)
	}

	//T3512 value *
	if p.IeiMark[nasie.IeiT3512Value] == true {
		var (
			t3512Octet   []byte
			t3512OLength byte = 0
		)
		//T
		t3512Octet = append(t3512Octet, byte(nasie.IeiT3512Value))
		//L
		t3512Octet = append(t3512Octet, t3512OLength)
		//V
		t3512ValueOctet := p.T3512.Encode()
		t3512Octet = append(t3512Octet, t3512ValueOctet[:]...)

		t3512OLength = byte(len(t3512ValueOctet))
		t3512Octet[1] = t3512OLength

		encBuf = append(encBuf, t3512Octet...)
	}

	//Rejected NSSAI
	//Configured NSSAI
	//5GS network feature support
	//PDU session status
	//PDU session reactivation result
	//PDU session reactivation result error cause
	//LADN information

	return encBuf, nil
}

// decode a registration request msg from nas octet stream
func (p *RegistrationAcceptMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	// mandatory IEs
	// registration result
	lenOctet3, _ := msgBuf.ReadByte()
	octet3, err := msgBuf.ReadByte()
	if octet3 != lenOctet3 {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "wrong length of 5GS registration result ie ")
		return fmt.Errorf("wrong length of 5GS registration result ie ")
	}
	regResult, err := utils.GetBitsValue(octet3, 1, 3)
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "fail to get register result value")
		return fmt.Errorf("fail to get register result value")
	}
	p.RegResult = RegistrationResult(regResult)
	smsAllowed, err := utils.GetBitValue(octet3, 4)
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "fail to get sms allowed value")
		return fmt.Errorf("fail to get sms allowed value")
	}
	p.AllowSmsOverNAS = smsAllowed

	// Optional IEs
	err = p.DecodeOptIes(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode optional ies : %s", err)
	}

	return nil
}

func (p *RegistrationAcceptMsg) DecodeOptIes(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	for {
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		switch nasie.Iei(ieType) {
		case nasie.IeiGuti5G:
			// 5GS mobile identity: chapter 9.11.3.4	5GS mobile identity
			err := p.Guti5g.Decode(msgBuf)
			if err != nil {
				return fmt.Errorf("failed to decode guti: %s ", err)
			}
			return nil //剩余的不再解码
		case nasie.IeiRejectedNSSAI:
			utils.SkipIe(msgBuf)
		case nasie.IeiConfiguredNSSAI:
			utils.SkipIe(msgBuf)
		case nasie.Iei5GSNetworkFeatureSupport:
			utils.SkipIe(msgBuf)
		case nasie.IeiPDUSessionStatus:
			utils.SkipIe(msgBuf)
		case nasie.IeiPDUSessionReactivationResult:
			utils.SkipIe(msgBuf)
		case nasie.IeiPDUSessionReactivationResultErrorCause:
			utils.SkipIe(msgBuf)
		case nasie.IeiLADNInformation:
			utils.SkipIe(msgBuf)
		case nasie.IeiNon3GPPDeRegistrationTimerValue:
			utils.SkipIe(msgBuf)
		case nasie.IeiEmergencyNumberList:
			utils.SkipIe(msgBuf)
		case nasie.IeiExtendedEmergencyNumberList:
			utils.SkipIe(msgBuf)
		case nasie.IeiSORTransparentContainer:
			utils.SkipIe(msgBuf)
		case nasie.IeiEAPMessage:
			utils.SkipIe(msgBuf)
		case nasie.IeiNSSAIInclusionMode:
			utils.SkipIe(msgBuf)
		case nasie.IeiOperatorDefinedAccessCategoryDefinitions:
			utils.SkipIe(msgBuf)
		case nasie.IeiNegotiatedDRXParameters:
			utils.SkipIe(msgBuf)
		}
	}
	return nil
}
