package sbicmn

import (
	"encoding/binary"
	"encoding/hex"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/udmdata"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/openapi/models"
	"net"
	"strconv"
)

func Trans_ModelsToN11_SMDataFormat(smData []models.SessionManagementSubscriptionData) (sm []udmdata.SessMgntSubscripitonData) {
	rlogger.FuncEntry(types.ModCmn, nil)
	sm = make([]udmdata.SessMgntSubscripitonData, len(smData))
	for k, v := range smData {
		snssai := nasie.SNssai{}
		snssai.Sst = byte(v.SingleNssai.Sst)
		if v.SingleNssai.Sd == "" {
			snssai.Ind = nasie.SstOnly
		} else {
			snssai.Ind = nasie.SstSd
			sd, _ := strconv.Atoi(v.SingleNssai.Sd)
			sdSlice := make([]byte, 4)
			binary.BigEndian.PutUint32(sdSlice, uint32(sd))
			copy(snssai.Sd[:], sdSlice[1:4])
		}
		sm[k].SingleNssai = snssai

		sm[k].DnnConfigs = make(map[string]*nasie.DNNConfiguration)
		for k2, v2 := range v.DnnConfigurations {
			dnnC := nasie.DNNConfiguration{}
			// DefaultPDUSessionType
			switch v2.PduSessionTypes.DefaultSessionType {
			case models.PduSessionType_IPV4:
				dnnC.DefaultPDUSessionType = types3gpp.Ipv4
			case models.PduSessionType_IPV6:
				dnnC.DefaultPDUSessionType = types3gpp.Ipv6
			case models.PduSessionType_IPV4_V6:
				dnnC.DefaultPDUSessionType = types3gpp.Ipv4v6
			case models.PduSessionType_UNSTRUCTURED:
				dnnC.DefaultPDUSessionType = types3gpp.Unstructured
			case models.PduSessionType_ETHERNET:
				dnnC.DefaultPDUSessionType = types3gpp.Ethernet
			default:
				dnnC.DefaultPDUSessionType = types3gpp.Ipv4
			}
			// AllowedSessionType
			allowPST := make([]types3gpp.PduSessType, len(v2.PduSessionTypes.AllowedSessionTypes))
			for k, v := range v2.PduSessionTypes.AllowedSessionTypes {
				switch v {
				case models.PduSessionType_IPV4:
					allowPST[k] = types3gpp.Ipv4
				case models.PduSessionType_IPV6:
					allowPST[k] = types3gpp.Ipv6
				case models.PduSessionType_IPV4_V6:
					allowPST[k] = types3gpp.Ipv4v6
				case models.PduSessionType_UNSTRUCTURED:
					allowPST[k] = types3gpp.Unstructured
				case models.PduSessionType_ETHERNET:
					allowPST[k] = types3gpp.Ethernet
				default:
					allowPST[k] = types3gpp.Ipv4
				}
			}
			dnnC.AllowedSessionType = allowPST
			// DefaultSSCMode
			switch v2.SscModes.DefaultSscMode {
			case models.SscMode__1:
				dnnC.DefaultSSCMode = nas.SSCMode1
			case models.SscMode__2:
				dnnC.DefaultSSCMode = nas.SSCMode2
			case models.SscMode__3:
				dnnC.DefaultSSCMode = nas.SSCMode3
			default:
				dnnC.DefaultSSCMode = nas.SSCMode1
			}
			// AllowedSSCMode
			allowSSC := make([]nas.SSCMode, len(v2.SscModes.AllowedSscModes))
			for k, v := range v2.SscModes.AllowedSscModes {
				switch v {
				case models.SscMode__1:
					allowSSC[k] = nas.SSCMode1
				case models.SscMode__2:
					allowSSC[k] = nas.SSCMode2
				case models.SscMode__3:
					allowSSC[k] = nas.SSCMode3
				default:
					allowSSC[k] = nas.SSCMode1
				}
			}
			dnnC.AllowedSSCMode = allowSSC
			// IwkEpsInd
			dnnC.IwkEpsInd = v2.IwkEpsInd
			// QosProf
			qosF := nasie.QosProfile{}
			switch v2.Var5gQosProfile.Var5qi {
			case 1:
				qosF.QI5 = nasie.QI1
			case 2:
				qosF.QI5 = nasie.QI2
			case 3:
				qosF.QI5 = nasie.QI3
			case 4:
				qosF.QI5 = nasie.QI4
			case 5:
				qosF.QI5 = nasie.QI5
			case 6:
				qosF.QI5 = nasie.QI6
			case 7:
				qosF.QI5 = nasie.QI7
			case 8:
				qosF.QI5 = nasie.QI8
			case 9:
				qosF.QI5 = nasie.QI9
			case 65:
				qosF.QI5 = nasie.QI65
			case 66:
				qosF.QI5 = nasie.QI66
			case 69:
				qosF.QI5 = nasie.QI69
			case 70:
				qosF.QI5 = nasie.QI70
			case 75:
				qosF.QI5 = nasie.QI75
			case 79:
				qosF.QI5 = nasie.QI79
			case 255:
				qosF.QI5 = nasie.QIMax
			default:
				qosF.QI5 = nasie.QI9
			}
			if v2.Var5gQosProfile.PriorityLevel != 0 {
				qosF.PriorityLevel = uint8(v2.Var5gQosProfile.PriorityLevel)
				qosF.IsPriorityLevelPrst = true
			}
			var arp types3gpp.ARP
			temp := v2.Var5gQosProfile.Arp.PriorityLevel
			arp.PriorityLevel = uint16(temp)
			switch v2.Var5gQosProfile.Arp.PreemptCap {
			case models.PreemptionCapability_NOT_PREEMPT:
				arp.PreemptCap = types3gpp.NOT_PREEMPT
			case models.PreemptionCapability_MAY_PREEMPT:
				arp.PreemptCap = types3gpp.MAY_PREEMPT
			default:
				arp.PreemptCap = types3gpp.NOT_PREEMPT
			}
			switch v2.Var5gQosProfile.Arp.PreemptVuln {
			case models.PreemptionVulnerability_NOT_PREEMPTABLE:
				arp.PreemptVuln = types3gpp.NOT_PREEMPTABLE
			case models.PreemptionVulnerability_PREEMPTABLE:
				arp.PreemptVuln = types3gpp.PREEMPTABLE
			default:
				arp.PreemptVuln = types3gpp.NOT_PREEMPTABLE
			}
			qosF.Arp = arp
			dnnC.QosProf = qosF
			// SessionAmbr
			var sambr nasie.SessionAmbr
			_ = sambr.StoreWithString(v2.SessionAmbr.Uplink, v2.SessionAmbr.Downlink)
			dnnC.SessionAmbr = sambr
			// ChargingCharacteristics
			dnnC.ChargingCharacteristics = v2.Var3gppChargingCharacteristics
			// StaticIpAddress
			ip := net.ParseIP(v2.StaticIpAddress[0].Ipv4Addr)
			dnnC.StaticIpAddress = ip
			// dnn
			sm[k].DnnConfigs[k2] = &dnnC
		}
	}
	return sm
}

