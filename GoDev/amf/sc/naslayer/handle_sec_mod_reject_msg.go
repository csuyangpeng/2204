package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *NasMgr) HandleSecurityModeReject(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	plainNasMsg *bytes.Reader) error {
	td := []interface{}{ctx, n2connData}
	rlogger.FuncEntry(types.ModuleAmfNas, td)

	p.securityModeReject.Reset()

	err := p.securityModeReject.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to decode "+
			"security mode complete nas message")
		return fmt.Errorf("failed to decode security mode complete message")
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "p.securityModeReject: %s", p.securityModeReject)

	//get ue context with amf ngap id
	ueContext, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "type assersion failed for ue ctext")
		return types.ErrFailFindUeCtxt
	}

	//cancel t3560
	err = gctxt.CancelPrcdTimer(ueContext, ctx, gctxt.T3560)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueContext.GetImsiPtr(),
			"failed to cancel T3560")
		return fmt.Errorf("failed to cancel T3560")
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "cancel t3560 timer for auth reject msg")

	switch p.securityModeReject.Cause {
	case nas.SecModRejectUnspec:
		fallthrough
	case nas.UeSecCapMismatch:
		// 24501 5.4.2.5 #24	security mode rejected, unspecified.
		// Upon receipt of the SECURITY MODE REJECT message,
		// the AMF shall stop timer T3560. The AMF shall also abort the ongoing procedure
		// that triggered the initiation of the NAS security mode control procedure.
		// Both the UE and the AMF shall apply the 5G NAS security context in use before the initiation of
		// the security mode control procedure, if any, to protect the SECURITY MODE REJECT message and
		// any other subsequent messages according to the rules in subclause 4.4.4 and 4.4.5.
		ueContext.SetProcCtxt(nil)
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "unknown reason(%s) for securityModeRejectMsg",
			p.securityModeReject.Cause)
	}
	return nil
}
