package procedure

import (
	"context"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func HandleN1N2MsgTransferRespMsg4PduSessEstb(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)

	// pdu session context
	pduSessCtxt, ok := ctxt.Value(types.SmfPduSessCtxtCK).(*gctxt.PduSessContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	// pdu sess estb procedure context
	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt.IMSI,
			"prcdctxt is not PduSessionEstbPrcdCtxt,but:", gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return
	}

	// get msgData payload : N1N2MessageTransferRspData
	if pCtxt.Cause == n11msg.N1_N2_TRANSFER_INITIATED {
		rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt.IMSI, "smf received N1N2MessageTransferRspData, "+
			"and AMF has initiated the N1/N2 message transfer towards the UE and/or the AN.")
	}

	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, pduSessCtxt.IMSI,
		"receive n1n2 msg transfer response msg. cause(%d)", pCtxt.Cause)

	//TODO do other process
	//store procedure info into pdu session ctxt

	//set next state
	err := pCtxt.SetNextState(statetype.StatePduSessEstbWfUpdateSmCtxtReq)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pduSessCtxt.IMSI, "fail to set state in sess estb proc")
	}
}
