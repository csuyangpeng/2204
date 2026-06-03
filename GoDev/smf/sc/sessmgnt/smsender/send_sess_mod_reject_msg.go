package smsender

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/types"
	"lite5gc/smf/idmgrsmf"
)

func SendPduSessModRejectMsg(ctxt context.Context, cause nas.Sm5gCause) error {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

	//get the pdu session ctxt
	//pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	//if !ok {
	//	rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
	//	return fmt.Errorf("failed to get pdu session context")
	//}

	//get the sm header
	header, ok := ctxt.Value(types.SmHeaderCK).(*nas.SmNasMessageHeader)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "no sm header found")
		return fmt.Errorf("no sm header found")
	}

	var rejectMsg nasmsg.PduSessionModifyRejectMsg
	rejectMsg.MsgHeader = *header
	rejectMsg.MsgHeader.MessageType = nas.PduSessionRelReject
	rejectMsg.SMCause = cause
	rejectBytes, err := rejectMsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "fail to encode session reject msg, err:(%s)", err)
		return fmt.Errorf("fail to encode session reject msg, err:(%s)", err)
	}

	var msgData n11msg.UpdateSMContextRequestData
	msgData.N1SmMsg = rejectBytes
	msgData.IeFlags.Set(n11msg.Ieid_n1SmMsg)
	msgData.Release = true
	//msgData.Cause =

	//send n11 msg to amf
	var msg n11msg.N11Msg
	//msg.AmfSmCtxtId = pduSessCtxt.AmfSmCtxtId
	//msg.SmfSmCtxtId = pduSessCtxt.SmfSmCtxtId
	msg.MsgType = n11msg.UpdateSmCtxtResp
	msg.MsgData = msgData

	amfScId := idmgrsmf.RetrieveScId(msg.AmfSmCtxtId)
	err = routeragent.SendIpcMessage(ctxt, router.ScGR, amfScId, msg)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to send ipv msg")
		return err
	}

	return nil
}
