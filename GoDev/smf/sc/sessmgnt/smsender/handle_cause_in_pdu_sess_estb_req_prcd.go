package smsender

import (
	"context"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func ProcessPduSessEstbReqCause(ctx context.Context, cause nas.Sm5gCause) {
	rlogger.FuncEntry(types.ModuleSmfSM, nil)

	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, ctx, "cause(%s)", nas.Mm5gCause(cause))

	if cause != nas.SuccessNoReason {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctx, "session reject cause: %x", nas.Mm5gCause(cause))
		err := SendPduSessEstbRejectMsg(ctx, cause)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctx, "failed to handle pdu session establish reject")
		}
	}
}
