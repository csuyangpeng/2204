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
)

type RegistrationRequestMsg struct {
	CmnInfo nas.CmnNasInfo

	ForPending bool //false - no follow-on request pending
	RegType    nas.RegistrationType

	NgKSI          nasie.NasKSI
	MobileIdentity nasie.MobileIdentity
	//NonCurNativeKSI uint8 //inter-system from s1 to n1, ue should mapped 5g nas sec ctxt to protect regist request

	MmCapbility    nasie.MmCapability5G
	UeSecCapablity types3gpp.SecurityCapability //nasie.UeSecurityCapability

	RequestNssai   nasie.Nssai
	LastVisitedTAI types3gpp.TAI

	//S1UeNwCapbility S1UeNetworkCapbility  //TODO supported in Phase 3
	UplinkDataStatus nasie.SessionStatus
	PduSessStatus    nasie.SessionStatus
	MicoIndication   bool //false - all plmn registartion area not allocated

	//UeStatus UeStatus  //used for interworking with EPS
	//AdditionalGuti types3gpp.Guti5G  // for inter-system
	AllowedPduSessStatus nasie.SessionStatus
	UeUsageSetting       nas.UeUsage //This IE shall be included if the UE is configured to support IMS voice.

	//RequestedDRX  //TODO spec is not ready
	//EPS NAS message container

	LadnInd nasie.LadnIndication
	//PayloadContainer PayloadContainer  //for s1 mode to n1 mode
	//8-Payload container type	Payload container type

	NetSliceIndication nasie.NetSliceIndicationType
	FiveGUpdate        nasie.FiveGUpdateType
	NasMsgContainer    []byte
	//Indicates whether an IE is assigned or it is an empty value
	IeiMark map[nasie.Iei]bool
}

func (p *RegistrationRequestMsg) Reset() {
	p.CmnInfo.Reset()
	p.ForPending = false
	p.RegType = 0
	p.NgKSI.Reset()
	p.MobileIdentity.Reset()
	p.MmCapbility.Reset()
	p.UeSecCapablity.Reset()
	p.RequestNssai.Reset()
	p.LastVisitedTAI.Reset()
	p.UplinkDataStatus.Reset()
	p.PduSessStatus.Reset()
	p.MicoIndication = false
	p.AllowedPduSessStatus.Reset()
	p.UeUsageSetting = 0
	p.LadnInd.Reset()
	p.NetSliceIndication.Reset()
	p.FiveGUpdate.Reset()
	p.NasMsgContainer = []byte{}
	p.IeiMark = make(map[nasie.Iei]bool)

}

// Print Registration Request Msg
func (p *RegistrationRequestMsg) String() string {
	var msgStr string
	msgStr = "Registration Request Message ( "
	msgStr += fmt.Sprintf("RegistrationType(%d) ", p.RegType)
	msgStr += fmt.Sprintf("ForPending(%v) ", p.ForPending)
	msgStr += fmt.Sprintf("MobileIdentity(%s) ", p.MobileIdentity)
	msgStr += fmt.Sprintf("RequestNssai(%s) ", p.RequestNssai)
	msgStr += fmt.Sprintf("MicoIndication(%v) ", p.MicoIndication)
	return msgStr
}

// decode a registration request msg from nas octet stream
func (p *RegistrationRequestMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	// mandatory IEs
	// RegistrationType && ngKSI
	regKSI, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read registration type and ngKSI")
	}
	regType, _ := utils.GetBitsValue(regKSI, 1, 3)
	p.RegType = nas.RegistrationType(regType)
	p.ForPending, _ = utils.GetBitValue(regKSI, 4)
	p.NgKSI.Ksi, _ = utils.GetBitsValue(regKSI, 5, 7)
	p.NgKSI.Tsc, _ = utils.GetBitValue(regKSI, 8)

	// 5GS mobile identity: chapter 9.11.3.4	5GS mobile identity
	err = p.MobileIdentity.Decode(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode mobile identity,  %s", err)
	}

	// Optional IEs
	err = p.DecodeOptIes(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode optional ies")
	}

	return nil
}

