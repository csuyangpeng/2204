package smsender

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func SendPduSessRelAcceptMsg(ctxt context.Context, prcdCtxt *prcdctxt.PduSessionRelReqPrcdCtxt) error {
	td := []interface{}{ctxt, prcdCtxt}
	rlogger.FuncEntry(types.ModuleSmfSM, td)

	// Encode NAS Msg: pdu session establishment accept msg
	var pduSessRelAcceptMsg nasmsg.PduSessionReleaseCommandMsg
	pduSessRelAcceptMsg.EPD = nas.Epd5gsSessMgntMsg
	pduSessRelAcceptMsg.MsgHeader.PrcdTransactionID = prcdCtxt.Pti
	pduSessRelAcceptMsg.MsgHeader.PduSessionID = prcdCtxt.PduSessId
	pduSessRelAcceptMsg.MsgHeader.MessageType = nas.PduSessionRelCommand
	if prcdCtxt.IsCauseExist {
		pduSessRelAcceptMsg.SMCause = prcdCtxt.Cause
	} else {
		pduSessRelAcceptMsg.SMCause = nas.RegularDeactivation
	}
	ctxt = context.WithValue(ctxt, types.SessrelAcptCK, pduSessRelAcceptMsg)

	//Encode N11 message: n1n2message transfer request
	var msgData n11msg.UpdateSMContextResponseData
	encodeMsg, err := pduSessRelAcceptMsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "fail to encode PduSessionReleaseCommandMsg")
		return fmt.Errorf("fail to encode PduSessionReleaseCommandMsg")
	}
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, td, "encodeMsg put in UpdateSmCtxtResp:", encodeMsg)
	msgData.N1SmMsg = string(encodeMsg)
	msgData.IeFlags.Set(n11msg.Ieid_n1SmMsg)

	//send n11 msg to amf
	modeldata := sbicmn.Trans_N11ToModels_SmContextUpdatedDataReq(&msgData)

	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, td, "pctxt, sbimessage: %v", prcdCtxt.SbiMessage.ResponseChan)
	SendSbiRespMsgUpdateSmCtxtResponse(modeldata, prcdCtxt.SbiMessage)

	return nil
}
