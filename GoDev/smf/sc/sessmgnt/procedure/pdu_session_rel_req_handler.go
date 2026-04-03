package procedure

import (
	"context"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/oam/pm"
	"lite5gc/openapi/models"
	"lite5gc/smf/sc/sessmgnt/smsender"
	"lite5gc/smf/sc/sessmgnt/smutility"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func HandlePduSessRelReqMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

	// ue context
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get ue context")
		return
	}
	ctxt = context.WithValue(ctxt, types.UeContextCK, ueCtxt)

	// pdu session context
	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}
	ctxt = context.WithValue(ctxt, types.SmfPduSessCtxtCK, pduSessCtxt)

	// procedure context
	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionRelReqPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt,
			"prcdctxt is not PduSessionRelReqPrcdCtxt, but:",
			gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return
	}

	td := []interface{}{ctxt, pCtxt}

	// print the status changes
	PrcdState := pduSessCtxt.GetPrcdCtxt().GetCurrentState()
	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, td, "Current, Session Procedure State is (%s).", PrcdState)

	// send N4 session establish request to UPF && get response msg immediately
	msgData := &gctxt.ScN4MsgData{}
	err := smutility.SendMsg2PfcpNode(ctxt, pduSessCtxt, msgData, pfcp.PFCP_Session_Deletion_Request)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to send msg to pfcp node")
		return
	}

}

func HandlePduSessRelUpdateSmCtxtReqSecond(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

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

	// set the next state
	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionRelReqPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil,
			"prcdctxt is not PduSessionRelReqPrcdCtxt, but:",
			gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return
	}
	err := pCtxt.SetNextState(statetype.StatePduSessRelEnd)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to set state")
		return
	}

	modelData := &models.UpdateSmContextResponse{}
	smsender.SendSbiRespMsgUpdateSmCtxtResponse(modelData, pCtxt.SbiMessage)

	//release session ctxt in smf
	//err = smutility.ReleasePduSessCtxt(ctxt, pduSessCtxt.PduSessionId)
	//if err != nil {
	//	rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "fail to release ue ctxt and session ctxt")
	//	return
	//}

	//cancel t3550
	timerMgr, ok := ctxt.Value(types.SmfScTimerMgrCK).(*timermgr.TimerMgr)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to get timer manager")
		return
	}

	timerMgr.CancelTimer(pCtxt.TimerId)
	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "cancel t3592 timer for release command msg")

	if len(ueCtxt.PduSessCtxts) == 0 && ueCtxt.PegIdleAlready == false {
		pm.PegCounter(statistics.IdleUserCounter)
		ueCtxt.PegIdleAlready = true
	}

	//release session ctxt in smf
	ctxt = context.WithValue(ctxt, types.UeContextCK, ueCtxt)
	err = smutility.ReleasePduSessCtxt(ctxt, pCtxt.PduSessId)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to release ue ctxt and session ctxt")
		return
	}
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "ue oriented session release done in SMF")

}
