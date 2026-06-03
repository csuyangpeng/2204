package procedure

import (
	"context"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/udmdata"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/idmgrsmf"
	"lite5gc/smf/sc/sessmgnt"
	"lite5gc/smf/sc/sessmgnt/smsender"
	"lite5gc/smf/sc/sessmgnt/smutility"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
	"lite5gc/smf/ueippool"
)

func HandlePduSessEstbReqMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

	// ue context
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get ue context")
		return
	}

	// pdu session context
	psi, ok := ctxt.Value(types.PduSessIdCK).(nas.PduSessID)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to get pdu session context")
		return
	}
	pduSessCtxt, ok := ueCtxt.PduSessCtxts[psi]
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to get pdu sess ctxt with psi(%d)", psi)
		return
	}

	// pdu sess estb procedure context
	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt,
			"the prcdctxt is not PduSessionEstbPrcdCtxt,but:",
			gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return
	}

	// Check if need to fetch data from UDM
	// If Session Management Subscription data for corresponding SUPI, DNN and S-NSSAI is not available,
	// then SMF retrieves the Session Management Subscription data
	// using Nudm_SDM_Get (SUPI, Session Management Subscription data, DNN, S-NSSAI)
	needSdmGetData := true
	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "debug smsd in ue context:%v", ueCtxt.SessMgntSubsDataMap)
	for _, v := range ueCtxt.SessMgntSubsDataMap {
		// snssai
		if pCtxt.IeFlags.Test(n11msg.Ieid_sNssai) {
			if v.SingleNssai.IsValid() && pCtxt.SNSSAI.IsValid() && (v.SingleNssai.String() == pCtxt.SNSSAI.String()) {
				rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "debug Snssai(%s), ueCtxt(%s)", pCtxt.SNSSAI.String(), v.SingleNssai.String())
				needSdmGetData = false
				break
			}
		}
		// dnn
		if pCtxt.IeFlags.Test(n11msg.Ieid_dnn) {
			_, ok = v.DnnConfigs[pCtxt.DNN.String()]
			if ok {
				rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "debug dnn(%s), uectxt(%v)", pCtxt.DNN.String(), v.DnnConfigs)
				needSdmGetData = false
				break
			}
		}
	}

	// ue context
	sm, ok := ctxt.Value(types.SessMgntCK).(*sessmgnt.SessMGMT)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get Session Management")
		return
	}
	ctxt = context.WithValue(ctxt, types.SmfPduSessCtxtCK, pduSessCtxt)
	// Get SMD data from udm
	if needSdmGetData == true {
		//  smf select data
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "send SBI GetSmData service request")
		smMessage := sbicmn.SbiMessage{}
		var getsm = &sbicmn.SbiGetSmDataMsg{}
		getsm.Supi = ueCtxt.IMSI.AddIMSIPrefix()
		getsm.Psi = psi
		smMessage.MsgData = getsm
		smMessage.MsgType = sbicmn.GetSmDataMsgRequest
		smMessage.ScInstId = sm.GetScId()

		err := routeragent.SendIpcMessage(ctxt, router.SbiGR, 1, &smMessage) // only one sbi layer
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to send sbi msg to udm. error(%s)", err)
			return
		}
		err = pCtxt.SetNextState(statetype.StatePduSessEstbWfSmGet)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt.IMSI, "fail to set state in sess estb proc")
		}
	} else {
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "don't need to GetSmData")
		HandlePduSessEstbUdmGetRespMsg(ctxt)
	}
	return
}

func HandlePduSessEstbUdmGetRespMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	var cause = nas.SuccessNoReason

	// ue context
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get ue context")
		return
	}

	// pdu session context
	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt, "failed to get pdu session context")
		return
	}

	// pdu sess estb procedure context
	if pduSessCtxt.GetPrcdCtxt() != nil {
		pCtxt := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)

		var smsd udmdata.SessMgntSubscripitonData
		if pCtxt.IeFlags.Test(n11msg.Ieid_sNssai) && pCtxt.IeFlags.Test(n11msg.Ieid_dnn) {
			rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, pduSessCtxt, "ueCtxt.SessMgntSubsDataMap(%v)", ueCtxt.SessMgntSubsDataMap)
			for _, v := range ueCtxt.SessMgntSubsDataMap {
				if v.SingleNssai.String() == pCtxt.SNSSAI.String() {
					smsd = v
				}
			}
			rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, pduSessCtxt, "smsd.DnnConfigs(%v)", smsd.DnnConfigs)
			dnnConfig, ok := smsd.DnnConfigs[pCtxt.DNN.String()]
			if ok != true {
				rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt, "failed to find Dnn Configuration with dnn(%s)", pCtxt.DNN.String())
				return
			}

			pCtxt.SessionAmbr = dnnConfig.SessionAmbr
		}

		pCtxt.SessMgntSubsData = &smsd

		// should check dnn string
		//pCtxt.SessionAmbr = ueCtxt.SessMgntSubsDataMap[pCtxt.Snssai.String()].DnnConfigs[pCtxt.DNN.String()].SessionAmbr

		//store the dnn configuration data
		if pCtxt.SessMgntSubsData != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "pCtxt.DNN.String()", pCtxt.DNN.String(), pCtxt.SessMgntSubsData.DnnConfigs)
			pCtxt.DnnConfig = pCtxt.SessMgntSubsData.DnnConfigs[pCtxt.DNN.String()]
		}
		// 每个S-NSSAI的订阅信息可能会包含一个默认DNN和多个DNN
		// 当UE发起PDU Session Establishment Request时没有提供S-NSSAI的DNN，
		// 那么, AMF选择默认DNN或本地配置DNN
		// 如果UE携带的DNN不被网络支持，且AMF也没能通过查询NRF选择到一个合适SMF，
		// 如果UE携带的DNN不被网络支持，且AMF也没能通过查询NRF选择到一个合适SMF，
		// 则AMF拒绝请求，原因值“DNN is not supported”
		if pCtxt.IeFlags.Test(n11msg.Ieid_selMode) {
			//SelMode DnnSelectionMode C
			//29502  6.1.6.3.10
			switch pCtxt.SelMode {
			case n11msg.VERIFIED:
				// UE or network provided DNN, subscription verified
				if pCtxt.DnnConfig != nil {
					pCtxt.SelectDnn = pCtxt.DNN
				} else {
					rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt, "provided dnn not support")
					//construct reject msg cause value
					//return types.Missing_or_unknown_DNN
				}
			case n11msg.UE_DNN_NOT_VERIFIED:
				// UE provided DNN, subscription not verified
				// if OPTIONAL IE : DNN is exist
				if pCtxt.IeFlags.Test(n11msg.Ieid_dnn) {
					pCtxt.SelectDnn = pCtxt.DNN
				}
			case n11msg.NW_DNN_NOT_VERIFIED:
				// Network provided DNN,subscription not verified
				// if OPTIONAL IE : DNN is exist
				if pCtxt.IeFlags.Test(n11msg.Ieid_dnn) {
					pCtxt.SelectDnn = pCtxt.DNN
				}
			}
		}
		pCtxt.SelectDnn = pCtxt.DNN // TODO temp code for mate20x IOT

		//check session ssc mode
		ssc := pCtxt.SscMode
		switch ssc {
		case nas.SSCMode1:
		case nas.SSCMode2, nas.SSCMode3, nas.AsSSCMode1, nas.AsSSCMode2, nas.AsSSCMode3:
			cause = nas.NotSupportedSSCMode
			rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "NotSupportedSSCMode")
			goto HandleCause
		default:
			cause = nas.NotSupportedSSCMode
			rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "default NotSupportedSSCMode")
			goto HandleCause
		}

		//check session type
		if pCtxt.SessionReqIeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_SessionType) {
			switch pCtxt.PduSessType {
			case types3gpp.Ipv4:
			case types3gpp.Ipv6, types3gpp.Ipv4v6, types3gpp.Unstructured, types3gpp.Ethernet:
				cause = nas.PDUSessionTypeIPv4OnlyAllowed
				rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "PDUSessionTypeIPv4OnlyAllowed")
				goto HandleCause
			default:
				cause = nas.UnknownPDUSessionType
				rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "UnknownPDUSessionType")
				goto HandleCause
			}
		}

		cause = handlePduSessEstbReq(ctxt, ueCtxt, pduSessCtxt)

	HandleCause:
		smsender.ProcessPduSessEstbReqCause(ctxt, cause)
		return
	}
	return
}

