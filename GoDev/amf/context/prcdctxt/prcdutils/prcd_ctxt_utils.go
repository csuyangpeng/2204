package prcdutils

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
)

func CancelAmfTimer(ctx context.Context, ueCtxt *gctxt.UeContext) error {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	timerMgr, ok := ctx.Value(types.ScTimerMgrCK).(*timermgr.TimerMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.ERROR, nil, "failed to get timer manager")
		return types.ErrFailFindTimerMgr
	}
	prcdCtxt := ueCtxt.GetProcCtxt()
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.INFO, nil, "prcdCtxt", prcdCtxt)
	if prcdCtxt != nil {
		switch prcdCtxt.(type) {
		case *prcdctxt.RegistrationPrcdCtxt:
			pc := prcdCtxt.(*prcdctxt.RegistrationPrcdCtxt)
			timerMgr.CancelTimer(pc.TimerId)
		case *prcdctxt.ServiceRequestPrcdCtxt:
			pc := prcdCtxt.(*prcdctxt.ServiceRequestPrcdCtxt)
			timerMgr.CancelTimer(pc.TimerId)
		default:
			rlogger.Trace(types.ModuleAmfCtxt, rlogger.ERROR, nil, "un support PrcdCtxt")
			return fmt.Errorf("un support PrcdCtxt")
		}
	}
	return nil
}
