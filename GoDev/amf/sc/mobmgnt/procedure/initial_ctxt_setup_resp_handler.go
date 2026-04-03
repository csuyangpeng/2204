package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/oam/pm"
)

func HandleInitCtxtSetupRespMsg4Register(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}

	//get the register procedure context
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}
	// send register accept msg
	err := mmsender.SendRegistrationAccept(ctxt, ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to send register accept msg")
		return
	}

	//set the next state for FSM
	err = prcdCtxt.SetNextState(statetype.StateRegisterWfRegisterCmp)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to set state")
		return
	}
}

func HandleInitCtxtSetupRespMsg4SerReq(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}

	// get procedure context
	procCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	sendMsgToSMF := false

	updateSmCtxtReq := n11msg.UpdateSMContextRequestData{}
	for i := 0; i < len(procCtxt.PduSessResSetupRespList); i++ {
		respItem := procCtxt.PduSessResSetupRespList[i]
		updateSmCtxtReq.N2SmInfo = respItem.PduSessResSetupRespTransfer
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n2SmInfo)

		updateSmCtxtReq.N2SmInfoType = n11msg.PduResSetupRsp
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n2SmInfoType)

		updateSmCtxtReq.UpCnxState = n11msg.ACTIVATED
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_upCnxState)

		psi := nas.PduSessID(respItem.PduSessionId)
		mmsender.SendUpdateSMCtxtSBIMsg(ctxt, psi, ueCtxt.GetImsi(), updateSmCtxtReq)
		//set next state
		procCtxt.Order = prcdctxt.SerReqUpDataSmCtxtRespSecond
		err := procCtxt.SetNextState(statetype.StateSrvReqWfUpSmCtxtRespSec)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}

		sendMsgToSMF = true
	}
	for i := 0; i < len(procCtxt.PduSessResFailedToSetupList); i++ {
		failedItem := procCtxt.PduSessResFailedToSetupList[i]
		updateSmCtxtReq.N2SmInfo = failedItem.UnsuccessTransfer
		updateSmCtxtReq.N2SmInfoType = n11msg.PduResSetupFail
		psi := nas.PduSessID(failedItem.PduSessionId)
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n2SmInfo)
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_n2SmInfoType)
		updateSmCtxtReq.UpCnxState = n11msg.DEACTIVATED
		updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_upCnxState)
		mmsender.SendUpdateSMCtxtSBIMsg(ctxt, psi, ueCtxt.GetImsi(), updateSmCtxtReq)

		//set next state
		procCtxt.Order = prcdctxt.SerReqUpDataSmCtxtRespSecond
		err := procCtxt.SetNextState(statetype.StateSrvReqWfUpSmCtxtRespSec)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}

		sendMsgToSMF = true
	}

	if !sendMsgToSMF {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"no update sm ctxt request need send to smf")
		timerMgr, ok := ctxt.Value(types.ScTimerMgrCK).(*timermgr.TimerMgr)
		if !ok {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to get timer manager")
			return
		}

		prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
		if !ok {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return
		}
		timerMgr.CancelTimer(prcdCtxt.TimerId)

		ServiceRequestFinished(ctxt)

		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "ueCtxt.TodoThing", ueCtxt.TodoThing)
		if ueCtxt.TodoThing == nas.DeregistrationRequestUeT {
			rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
				"start DeregistrationRequestUeT after service request")
			ueCtxt.SetProcCtxt(prcdctxt.NewDeRegistration())

			var deregisterReqMsg nasmsg.DeRegistrationRequestUeTMsg
			deregisterReqMsg.DeRegistrationType.SwithOff = false
			deregisterReqMsg.DeRegistrationType.IsReRegistrationRequired = false
			deregisterReqMsg.DeRegistrationType.AccessType = nasie.ThreeGppAccess
			encodeMsg, err := deregisterReqMsg.Encode()
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to encode deregister request msg,error %s", err)
				return
			}

			//msg counter
			pm.PegCounter(statistics.DeregistrationRequestUeCounter)

			err, _ = mmsender.SendDownLinkNasMsg(ctxt, ueCtxt, encodeMsg)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"fail to send downlink nas ngap msg")
				return
			}
		}
	}
	return
}
