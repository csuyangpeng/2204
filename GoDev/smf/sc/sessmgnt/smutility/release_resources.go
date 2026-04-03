package smutility

import (
	"context"
	"fmt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/ueippool"
)

func ReleasePduSessCtxt(ctxt context.Context, psi nas.PduSessID) error {
	rlogger.FuncEntry(types.ModuleSmfSM, ctxt)

	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to get ue context")
		return fmt.Errorf("failed to get smf ue context")
	}

	pduSessCtxt := ueCtxt.GetPduSessCtxt(psi)
	pduSessCtxt.SetPrcdCtxt(nil)

	if pduSessCtxt != nil {
		//release ip by dnn
		err := ueippool.ReleaseIPByDNN(ueippool.DnnNameKey(pduSessCtxt.DNN.String()), pduSessCtxt.UEIP)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, ctxt,
				"fail to get release by dnn: %s", pduSessCtxt.DNN.String())
		}
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, ctxt,
			"release ip(%s) [dnn(%s)] for ue(%s)", pduSessCtxt.UEIP, pduSessCtxt.DNN.String(), ueCtxt.IMSI.String())
		pduSessCtxt.UEIP = nil

		err = gctxt.DeleteSessionContext(gctxt.SeidKey(pduSessCtxt.SEID))
		if err != nil {
			rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to del session ctxt")
			return fmt.Errorf("fail to del session ctxt")
		}
	}

	err := ueCtxt.DeletePduSessCtxt(psi)
	if err != nil {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "fail to del session ctxt")
		return fmt.Errorf("fail to del session ctxt")
	}

	if len(ueCtxt.PduSessCtxts) == 0 {
		//pm.PegCounter(statistics.IdleUserCounter)
		if pm.GetCounter(statistics.AmfStatisticRegistry, statistics.ActiveUserCounterKey) > 0 {
			pm.UnPegCounter(statistics.ActiveUserCounter)
		}
	}

	return nil
}