//amf receive Nudm msg, trans modelsData to udmData
func Trans_ModelsToN11_AMSDDataFormat(modelsData *models.AccessAndMobilitySubscriptionData) (udmData udmdata.AccMobSubscribeData) {
	rlogger.FuncEntry(types.ModCmn, nil)
	//ambr
	if modelsData.SubscribedUeAmbr != nil {
		v, err := types3gpp.StoreAmbrBitRate(modelsData.SubscribedUeAmbr.Uplink)
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.WARN, nil, "failed to get ambr, error(%s)", err)
		} else {
			udmData.SubsUeAmbr.Uplink = v
		}
	}

	if modelsData.SubscribedUeAmbr != nil {
		v, err := types3gpp.StoreAmbrBitRate(modelsData.SubscribedUeAmbr.Downlink)
		if err != nil {
			rlogger.Trace(types.ModCmn, rlogger.WARN, nil, "failed to store ambr, error(%s)", err)
		} else {
			udmData.SubsUeAmbr.Downlink = v
		}
	}
	//nssai
	data := udmdata.Nssai{}
	snssai := nasie.SNssai{}
	if modelsData.Nssai == nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "failed to get nssai")
	} else {
		if len(modelsData.Nssai.DefaultSingleNssais) != 0 {
			v := modelsData.Nssai.DefaultSingleNssais[0]
			snssai.Sst = byte(v.Sst)
			if v.Sd == "" {
				snssai.Ind = nasie.SstOnly
			} else {
				snssai.Ind = nasie.SstSd
				sd, _ := strconv.Atoi(v.Sd)
				sdSlice := make([]byte, 4)
				binary.BigEndian.PutUint32(sdSlice, uint32(sd))
				copy(snssai.Sd[:], sdSlice[1:4])
			}
		}
		data.DefSnssai = snssai
	}
	udmData.Nssai = data
	return udmData
}