func (p *RegistrationRequestMsg) DecodeOptIes(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	for {
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		// IE的标识被编码进了第一个字节，所以要单独拎出来

		switch nasie.Iei(ieType & 0xF0) {
		case nasie.IeiNoncurNativeNasKSI:
			utils.SkipIeOneByte(msgBuf)
			// The UE shall include this IE if the UE has a valid non-current native 5G NAS security context
			// when the UE performs a inter-system change from S1 mode to N1 mode in 5GMM-CONNECTED mode and
			// the UE uses a mapped 5G NAS security context to protect the REGISTRATION REQUEST message.
			// Don't support inter-system scenario yet
		case nasie.IeiMicoIndication:
			p.MicoIndication, _ = utils.GetBitValue(ieType, 1)
		}

		// 第一个字节就是IE的标志，直接识别即可
		switch nasie.Iei(ieType) {
		case nasie.Iei5GmmCapability:
			err = p.MmCapbility.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode 5GmmCapability failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiUeSecurityCapability:
			err = p.UeSecCapablity.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode UeSecCapability failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiRequestNssai:
			err = p.RequestNssai.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode Request Nssai failed")
				return nas.ErrDecodeNasIeFailed
			}

		case nasie.IeiLastVisitedRegTAI:
			err = p.LastVisitedTAI.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode Last Visited TAI failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiUplinkDataStatus:
			err = p.UplinkDataStatus.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode upLink Data status failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiPduSessionStatus:
			err = p.UplinkDataStatus.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode upLink Data status failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiAllowedPduSessStatus:
			utils.SkipIe(msgBuf)
		case nasie.IeiUeUsageSetting:
			octet, _ := msgBuf.ReadByte()

			rt, _ := utils.GetBitValue(octet, 1)
			if rt == true {
				p.UeUsageSetting = nas.DataCentric
			} else {
				p.UeUsageSetting = nas.VoiceCentric
			}
		case nasie.IeiLandIndication:
			err = p.LadnInd.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode LADN indication failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiS1UeNwCapability:
			utils.SkipIe(msgBuf)
		case nasie.IeiUeStatus:
			utils.SkipIe(msgBuf)
		case nasie.IeiAdditionalGuti:
			utils.SkipIe(msgBuf)
		case nasie.IeiReqDrxParameters:
			utils.SkipIe(msgBuf)
		case nasie.IeiEpsNasMsgContainer:
			utils.SkipIe(msgBuf)
		case nasie.IeiPayloadContainer:
			utils.SkipIe(msgBuf)
		case nasie.IeiNetSliceIndication:
			utils.SkipIe(msgBuf)
		case nasie.IeiFiveGUpdate:
			utils.SkipIe(msgBuf)
		case nasie.IeiNasMsgContainer:
			utils.SkipBigIe(msgBuf)
		default:
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "unknow IE Type(%x)", ieType)
		}
	}

	return nil
}

//encode a registration request msg from nas octet stream
func (p *RegistrationRequestMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	// mandatory IEs

	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.RegistrationRequest
	encBuf = append(encBuf, header.Encode()...)

	//RegistrationType && ngKSI
	regKSI := byte(utils.BoolToByte(p.NgKSI.Tsc)) << 7
	regKSI |= p.NgKSI.Ksi << 4
	regKSI |= utils.BoolToByte(p.ForPending) << 3
	regKSI |= byte(p.RegType)
	encBuf = append(encBuf, regKSI)

	//MobileIdentity, it's format is LV
	var mobileId []byte
	mobOctetValue, err := p.MobileIdentity.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "failed to encode mobile id, err(%s)", err)
		return nil, err
	}
	// L
	lenMobID := make([]byte, 2)
	binary.BigEndian.PutUint16(lenMobID, uint16(len(mobOctetValue)))
	mobileId = append(mobileId, lenMobID...)
	// V
	mobileId = append(mobileId, mobOctetValue...)
	encBuf = append(encBuf, mobileId...)

	// Optional IEs
	optIeOctet, err := p.EncodeOptIes()
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "failed to encode Ies, err(%s)", err)
		return nil, err
	}
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p *RegistrationRequestMsg) EncodeOptIes() ([]byte, error) {

	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	var encBuf []byte

	//RequestNssai
	if p.IeiMark[nasie.IeiRequestNssai] == true {
		var (
			nssaiOctet  []byte
			nssaiLength byte = 0
		)
		//T
		nssaiOctet = append(nssaiOctet, byte(nasie.IeiRequestNssai))
		//L
		nssaiOctet = append(nssaiOctet, nssaiLength)
		//V
		reValue := p.RequestNssai.Encode()
		nssaiOctet = append(nssaiOctet, reValue[:]...)

		nssaiLength = byte(len(reValue))
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

	return encBuf, nil
}
