package naslayer

import (
	"bytes"
	"context"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/oam/pm"
	"lite5gc/smf/sc/statemgr"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func (p *NasMgr) HandleSmPduSessRelRequest(ctx context.Context,
	header *nas.SmNasMessageHeader, plainNasMsg *bytes.Reader) nas.Sm5gCause {

	rlogger.FuncEntry(types.ModuleSmfNas, ctx)
	if plainNasMsg == nil || header == nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "input para is nil")
		return nas.OtherValue
	}
	//msg counter
	pm.PegCounter(statistics.PduSessReleaseRequestCounter)

	cause := nas.SuccessNoReason

	p.PduSessRelRequest.Reset()

	//decode the PduSessEstbRequest message
	p.PduSessRelRequest.MsgHeader.PduSessionID = header.PduSessionID
	p.PduSessRelRequest.MsgHeader.PrcdTransactionID = header.PrcdTransactionID
	err := p.PduSessRelRequest.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "failed to decode pdu session release request message, err: ", err)
		return nas.OtherValue
	}

	//get ue context
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "failed to get ue context")
		return nas.OtherValue
	}

	// check pti
	pti := header.PrcdTransactionID
	if pti < nas.MinPTI || pti > nas.MaxPTI {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "pti is out of range")
		return nas.InvalidPTIValue
	}

	//check the pdu session id exist
	psi := header.PduSessionID
	if psi < nas.MinPSI || psi > nas.MaxPSI {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "psi value is out of range")
		return nas.InvalidPDUSessionIdentity
	}

	//get the pdu session ctxt
	pduSessCtxt := ueCtxt.GetPduSessCtxt(psi)
	if pduSessCtxt == nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "failed to get pdu session context")
		return nas.OtherValue
	}

	if psi != pduSessCtxt.PduSessionId {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx, "psi mismatch between nas and smCtxtCreateData")
		return nas.InvalidPDUSessionIdentity
	}

	prcdCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionRelReqPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, ctx,
			"prcdctxt is not PduSessionRelReqPrcdCtxt, but:",
			gctxt.PrcdCtxt2String(pduSessCtxt.GetPrcdCtxt()))
		return nas.OtherValue
	}

	if p.PduSessRelRequest.IeFlags.Test(nas.Ieid_SMCause) {
		prcdCtxt.Cause = p.PduSessRelRequest.SMCause
		prcdCtxt.IsCauseExist = true
	} else {
		prcdCtxt.IsCauseExist = false
	}

	prcdCtxt.Pti = pti
	prcdCtxt.PduSessId = psi
	prcdCtxt.Imsi = ueCtxt.IMSI
	pduSessCtxt.SetPrcdCtxt(prcdCtxt)
	ctx = context.WithValue(ctx, types.SmfPduSessCtxtCK, pduSessCtxt)

	err = statemgr.TriggerSmfFsm(ctx,
		statemgr.SessionRel,
		prcdCtxt.GetCurrentState(),
		statetype.EventPduSessRelReq)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
		return nas.OtherValue
	}

	return cause
}
