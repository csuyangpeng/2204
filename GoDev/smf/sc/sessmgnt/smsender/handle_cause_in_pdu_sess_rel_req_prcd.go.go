package smsender

import (
	"context"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/sessmgnt/smutility"
)

func ProcessPduSessRelReqCause(ctx context.Context, psi nas.PduSessID, cause nas.Sm5gCause) {
	rlogger.FuncEntry(types.ModuleSmfSM, ctx)
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, ctx, "cause:", cause)
	if cause != nas.SuccessNoReason {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctx, "session release reject cause:", nas.Sm5gCause(cause))
		err := SendPduSessRelRejectMsg(ctx, cause)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctx, "failed to handle pdu session release reject")
		}

		//release resource
		err = smutility.ReleasePduSessCtxt(ctx, psi)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to release ue ctxt and session ctxt")
			return
		}
	}
}