//amf receive Nudm msg, trans modelsData to udmData
func Trans_ModelsToN11_SmfSelDataFormat(modelsData *models.SmfSelectionSubscriptionData) (udmData udmdata.SmfSelSubscribeData) {
	rlogger.FuncEntry(types.ModCmn, nil)
	udmData.SnssaiInfoList = make(map[string]udmdata.SnssaiInfo)
	for k, v := range modelsData.SubscribedSnssaiInfos {
		dnnInfo := make([]types3gpp.DnnInfo, len(v.DnnInfos))
		for k, v := range v.DnnInfos {
			err := dnnInfo[k].Dnn.StoreWithString(v.Dnn)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail get dnn")
				return
			}
			dnnInfo[k].IwkEpsInd = v.IwkEpsInd
			dnnInfo[k].LboRoamingAllowed = v.LboRoamingAllowed
			dnnInfo[k].DefDnnInd = v.DefaultDnnIndicator
		}
		sn := udmdata.SnssaiInfo{}
		sn.DnnInfos = dnnInfo
		udmData.SnssaiInfoList[k] = sn
	}

	return udmData
}

//amf receive Nudm msg, trans modelsData to udmData
func Trans_ModelsToN11_UCSDDataFormat(modelsData *models.UeContextInSmfData) (udmData udmdata.UeContextInSmfData) {
	rlogger.FuncEntry(types.ModCmn, nil)
	return udmData
}

//amf receive Nudm msg
func Trans_ModelsToN11_AuthDataFormat(modelsData *models.AuthenticationInfoResult) types.HeAvType {
	rlogger.FuncEntry(types.ModCmn, nil)

	heav := types.HeAvType{}
	authVec := modelsData.AuthenticationVector

	rand, _ := hex.DecodeString(authVec.Rand)
	copy(heav.Rand[:], rand)

	xres, _ := hex.DecodeString(authVec.Xres)
	copy(heav.XRes[:], xres)

	autn, _ := hex.DecodeString(authVec.Autn)
	copy(heav.Autn[:], autn)

	ck, _ := hex.DecodeString(authVec.CkPrime)
	copy(heav.CK[:], ck)

	ik, _ := hex.DecodeString(authVec.IkPrime)
	copy(heav.IK[:], ik)

	xress, _ := hex.DecodeString(authVec.XresStar)
	copy(heav.XResStar[:], xress)

	kausf, _ := hex.DecodeString(authVec.Kausf)
	copy(heav.Kausf[:], kausf)

	//todo
	//heav.MACA
	//heav.SqnXorAk
	//heav.AK

	rlogger.Trace(types.ModCmn, rlogger.INFO, nil, "HeAv Infos:%s", &heav)

	return heav
}

//amf receive Nsmf msg, trans modelsData to n11Data
func Trans_ModelsToN11_SmContextCreatedDataFormat(modelsData *models.SmContextCreatedData) (n11Data n11msg.SMContextCreatedData) {
	rlogger.FuncEntry(types.ModCmn, nil)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "modelsData(%v)", modelsData)
	if modelsData.PduSessionId != 0 {
		n11Data.PduSessionId = nas.PduSessID(modelsData.PduSessionId)
	}
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "n11Data(%v)", n11Data)
	return n11Data
}

