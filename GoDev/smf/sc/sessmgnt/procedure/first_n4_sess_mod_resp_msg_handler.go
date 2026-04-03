package procedure

import (
	"context"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/message/udmdata"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/sc/sessmgnt"
	"lite5gc/smf/sc/sessmgnt/smsender"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
	"net"
)

func Handle1stN4SessModResp(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	// pdu session context
	pCtxt, ok := ctxt.Value(types.PduSessionModPrcdCtxtCK).(*prcdctxt.PduSessionModPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session prcd context")
		return
	}
	ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(pCtxt.IMSI.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "failed  to find ue context by imsi %s", pCtxt.IMSI.String())
		return
	}

	var modCmdMsg nasmsg.PDUSessionModifyCommandMsg
	modCmdMsg.EPD = nas.Epd5gsSessMgntMsg
	modCmdMsg.MsgHeader.PrcdTransactionID = pCtxt.Pti
	modCmdMsg.MsgHeader.PduSessionID = pCtxt.PduSessId
	modCmdMsg.MsgHeader.MessageType = nas.PduSessionModCommand
	modCmdBytes, err := modCmdMsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "fail to encode session mod cmd msg, err:(%s)", err)
		return
	}

	//n2
	n2msg := ngapmsg.NewPduSessResSetupReqTransfer()
	sessMgnt, ok := ctxt.Value(types.SessMgntCK).(*sessmgnt.SessMGMT)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to get session mangement")
		return
	}
	n2msg.SetOssCodecCtxt(sessMgnt.GetOssCtxt().GetOssCtxtPtr_m())
	sessSubsData := udmdata.FindSessMgntSubsDataWithSnssai(pCtxt.SNSSAI.String(), ueCtxt.SessMgntSubsDataMap)
	if sessSubsData == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "failed to get sessionSubMsg")
		return
	}
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, ueCtxt.SessMgntSubsDataMap)
	dnnConfig := sessSubsData.DnnConfigs[pCtxt.DNN.String()]
	if dnnConfig != nil {
		n2msg.PduSessType = dnnConfig.DefaultPDUSessionType.Convert2NgApType()

		c, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(pCtxt.Seid))
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to get GetN4Context by seid, %s", err)
			return
		} else {
			rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "c.PDRs[0].PDI.LocalFTEID", c.PDRs[0].PDI.LocalFTEID, c)
			if c.PDRs[0].PDI.LocalFTEID != nil {
				n2msg.GtpTunnel.SetTeid(uint32(c.PDRs[0].PDI.LocalFTEID.TEID))
				n2msg.GtpTunnel.SetIpAddr(c.PDRs[0].PDI.LocalFTEID.IPv4Addr)
			} else {
				n2msg.GtpTunnel.SetTeid(0)
				n2msg.GtpTunnel.SetIpAddr(net.IP{})
			}
			rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, []interface{}{ctxt, pCtxt, c}, "gtp tunnel info : %s", n2msg.GtpTunnel.String())
		}

		qosflow := &types3gpp.QosFlowSetupReqest{}
		qosflow.QosFlowInd = 1
		//qosflow.QosFlowInd = uint(configure.GetQosRlue1Qfi())
		qosflow.QosLevelParam.Arp.PriorityLevel = dnnConfig.QosProf.Arp.PriorityLevel
		qosflow.QosLevelParam.Arp.PreemptCap = dnnConfig.QosProf.Arp.PreemptCap
		qosflow.QosLevelParam.Arp.PreemptVuln = dnnConfig.QosProf.Arp.PreemptVuln
		qosflow.QosLevelParam.QosChats.IsDynamic = false
		// before:
		// qosflow.QosLevelParam.QosChats.NonDynamic5qi.FiveQI = dnnConfig.QosProf.NonDynamic5qi.FiveQI
		// after:
		qosflow.QosLevelParam.QosChats.NonDynamic5qi.FiveQI = uint(dnnConfig.QosProf.QI5)
		//qosflow.QosLevelParam.QosChats.NonDynamic5qi.OptFlags.Set(types3gpp.ND5QI_AverageWind)
		//qosflow.QosLevelParam.QosChats.NonDynamic5qi.AverageWindow = dnnConfig.QosProf.NonDynamic5qi.AverageWindow

		n2msg.QosFlowSetupReqList = []*types3gpp.QosFlowSetupReqest{qosflow}

		n2msg.PduSessAMBR.Uplink, err = types3gpp.StoreAmbrBitRate(dnnConfig.SessionAmbr.UpLink.String())

		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to set Sesson AMBR uplink")
		}
		n2msg.PduSessAMBR.Downlink, err = types3gpp.StoreAmbrBitRate(dnnConfig.SessionAmbr.DownLink.String())
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to set Sesson AMBR downlink")
		}
		rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "n2msg.PduSessAMBR", n2msg.PduSessAMBR)

		n2msg.OptFlags.Set(ngapmsg.PSRSRT_PduSessAMBR)

	} else {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "no dnn config in procedure context")
	}
	var msgData n11msg.UpdateSMContextResponseData

	msgData.N1SmMsg = string(modCmdBytes)
	msgData.IeFlags.Set(n11msg.Ieid_n1SmMsg)
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "msgData.N1SmMsg", msgData.N1SmMsg)

	msgData.N2SmInfo = string(n2msg.Encode())
	msgData.IeFlags.Set(n11msg.Ieid_n2SmInfo)

	msgData.N2SmInfoType = n11msg.PduResModReq
	msgData.IeFlags.Set(n11msg.Ieid_n2SmInfoType)

	msgData.UpCnxState = n11msg.ACTIVATING
	msgData.IeFlags.Set(n11msg.Ieid_upCnxState)

	modelsData := sbicmn.Trans_N11ToModels_SmContextUpdatedDataReq(&msgData)

	err = smsender.SendSbiRespMsgUpdateSmCtxtResponseWithN1N2Info(modelsData, pCtxt.SbiMessage)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to send sbi msg")
		return
	}

	err = pCtxt.SetNextState(statetype.StatePduSessModWf2ndUpdateSMCtxtReq)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to set state")
		return
	}
}
