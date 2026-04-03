package naslayer

import (
	"context"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/oam/pm"
	"lite5gc/smf/sc/statemgr"
	"lite5gc/smf/smfcontext/gctxt"
)

func (p *NasMgr) HandleSmPduSessRelComplete(ctx context.Context,
	header *nas.SmNasMessageHeader) {
	if header == nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "input para is nil")
		return
	}
	rlogger.FuncEntry(types.ModuleSmfNas, ctx)
	//msg counter
	pm.PegCounter(statistics.PduSessReleaseCompleteCounter)

	//get ue context
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "failed to get ue context")
	}
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	//get the pdu session ctxt
	pduSessCtxt := ueCtxt.GetPduSessCtxt(header.PduSessionID)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "failed to get pdu session context")
	}

	err := statemgr.TriggerSmfFsm(ctx,
		statemgr.SessionRel,
		pduSessCtxt.GetPrcdCtxt().GetCurrentState(),
		statetype.EventPduSessRelUpdateSmCtxtReqSec)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
		return
	}
}
