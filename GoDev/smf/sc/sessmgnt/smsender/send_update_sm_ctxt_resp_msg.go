package smsender

import (
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func SendUpdateSmCtxtResponseMsg4AnRelPrcd(prcdCtxt *prcdctxt.AnRelSerReqPrcdCtxt) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	if prcdCtxt == nil{
		rlogger.Trace(types.ModuleSmfSM,rlogger.ERROR,nil,"input para is nil")
		return
	}
	var sbimsg *sbicmn.SbiHandlerMessage
	sbimsg = prcdCtxt.SbiMessage

	n11MsgData := &n11msg.UpdateSMContextResponseData{}
	n11MsgData.UpCnxState = n11msg.DEACTIVATED
	n11MsgData.IeFlags.Set(n11msg.Ieid_upCnxState)

	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "Update Response:%c", n11MsgData)

	msgData := sbicmn.Trans_N11ToModels_SmContextUpdatedDataReq(n11MsgData)
	SendSbiRespMsgUpdateSmCtxtResponse(msgData, sbimsg)
}

func SendUpdateSmCtxtResponseMsg4SessModPrcd(prcdCtxt *prcdctxt.PduSessionModPrcdCtxt) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	if prcdCtxt == nil{
		rlogger.Trace(types.ModuleSmfSM,rlogger.ERROR,nil,"input para is nil")
		return
	}
	var sbimsg *sbicmn.SbiHandlerMessage
	sbimsg = prcdCtxt.SbiMessage

	n11MsgData := &n11msg.UpdateSMContextResponseData{}
	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "Update Response:%c", n11MsgData)

	msgData := sbicmn.Trans_N11ToModels_SmContextUpdatedDataReq(n11MsgData)
	SendSbiRespMsgUpdateSmCtxtResponse(msgData, sbimsg)
}
