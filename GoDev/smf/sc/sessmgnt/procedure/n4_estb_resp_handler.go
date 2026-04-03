package procedure

import (
	"context"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/smf/sc/sessmgnt/smsender"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func HandleN4SessEstbResp(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	// pdu session context
	pCtxt, ok := ctxt.Value(types.PduSessionEstbPrcdCtxtCK).(*prcdctxt.PduSessionEstbPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get pdu session context")
		return
	}

	// Step 3: send pdu session establish accept to UE
	//         send n1n2 transfer msg to ran
	err := smsender.SendPduSessEstbAcceptMsg(ctxt, pCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to send pdu sess estb accept msg")
		return
	}

	// set the next state
	err = pCtxt.SetNextState(statetype.StatePduSessEstbWfN1N2MsgTransferResp)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to set state in sess estb proc")
	}
}