func handlePduSessEstbReq(ctxt context.Context,
	ueCtxt *gctxt.UeContext,
	pduSessCtxt *gctxt.PduSessContext) nas.Sm5gCause {
	td := []interface{}{ueCtxt, pduSessCtxt}
	rlogger.FuncEntry(types.ModuleSmfSM, td)

	var cause = nas.SuccessNoReason

	// print the status changes
	PrcdState := pduSessCtxt.GetPrcdCtxt().GetCurrentState()
	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, td, "Current, Session Procedure State is (%s).", PrcdState)

	// get the procedure context
	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td,
			"prcdctxt is not PduSessionEstbPrcdCtxt,but:",
			gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return nas.OtherValue
	}

	// Step 1: send SMContextCreatedData (response) to AMF
	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, td, "pctxt, sbimessage: %v", pCtxt.SbiMessage.ResponseChan)
	err := smsender.SendSbiRespMsgSMContextCreatedData(pduSessCtxt, pCtxt.SbiMessage)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "failed to send SMContextCreatedData to Sbi Layer")
		return nas.OtherValue
	}

	// Step 2: send N4 session establish request to UPF && get response msg immediately
	//save in pdu session ctxt directly for PDR FAR
	// get the session subscriber data from UDM
	dnnConfig := pCtxt.DnnConfig
	if dnnConfig == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "no dnn configuration.")
		cause = nas.MissingOrUnknownDNNInASlice
	} else {
		if pCtxt.DnnConfig.StaticIpAddress == nil {
			ip, err := ueippool.AllocateByDNN(ueippool.DnnNameKey(pCtxt.DNN.String()))
			if err != nil {
				rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "fail to get allocator by dnn: %s", pCtxt.DNN.String())
				return nas.OtherValue
			}
			pCtxt.UEIP = ip
			rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, td, "allocat ip:%s by dnn:%s for ue:%s",
				pCtxt.UEIP, pCtxt.DNN.String(), ueCtxt.IMSI.String())
		} else {
			pCtxt.UEIP = dnnConfig.StaticIpAddress

		}

		if pCtxt.SessionReqIeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_SscMode) {
			if pCtxt.SscMode != dnnConfig.DefaultSSCMode {
				switch dnnConfig.DefaultSSCMode {
				case nas.SSCMode1:
				case nas.SSCMode2, nas.SSCMode3, nas.AsSSCMode1, nas.AsSSCMode2, nas.AsSSCMode3:
					return nas.NotSupportedSSCMode
				default:
					return nas.NotSupportedSSCMode
				}

			}
		}
		// allocated seid
		pduSessCtxt.SEID, err = idmgrsmf.GetSEID()
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "fail to alloc SEID")
			return nas.OtherValue
		}
		//save in pdu session table
		err := gctxt.AddIndexSessContext(gctxt.SeidKey(pduSessCtxt.SEID), pduSessCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "failed to add pdu sess in pduSessTable with SEDI(%d)", pduSessCtxt.SEID)
		} else {
			rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, td, "add pdu sess in pduSessTable with SEDI(%d)", pduSessCtxt.SEID)
		}
		pCtxt.Seid = pduSessCtxt.SEID

		// UPInactiveTimer
		pduSessCtxt.DNN = pCtxt.DNN
		pduSessCtxt.SNSSAI = pCtxt.SNSSAI
		pduSessCtxt.UEIP = pCtxt.UEIP
		//pduSessCtxt.UPInactiveTimer = configure.GetUserPlaneInactTimer()
		//ask upf to return teid

		msgData := &gctxt.ScN4MsgData{}
		//TODO Diretcty save into pdu session context, since they are from config manager
		msgData.PFCPParameters = configure.GetPFCPParams()
		msgData.PFCPParameters.PDRs[0].PDI.IeFlags.Set(pfcp.IE_F_TEID) //tell node to send upf-teid in pfcp-req msg
		msgData.PFCPParameters.PDRs[1].PDI.UEIPAddress = pduSessCtxt.UEIP
		msgData.PFCPParameters.PDRs[1].PDI.IeFlags.Set(pfcp.IE_UE_IP_Address)

		err = smutility.SendMsg2PfcpNode(ctxt, pduSessCtxt, msgData, pfcp.PFCP_Session_Establishment_Request)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "fail to send msg to pfcp node")
			return nas.OtherValue
		}

		// set the next state
		err = pCtxt.SetNextState(statetype.StatePduSessEstbWfN4SessEstbResp)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt.IMSI, "fail to set state in sess estb proc")
			return nas.OtherValue
		}

	}
	return cause
}