func Trans_ModelsToN11_SmContextCreateDataFormat(modelsData *models.SmContextCreateData) (n11Data n11msg.SmContextCreateData) {
	rlogger.FuncEntry(types.ModCmn, nil)

	if modelsData.Supi != "" {
		n11Data.Supi.Store(modelsData.Supi, types3gpp.IMSIType)
		n11Data.IeFlags.Set(n11msg.Ieid_supi)
	}

	n11Data.PduSessionId = nas.PduSessID(modelsData.PduSessionId)
	n11Data.IeFlags.Set(n11msg.Ieid_pdusessionId)

	err := n11Data.Dnn.StoreWithString(modelsData.Dnn)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail get dnn")
		return
	}
	n11Data.IeFlags.Set(n11msg.Ieid_dnn)

	if modelsData.SNssai != nil {
		n11Data.SNssai.Sst = byte(modelsData.SNssai.Sst)
		if modelsData.SNssai.Sd != "" {
			sd, _ := hex.DecodeString(modelsData.SNssai.Sd[:])
			copy(n11Data.SNssai.Sd[:], sd[0:3])
			n11Data.SNssai.Ind = nasie.SstSd
		} else {
			n11Data.SNssai.Ind = nasie.SstOnly
		}

		n11Data.IeFlags.Set(n11msg.Ieid_sNssai)
	}

	if modelsData.Guami != nil {
		_ = n11Data.Guami.PlmnId.SetString(modelsData.Guami.PlmnId.Mcc + modelsData.Guami.PlmnId.Mnc)
		_ = n11Data.Guami.AmfId.StoreWithString(modelsData.Guami.AmfId)
		n11Data.IeFlags.Set(n11msg.Ieid_guami)
	}

	if modelsData.ServingNetwork != nil {
		_ = n11Data.ServingNetwork.SetString(modelsData.ServingNetwork.Mcc + modelsData.ServingNetwork.Mnc)
		n11Data.IeFlags.Set(n11msg.Ieid_servingNetwork)
	}

	if modelsData.RequestType != "" {
		n11Data.RequestType = Trans_ModelsToN11_RequestType(modelsData.RequestType)
		n11Data.IeFlags.Set(n11msg.Ieid_requestType)
	}

	if modelsData.AnType != "" {
		n11Data.AnType = Trans_ModelsToN11_AnType(modelsData.AnType)
		n11Data.IeFlags.Set(n11msg.Ieid_anType)
	}

	if modelsData.N1SmMsg != nil {
		n11Data.N1SmMsg, _ = hex.DecodeString(modelsData.N1SmMsg.ContentId)
		n11Data.IeFlags.Set(n11msg.Ieid_n1SmMsg)
	}

	if modelsData.RatType != "" {
		n11Data.RatType = Trans_ModelsToN11_RatType(modelsData.RatType)
		n11Data.IeFlags.Set(n11msg.Ieid_ratType)
	}

	if modelsData.SelMode != "" {
		n11Data.SelMode = Trans_ModelsToN11_SelMode(modelsData.SelMode)
		n11Data.IeFlags.Set(n11msg.Ieid_selMode)
	}
	return n11Data
}

//amf receive Nsmf msg, trans modelsData to n11Data
func Trans_ModelsToN11_SmContextUpdatedDataFormat(modelsData *models.UpdateSmContextResponse) (n11Data n11msg.UpdateSMContextResponseData) {
	rlogger.FuncEntry(types.ModCmn, nil)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "modelsData (%v)", modelsData)

	if modelsData.JsonData != nil {
		if modelsData.JsonData.UpCnxState != "" {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "get UpCnxState(%v)", modelsData.JsonData.UpCnxState)
			n11Data.UpCnxState = Trans_ModelsToN11_UpCnxState(modelsData.JsonData.UpCnxState)
			n11Data.IeFlags.Set(n11msg.Ieid_upCnxState)
		}

		if modelsData.BinaryDataN2SmInformation != nil && modelsData.JsonData.N2SmInfoType != "" {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "get n2 msg(%v) ", modelsData.BinaryDataN2SmInformation,
				modelsData.JsonData.N2SmInfoType)
			n11Data.N2SmInfo = string(modelsData.BinaryDataN2SmInformation)
			n11Data.IeFlags.Set(n11msg.Ieid_n2SmInfo)
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "n11Data.N2SmInfo (%v)", n11Data.N2SmInfo)

			n11Data.N2SmInfoType = Trans_ModelsToN11_N2SmInfoType(modelsData.JsonData.N2SmInfoType)
			n11Data.IeFlags.Set(n11msg.Ieid_n2SmInfoType)
		}

		if modelsData.BinaryDataN1SmMessage != nil && modelsData.JsonData.N1SmMsg.ContentId != "" {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "get n1 msg (%v)", modelsData.BinaryDataN1SmMessage,
				modelsData.JsonData.N1SmMsg.ContentId)
			n11Data.N1SmMsg = string(modelsData.BinaryDataN1SmMessage)
			n11Data.IeFlags.Set(n11msg.Ieid_n1SmMsg)
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "n11Data.N1SmMsg (%v)", n11Data.N1SmMsg)
		}
	}
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "n11Data(%v) ", n11Data)

	return n11Data
}

