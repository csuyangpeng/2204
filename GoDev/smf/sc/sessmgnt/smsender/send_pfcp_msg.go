package smsender

import (
	"context"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/sc/sessmgnt/smutility"
	"lite5gc/smf/smfcontext/gctxt"
)

func SendPduResModifyReq2Upf(ctxt context.Context,
	pduSessCtxt *gctxt.PduSessContext,
	id pfcp.PFCPMSG,
	transfer *ngapmsg.PduSessResSetupRespTransfer) {

	rlogger.FuncEntry(types.ModuleSmfSM, nil)

	msgData := &gctxt.ScN4MsgData{}

	// new params
	msgData.PFCPParameters.PDRs = make([]*gctxt.PacketDetectionRule, 1)

	msgData.PFCPParameters.PDRs[0] = &(gctxt.PacketDetectionRule{})
	msgData.PFCPParameters.IeFlags.Set(pfcp.IE_Create_PDR)

	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.DstInterface = gctxt.DestinationInterface_Access
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Destination_Interface)

	msgData.PFCPParameters.PDRs[0].FarI.Far.RuleID = 2
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.Action = gctxt.FORW
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Apply_Action)

	msgData.PFCPParameters.PDRs[0].FarI.Far.OuterHeaderCreation.TEID = types3gpp.Teid(transfer.GtpTunnel.GetTeid())
	msgData.PFCPParameters.PDRs[0].FarI.Far.OuterHeaderCreation.IPv4Addr = transfer.GtpTunnel.GetIpAddr()
	msgData.PFCPParameters.PDRs[0].FarI.Far.OuterHeaderCreation.Description = pfcp.IEOuterHCreation_GTPU_UDP_IPv4
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Outer_Header_Creation)

	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "pduSessCtxt.PFCPParameters.PDRs", msgData.PFCPParameters.PDRs)

	err := smutility.SendMsg2PfcpNode(ctxt, pduSessCtxt, msgData, id)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to send msg to pfcp node")
		return
	}
}

func SendPfcpSessDelReqMsg(ctxt context.Context,
	pduSessCtxt *gctxt.PduSessContext,
	id pfcp.PFCPMSG) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)

	msgData := &gctxt.ScN4MsgData{}

	err := smutility.SendMsg2PfcpNode(ctxt, pduSessCtxt, msgData, id)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to send msg to pfcp node")
		return
	}
}

func DeactivePduSession(ctxt context.Context, pduSessCtxt *gctxt.PduSessContext, pfcpMsgType pfcp.PFCPMSG) {
	// for an release procedure
	msgData := &gctxt.ScN4MsgData{}

	//pduSessCtxt.UpCnxState = n11msg.DEACTIVATED
	//pduSessCtxt.DirectRespFlag = false

	msgData.PFCPParameters.PDRs = make([]*gctxt.PacketDetectionRule, 1)
	msgData.PFCPParameters.PDRs[0] = &(gctxt.PacketDetectionRule{})
	msgData.PFCPParameters.IeFlags.Set(pfcp.IE_Create_PDR)

	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.DstInterface = gctxt.DestinationInterface_Access
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Destination_Interface)

	msgData.PFCPParameters.PDRs[0].FarI.Far.RuleID = 2
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.Action = gctxt.NOCPBUFF
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Apply_Action)

	//msgData.PFCPParameters.PDRs[0].FarI.Far.BAR = configure.SmfConf.N4Conf.PFCPParameters.PDRs[1].FarI.Far.BAR
	//msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Create_BAR)

	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "ready to send pfcp msg")

	msgData.PFCPParameters.PDRs[0].Tag = gctxt.UpdatePDRs
	err := smutility.SendMsg2PfcpNode(ctxt, pduSessCtxt, msgData, pfcpMsgType)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to send msg to pfcp node")
		return
	}
	return
}

func activatePduSession(ctxt context.Context, pduSessCtxt *gctxt.PduSessContext) {

	//pduSessCtxt.UpCnxState = n11msg.ACTIVATED
	//pduSessCtxt.DirectRespFlag = true
	msgData := &gctxt.ScN4MsgData{}

	msgData.PFCPParameters.PDRs = make([]*gctxt.PacketDetectionRule, 1)
	msgData.PFCPParameters.PDRs[0] = &(gctxt.PacketDetectionRule{})
	msgData.PFCPParameters.IeFlags.Set(pfcp.IE_Create_PDR)

	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.RuleID = 2
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.Action = gctxt.FORW
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Apply_Action)

	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "send pfcp msg")

	msgData.PFCPParameters.PDRs[0].Tag = gctxt.UpdatePDRs
	err := smutility.SendMsg2PfcpNode(ctxt, pduSessCtxt, msgData, pfcp.PFCP_Session_Modification_Request)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to send msg to pfcp node")
		return
	}
	return
}

func SendActivatePfcpSessModRequest(ctxt context.Context, pduSessCtxt *gctxt.PduSessContext) {

	//pduSessCtxt.UpCnxState = n11msg.ACTIVATED
	//pduSessCtxt.DirectRespFlag = true
	msgData := &gctxt.ScN4MsgData{}

	msgData.PFCPParameters.PDRs = make([]*gctxt.PacketDetectionRule, 1)
	msgData.PFCPParameters.PDRs[0] = &(gctxt.PacketDetectionRule{})
	msgData.PFCPParameters.IeFlags.Set(pfcp.IE_Create_PDR)

	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.RuleID = 2
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.Action = gctxt.FORW
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Apply_Action)

	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt, "send pfcp msg")

	msgData.PFCPParameters.PDRs[0].Tag = gctxt.UpdatePDRs
	err := smutility.SendMsg2PfcpNode(ctxt, pduSessCtxt, msgData, pfcp.PFCP_Session_Modification_Request)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to send msg to pfcp node")
		return
	}
	return
}
