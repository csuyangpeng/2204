package smsender

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/openapi/models"
	"lite5gc/smf/sc/sessmgnt"
	"lite5gc/smf/smfcontext/gctxt"
	"net"
)

func SendN1N2MsgTransfer4Paging(ctxt context.Context, pduSessCtxt *gctxt.PduSessContext, scInst uint32) error {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)

	var err error
	var msgData n11msg.N1N2MessageTransferReqData

	msgData.SessionId = pduSessCtxt.PduSessionId
	msgData.IeFlags.Set(n11msg.Ieid_pdusessionId)

	// Encode N2 message
	n2Info := &n11msg.N2SmInformation{}
	n2Info.PduSessionId = pduSessCtxt.PduSessionId
	n2Info.IeFlags.Set(n11msg.Ieid_pdusessionId)
	n2Info.Snssai = pduSessCtxt.SNSSAI
	n2Info.IeFlags.Set(n11msg.Ieid_sNssai)

	n2msg := ngapmsg.NewPduSessResSetupReqTransfer()

	sessMgnt, ok := ctxt.Value(types.SessMgntCK).(*sessmgnt.SessMGMT)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to get session mangement")
		return fmt.Errorf("failed to get session mangement")
	}
	n2msg.SetOssCodecCtxt(sessMgnt.GetOssCtxt().GetOssCtxtPtr_m())

	dnnConfig := pduSessCtxt.DnnConfig
	if dnnConfig != nil {
		n2msg.PduSessType = dnnConfig.DefaultPDUSessionType.Convert2NgApType()

		qosflow := &types3gpp.QosFlowSetupReqest{}
		//qosflow.QosFlowInd = uint(configure.GetQosRlue1Qfi())
		qosflow.QosFlowInd = 1
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

		n2msg.PduSessAMBR.Uplink, err = types3gpp.StoreAmbrBitRate(pduSessCtxt.SessionAmbr.UpLink.String())

		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to set Sesson AMBR uplink")
		}
		n2msg.PduSessAMBR.Downlink, err = types3gpp.StoreAmbrBitRate(pduSessCtxt.SessionAmbr.DownLink.String())
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to set Sesson AMBR downlink")
		}
		rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "n2msg.PduSessAMBR", n2msg.PduSessAMBR)

		n2msg.OptFlags.Set(ngapmsg.PSRSRT_PduSessAMBR)

	} else {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "no dnn config in procedure context")
	}

	c, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(pduSessCtxt.SEID))
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to get GetN4Context by seid, %s", err)
		return fmt.Errorf("fail to get GetN4Context by seid, %s", err)
	} else {
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "c.PDRs[0].PDI.LocalFTEID", c.PDRs[0].PDI.LocalFTEID, c)
		if c.PDRs[0].PDI.LocalFTEID != nil {
			n2msg.GtpTunnel.SetTeid(uint32(c.PDRs[0].PDI.LocalFTEID.TEID))
			n2msg.GtpTunnel.SetIpAddr(c.PDRs[0].PDI.LocalFTEID.IPv4Addr)
		} else {
			n2msg.GtpTunnel.SetTeid(0)
			n2msg.GtpTunnel.SetIpAddr(net.IP{})
		}
	}

	n2Info.N2InfoCont = &n11msg.N2InfoContent{}
	n2Info.N2InfoCont.NgapIeType = n11msg.NgapPduResSetupReq
	encodeMsg := n2msg.Encode()
	n2Info.N2InfoCont.NgapData = append(n2Info.N2InfoCont.NgapData, encodeMsg...)
	n2Info.IeFlags.Set(n11msg.Ieid_n2InfoContent)

	//set n2 info msg into msg data
	N2Container := &n11msg.N2InfoContainerIE{N2InforClass: n11msg.SM_N2Info, SmInfo: n2Info}
	N2Container.IeFlags.Set(n11msg.Ieid_n2SmInfo)
	msgData.N2InfoContainer = N2Container
	msgData.IeFlags.Set(n11msg.Ieid_n2InfoContainer)

	//send n1n2 message request to SBI
	//reqData := sbicmn.Trans_N11ToModels_N1N2MsgTransferReqFormat(&msgData)

	requestData := &models.N1N2MessageTransferRequest{}
	requestData.JsonData = sbicmn.Trans_N11ToModels_N1N2MsgTransferReqFormat(&msgData)
	//requestData.BinaryDataN1Message = msgData.N1MessageContainer.N1MessageContent
	requestData.BinaryDataN2Information = msgData.N2InfoContainer.SmInfo.N2InfoCont.NgapData

	SendN1N2MsgTransfer(ctxt, requestData, pduSessCtxt.IMSI.AddIMSIPrefix(), pduSessCtxt.PduSessionId, scInst)

	return nil
}
func SendN1N2MsgTransfer(ctxt context.Context, reqData *models.N1N2MessageTransferRequest, imsi string, psi nas.PduSessID, scId uint32) error {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)

	smMessage := sbicmn.SbiMessage{}
	var n1n2Msg = &sbicmn.SbiPostN1N2MsgTransferMsg{}
	n1n2Msg.Supi = imsi
	n1n2Msg.Psi = psi
	n1n2Msg.ReqData = reqData

	smMessage.MsgData = n1n2Msg
	smMessage.MsgType = sbicmn.N1N2MessageTransferReq
	smMessage.ScInstId = scId

	err := routeragent.SendIpcMessage(ctxt, router.SbiGR, 1, &smMessage) // only one sbi layer
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to send sbi msg to sbi layer. error(%s)", err)
		return err
	}

	return nil
}
