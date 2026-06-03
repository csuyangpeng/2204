package naslayer

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *NasMgr) HandleAuthReject(ctx context.Context) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	//get ue context with amf ngap id
	ueContext, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "type assersion failed for ue ctext")
		return types.ErrFailFindUeCtxt
	}

	//cancel t3560
	err := gctxt.CancelPrcdTimer(ueContext, ctx, gctxt.T3560)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueContext.GetImsiPtr(), "fail to cancel t3560 timer")
		return err
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueContext.GetImsiPtr(), "cancel t3560 timer for auth reject msg")

	return nil
}
