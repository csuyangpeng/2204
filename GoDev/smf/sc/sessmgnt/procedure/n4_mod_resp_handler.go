package procedure

import (
	"context"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/openapi/models"
	"lite5gc/smf/sc/sessmgnt/smsender"
	"lite5gc/smf/sc/sessmgnt/smutility"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func HandleN4SessModResp(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	// pdu session context
	pCtxt, ok := ctxt.Value(types.PduSessionEstbPrcdCtxtCK).(*prcdctxt.PduSessionEstbPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	respData := &models.UpdateSmContextResponse{}

	if pCtxt.SbiMessage == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "invalid SbiMessage in procedure context")
		return
	}

	sbiMsg := pCtxt.SbiMessage

	smsender.SendSbiRespMsgUpdateSmCtxtResponse(respData, sbiMsg)

	//finished the pdu session establish procedure
	PduSessEstbFinished(pCtxt)
}

func HandleN4SessModResp4SerReq(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)

	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	// procedure context
	prcdCtxt, ok := ctxt.Value(types.PduSessionAnRelSerReqCtxtCK).(*prcdctxt.AnRelSerReqPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	var sbimsg *sbicmn.SbiHandlerMessage

	n11MsgData := &n11msg.UpdateSMContextResponseData{}

	sbimsg = prcdCtxt.SbiMessage

	n11MsgData.UpCnxState = n11msg.ACTIVATED
	n11MsgData.IeFlags.Set(n11msg.Ieid_upCnxState)
	respData := &smutility.UpdateSmCtxtRespData{}
	respData.Seid = pduSessCtxt.SEID
	respData.Dnn = pduSessCtxt.DNN.String()
	respData.Snssai = pduSessCtxt.SNSSAI.String()
	respData.UpCnxState = n11msg.ACTIVATED
	respData.Imsi = pduSessCtxt.IMSI
	err, encodN2SmInfo := smutility.GenPduSessResSetupReqTransferMsg(ctxt, respData)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to construct PduSessResSetupReq msg")
		return
	}

	n11MsgData.N2SmInfo = string(encodN2SmInfo)
	// GenerateN2InUpdateSmCtxt == ngapsmf.NewPduSessResSetupReqTransfer().encode()
	// == n1n2Req n2Info SmInfo N2InfoCont NgapData
	n11MsgData.IeFlags.Set(n11msg.Ieid_n2SmInfo)

	n11MsgData.N2SmInfoType = n11msg.PduResSetupReq
	n11MsgData.IeFlags.Set(n11msg.Ieid_n2SmInfoType)

	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "Update Response:%c", n11MsgData)

	msgData := sbicmn.Trans_N11ToModels_SmContextUpdatedDataReq(n11MsgData)
	smsender.SendSbiRespMsgUpdateSmCtxtResponse(msgData, sbimsg)

	prcdCtxt.IsWfSecN4ModifyResp = true
	err = prcdCtxt.SetNextState(statetype.StatePduSessSerReqWfUpdateSmCtxtReqSec)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to set state")
		return
	}
}

func HandleN4SessModRespSec4SerReq(ctxt context.Context) {
	// pdu session context
	pCtxt, ok := ctxt.Value(types.PduSessionAnRelSerReqCtxtCK).(*prcdctxt.AnRelSerReqPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	respData := &models.UpdateSmContextResponse{}

	if pCtxt.SbiMessage == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "invalid SbiMessage in procedure context")
		return
	}

	sbiMsg := pCtxt.SbiMessage

	smsender.SendSbiRespMsgUpdateSmCtxtResponse(respData, sbiMsg)

	//finished the service request procedure
	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	pduSessCtxt.SetPrcdCtxt(nil)

}
