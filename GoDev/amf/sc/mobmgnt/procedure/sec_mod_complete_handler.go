package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func HandleSecModCmpMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}

	// stop timer T3560
	err := gctxt.CancelPrcdTimer(ueCtxt, ctxt, gctxt.T3560)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to cancel T3560")
		return
	}

	//set status for FSM
	switch ueCtxt.GetProcCtxt().(type) {
	case *prcdctxt.RegistrationPrcdCtxt:
		HandleRegistrationMsg(ctxt)
	case *prcdctxt.ServiceRequestPrcdCtxt:
		HandleServiceRequestMsg(ctxt)
	}
}
