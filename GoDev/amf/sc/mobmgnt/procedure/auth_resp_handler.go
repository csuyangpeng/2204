package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/statistics"
	"lite5gc/amf/security/seaf"
	"lite5gc/amf/security/secmgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/oam/pm"
)

func HandleAuthResponseMsgRegistration(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, "no ue context.")
		return
	}

	// stop timer T3560
	err := gctxt.CancelPrcdTimer(ueCtxt, ctxt, gctxt.T3560)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to cancel T3560")
		return
	}

	//get the register procedure ctxt
	if ueCtxt.GetProcCtxt() != nil {
		prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
		if !ok {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return
		}

		// match RES* -> HRES*   with   XHRES*
		err := seaf.HandleAuthenticationResponse(ueCtxt, prcdCtxt.ResStart[:])
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed for RES verification, error(%s)", err)
			//TODO send authentication failure to UE
			return
		}

		// Select the Enc And Int Algo
		selSecAlgo, err := secmgr.SelectNasAlgo(ueCtxt, &prcdCtxt.UeSecCapablity)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to select Nas Security Algo, error(%s)", err)
			//TODO send authentication failure to UE
			return
		}

		// send Security Mode Command message to UE
		err = mmsender.SendSecurityModeCommandMsg(ctxt, selSecAlgo)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to send security mode cmd msg")
			return
		}

		//msg counter
		pm.PegCounter(statistics.SecurityModeCommandCounter)

		//set the next state
		err = prcdCtxt.SetNextState(statetype.StateRegisterWfSecModCmp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"ueCtxt.GetProcCtxt() is nil")
	}
}

func HandleAuthResponseMsgSerReq(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, "no ue context.")
		return
	}

	// stop timer T3560
	err := gctxt.CancelPrcdTimer(ueCtxt, ctxt, gctxt.T3560)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to cancel T3560")
		return
	}

	if ueCtxt.GetProcCtxt() != nil {
		//get the register procedure ctxt
		prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
		if !ok {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return
		}

		// match RES* -> HRES*   with   XHRES*
		err = seaf.HandleAuthenticationResponse(ueCtxt, prcdCtxt.ResStart[:])
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed for RES verification, error(%s)", err)
			//TODO send authentication failure to UE
		}

		// Select the Enc And Int Algo
		selSecAlgo, err := secmgr.SelectNasAlgo(ueCtxt, &prcdCtxt.UeSecCapablity)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to select Nas Security Algo, error(%s)", err)
			//TODO send authentication failure to UE
			return
		}

		// send Security Mode Command message to UE
		err = mmsender.SendSecurityModeCommandMsg(ctxt, selSecAlgo)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to send security mode cmd msg")
			return
		}

		//msg counter
		pm.PegCounter(statistics.SecurityModeCommandCounter)

		//set the next state
		err = prcdCtxt.SetNextState(statetype.StateSrvReqWfSecModCmp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"ueCtxt.GetProcCtxt() is nil")
	}
}
