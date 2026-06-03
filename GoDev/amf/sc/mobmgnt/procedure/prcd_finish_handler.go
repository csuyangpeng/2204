package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
)

func RegisterPrcdFinished(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, "no ue context.")
		return
	}

	//set state for ue
	err := ueCtxt.SetRmState(types.StateRmRegistered)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}

	//msg counter
	pm.PegCounter(statistics.OnLineUserCounter)

	//set the prcd ctxt to nil, finished the registration procedure
	ueCtxt.SetProcCtxt(nil)
}

func DeregisterPrcdFinished(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, "no ue context.")
		return
	}
	ueCtxt.SetPduSessCtxtNil()
	ueCtxt.PduSessStatus = ueCtxt.GetUePsiStatus()

	err := ueCtxt.SetRmState(types.StateRmDeRegistered)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}

	//msg counter
	pm.PegCounter(statistics.OffLineUserCounter)
	pm.UnPegCounter(statistics.OnLineUserCounter)

	if len(ueCtxt.GetPDUSessionCtxts()) == 0 {
		if pm.GetCounter(statistics.AmfStatisticRegistry, statistics.IdleUserCounterKey) > 0 {
			pm.UnPegCounter(statistics.IdleUserCounter)
		}
	}
	err = ueCtxt.SetCmState(types.CmIdle)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}

	ueCtxt.SetProcCtxt(nil)

	ueCtxt.TodoThing = 0
	ueCtxt.IsReRegistrationRequired = false

	//Release ueContext's index
	//AmfUeNgApId
	err = gctxt.DeleteUeContext(gctxt.AmfUeNgApId(ueCtxt.GetAmfUeNgapId()), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to release ueCtxt by index(AmfUeNgApId:%d),error(%s)",
			ueCtxt.GetAmfUeNgapId(), err)
		return
	}

	//return amf ue ngap id
	tmpId := ueCtxt.GetAmfUeNgapId() & 0xFFFFFFFF
	err = idmgr.GetInst().ReturnID(string(types.AMFUeNgapId), uint32(tmpId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to release AmfUeNgApId:(%d),error(%s)", ueCtxt.GetAmfUeNgapId(), err)
	}

	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(), "deRegister is done in AMF")
}

func ClearUEContextInAMF(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "no ue context.")
		return
	}
	ueCtxt.SetPduSessCtxtNil()
	ueCtxt.PduSessStatus = ueCtxt.GetUePsiStatus()

	err := ueCtxt.SetRmState(types.StateRmDeRegistered)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}

	//msg counter
	pm.ClearCounter(statistics.OffLineUserCounter)
	pm.ClearCounter(statistics.OnLineUserCounter)
	pm.ClearCounter(statistics.IdleUserCounter)
	pm.ClearCounter(statistics.ActiveUserCounter)

	err = ueCtxt.SetCmState(types.CmIdle)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}

	ueCtxt.SetProcCtxt(nil)

	//Release ueContext's index
	//AmfUeNgApId
	err = gctxt.DeleteUeContext(gctxt.AmfUeNgApId(ueCtxt.GetAmfUeNgapId()), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to release ueCtxt by index(AmfUeNgApId:%d) ,error(%s)", ueCtxt.GetAmfUeNgapId(), err)
		return
	}
}

func AnReleasePrcdFinished(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "no ue context.")
		return
	}
	ueCtxt.PduSessStatus = ueCtxt.GetUePsiStatus()
	err := ueCtxt.SetCmState(types.CmIdle)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}

	//msg counter
	pm.PegCounter(statistics.IdleUserCounter)
	if pm.GetCounter(statistics.AmfStatisticRegistry, statistics.ActiveUserCounterKey) > 0 {
		pm.UnPegCounter(statistics.ActiveUserCounter)
	}

	ueCtxt.SetProcCtxt(nil)

	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "an release done in AMF")
}

func ServiceRequestFinished(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "no ue context.")
		return
	}

	if len(ueCtxt.GetPsiList(types3gpp.SessActived)) > 0 {
		err := ueCtxt.SetCmState(types.CmConnected)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	}
	ueCtxt.PduSessStatus = ueCtxt.GetUePsiStatus()
	// get procedure context
	pCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	//set n2 context id
	ueCtxt.SetRanUeNgapId(pCtxt.GnbConnID)
	//TODO store n2 related info into ue context
	//TODO GnbInfo, TAI etc.

	ueCtxt.SetProcCtxt(nil)
	return
}

func SessionReleasePrcdFinished(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, "no ue context.")
		return
	}

	//get the prcdCtxt from ueCtxt
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.SessionReleasePrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get SessionReleasePrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	//set the next state
	err := prcdCtxt.SetNextState(statetype.StateSessRelEnd)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}

	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(),
		"clear pdu session context in ue context,psi :", prcdCtxt.Psi)
	ueCtxt.DelPduSessCtxt(prcdCtxt.Psi)

	ueCtxt.PduSessStatus = ueCtxt.GetUePsiStatus()

	ueCtxt.SetProcCtxt(nil)

	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
		"ue oriented session release done in AMF")
}
