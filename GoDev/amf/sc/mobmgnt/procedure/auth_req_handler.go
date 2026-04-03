package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

// trigger send authentication request message to UE
func HandleAuthRequest(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, "no ue context.")
		return
	}

	// call sbi api
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(), "send SBI GetAuthData service request")
	// for debug
	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "ue_ngap_id(%x)", ueCtxt.GetRanUeNgapId())

	var getAuthData = &sbicmn.SbiGetAuthDataMsg{}
	imsi := ueCtxt.GetImsi()
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(), "imsi: %s", imsi.String())
	getAuthData.Supi = imsi.AddIMSIPrefix()

	scId, ok := ctxt.Value(types.ScIdCK).(uint32)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "no sc id found")
		return
	}
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(), "sc id: %d", scId)

	message := sbicmn.SbiMessage{}
	message.MsgData = getAuthData
	message.MsgType = sbicmn.GetAuthDataMsgRequest
	message.ScInstId = scId

	err := routeragent.SendIpcMessage(ctxt, router.SbiGR, message.ScInstId, &message)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to send sbi msg to udm. error(%s)", err)
		return
	}

	//get procedure context
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	} else {
		err = prcdCtxt.SetNextState(statetype.StateRegisterWfUdmGetAuth)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to set state")
			return
		}
	}
}

func HandleAuthRequestForUdmResp(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "no ue context.")
		return
	}

	// print the status changes
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(), "Current RM State is (%s), "+
		"Register Procedure State is (%s).", ueCtxt.GetRmState(), ueCtxt.GetProcCtxt().GetCurrentState())

	//for debug
	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "ran_ue_ngap_id(%x) ",
		ueCtxt.GetRanUeNgapId())

	// update ksi
	ueCtxt.NgKsi.Update()

	err := mmsender.SendAuthenticationRequestMsg(ctxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to send authentication request msg")
		return
	}

	//set status for FSM
	switch ueCtxt.GetProcCtxt().(type) {
	case *prcdctxt.RegistrationPrcdCtxt:
		prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
		if !ok {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return
		} else {
			err = prcdCtxt.SetNextState(statetype.StateRegisterWfAuthResp)
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to set state")
				return
			}
		}
	case *prcdctxt.ServiceRequestPrcdCtxt:
		prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
		if !ok {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return
		} else {
			err = prcdCtxt.SetNextState(statetype.StateSrvReqWfAuthResp)
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to set state")
				return
			}
		}
	}
}