//amf receive Nsmf msg, trans modelsData to n11Data
func Trans_ModelsToN11_N1N2MessageTransferReqDataFormat(data models.N1N2MessageTransferRequest) n11msg.N1N2MessageTransferReqData {
	rlogger.FuncEntry(types.ModCmn, nil)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "data(%v)", data)

	n11Data := n11msg.N1N2MessageTransferReqData{}

	if data.JsonData != nil {

		modelsData := data.JsonData
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "modelsData(%v) ", modelsData)

		if len(data.BinaryDataN1Message) != 0 && modelsData.N1MessageContainer.N1MessageClass != "" {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "get n1 msg(%v) ", data.BinaryDataN1Message,
				modelsData.N1MessageContainer.N1MessageClass)
			n11Data.N1MessageContainer = &n11msg.N1MessageContainerIE{}
			n11Data.N1MessageContainer.N1MessageContent = data.BinaryDataN1Message
			n11Data.N1MessageContainer.N1MsgClass = Trans_ModelsToN11_N1MessageClass(modelsData.N1MessageContainer.N1MessageClass)
			n11Data.IeFlags.Set(n11msg.Ieid_n1MessageContainer)
		}

		if modelsData.PduSessionId != 0 {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "get psi(%d) ", modelsData.PduSessionId)
			n11Data.SessionId = nas.PduSessID(modelsData.PduSessionId)
			n11Data.IeFlags.Set(n11msg.Ieid_pdusessionId)
		}

		if len(data.BinaryDataN2Information) != 0 && modelsData.N2InfoContainer.N2InformationClass != "" &&
			modelsData.N2InfoContainer.SmInfo.N2InfoContent.NgapIeType != "" {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "get n2 info(%v) ", data.BinaryDataN2Information,
				modelsData.N2InfoContainer.N2InformationClass)
			n11Data.N2InfoContainer = &n11msg.N2InfoContainerIE{}
			n11Data.N2InfoContainer.SmInfo = &n11msg.N2SmInformation{}
			n11Data.N2InfoContainer.SmInfo.N2InfoCont = &n11msg.N2InfoContent{}
			n11Data.N2InfoContainer.N2InforClass = Trans_ModelsToN11_N2InformationClass(modelsData.N2InfoContainer.N2InformationClass)
			n11Data.N2InfoContainer.SmInfo.N2InfoCont.NgapData = data.BinaryDataN2Information
			n11Data.N2InfoContainer.SmInfo.N2InfoCont.NgapIeType = Trans_ModelsToN11_NgapIeType(modelsData.N2InfoContainer.SmInfo.N2InfoContent.NgapIeType)
			n11Data.IeFlags.Set(n11msg.Ieid_n2InfoContainer)
			n11Data.N2InfoContainer.IeFlags.Set(n11msg.Ieid_n2SmInfo)
			n11Data.N2InfoContainer.SmInfo.IeFlags.Set(n11msg.Ieid_n2InfoContent)
		}

		if modelsData.Ppi != 0 {
			n11Data.Ppi = int(modelsData.Ppi)
			n11Data.IeFlags.Set(n11msg.Ieid_ppi)
		}
	} else {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "JsonData is nil")
	}

	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "n11Data (%v)", n11Data)
	return n11Data
}

func Trans_ModelsToN11_N1N2MsgTransferRespFormat(modelsData *models.N1N2MessageTransferRspData) (n11Data n11msg.N1N2MessageTransferRspData) {
	rlogger.FuncEntry(types.ModCmn, nil)
	return n11Data
}

func Trans_ModelsToN11_SmContextReleaseDataFormat(mData *models.SmContextReleaseData) (n11Data n11msg.ReleaseSMContextRequestData) {
	rlogger.FuncEntry(types.ModCmn, nil)

	if mData.NgApCause != nil {
		n11Data.NgApCause.Value = types3gpp.CauseValue(mData.NgApCause.Value)
		n11Data.NgApCause.Type = types3gpp.CauseType(mData.NgApCause.Group)
	}

	return n11Data
}

func Trans_ModelsToN11_SmContextUpdateDataReq(mData *models.UpdateSmContextRequest) (n11Data n11msg.UpdateSMContextRequestData) {
	rlogger.FuncEntry(types.ModCmn, nil)

	n11Data.UpCnxState = Trans_ModelsToN11_UpCnxState(mData.JsonData.UpCnxState)
	n11Data.IeFlags.Set(n11msg.Ieid_upCnxState)

	if mData.BinaryDataN1SmMessage != nil {
		n11Data.N1SmMsg = mData.BinaryDataN1SmMessage
		n11Data.IeFlags.Set(n11msg.Ieid_n1SmMsg)
	}

	if mData.BinaryDataN2SmInformation != nil {
		n11Data.N2SmInfo = string(mData.BinaryDataN2SmInformation)
		n11Data.IeFlags.Set(n11msg.Ieid_n2SmInfo)
	}

	n11Data.N2SmInfoType = Trans_ModelsToN11_N2SmInfoType(mData.JsonData.N2SmInfoType)
	n11Data.IeFlags.Set(n11msg.Ieid_n2SmInfoType)

	n11Data.Release = mData.JsonData.Release
	n11Data.IeFlags.Set(n11msg.Ieid_release)

	//rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "modle: %v. n11 update sm context: %v", *(mData.JsonData), n11Data)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "N11 UpCnxState(%s),Release(%v)", n11Data.UpCnxState.String(), n11Data.Release)
	return n11Data
}
