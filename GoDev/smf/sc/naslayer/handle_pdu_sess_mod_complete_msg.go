package naslayer

import (
	"context"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/sessmgnt/smutility"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func (p *NasMgr) HandleSmPduSessModComplete(ctx context.Context, header *nas.SmNasMessageHeader) {
	rlogger.FuncEntry(types.ModuleSmfNas, nil)
	if header == nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "input para is nil")
		return
	}
	// get ue context
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "failed to get ue context")
		return
	}

	//decode the PdqSessModComplete message
	p.PdqSessModComplete.SmHeader.PduSessionID = header.PduSessionID
	p.PdqSessModComplete.SmHeader.PrcdTransactionID = header.PrcdTransactionID
	p.PdqSessModComplete.SmHeader.MessageType = nas.PduSessionModComplete

	//get the pdu session ctxt
	pduSessCtxt, ok := ctx.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to get pdu session context")
		return
	}

	//create a pdu session modification prcd ctxt
	pCtxt := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionModPrcdCtxt)
	pCtxt.IMSI = ueCtxt.IMSI
	pCtxt.DNN = pduSessCtxt.DNN
	pCtxt.SNSSAI = pduSessCtxt.SNSSAI
	pCtxt.UEIP = pduSessCtxt.UEIP
	sbimsg, _ := ctx.Value(types.SmfSbiHandlerMsgCK).(*sbicmn.SbiHandlerMessage)
	pCtxt.SbiMessage = sbimsg

	msgData := &gctxt.ScN4MsgData{}
	msgData.IMSI = ueCtxt.IMSI
	//msgData.DNN = pduSessCtxt.DNN
	msgData.SNSSAI = pduSessCtxt.SNSSAI
	//msgData.UEIP = pduSessCtxt.UEIP
	msgData.PFCPParameters.PDRs = make([]*gctxt.PacketDetectionRule, 1)
	msgData.PFCPParameters.PDRs[0] = &(gctxt.PacketDetectionRule{})
	msgData.PFCPParameters.IeFlags.Set(pfcp.IE_Create_PDR)

	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.RuleID = 2
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	msgData.PFCPParameters.PDRs[0].FarI.Far.Action = gctxt.FORW
	msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Apply_Action)

	rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, pduSessCtxt, "ready to send pfcp msg")

	msgData.PFCPParameters.PDRs[0].Tag = gctxt.UpdatePDRs
	err := smutility.SendMsg2PfcpNode(ctx, pduSessCtxt, msgData, pfcp.PFCP_Session_Modification_Request)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to send msg to pfcp node")
		return
	}
}
