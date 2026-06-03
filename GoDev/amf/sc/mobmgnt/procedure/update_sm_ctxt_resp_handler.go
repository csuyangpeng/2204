package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
)

func HandleUpdateSmCtxtResp4AnRelease(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ctxt, "no ue context.")
		return
	}

	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.AnReleasePrcdCtxt)
	if ok {
		counter := prcdCtxt.GetCounter()

		// get pdu session
		for i := 0; i < len(prcdCtxt.ReleasePsiList); i++ {
			pduSessCtxt := ueCtxt.GetPduSessCtxt(nas.PduSessID(prcdCtxt.ReleasePsiList[i]))
			//set the pdu session status
			if prcdCtxt.UpCnxState == types3gpp.PduSessStatus(n11msg.DEACTIVATED) {
				pduSessCtxt.Status = types3gpp.SessDeactive
				rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
					"pduSessCtxt.Status:", pduSessCtxt.Status)
			}
		}

		// for an release update sm ctxt ack msg
		if counter-1 > 0 {
			counter-- // continue to wait for the rest of the message

			// reset the state
			err := prcdCtxt.SetNextState(statetype.StateAnRelWfUpSmCtxtAck)
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to set state")
				return
			}
			prcdCtxt.SetCounter(counter)
			return
		} else {
			ReleaseN2ConnId(ueCtxt)

			AnReleasePrcdFinished(ctxt)
		}
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get AnReleasePrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	return
}

func HandleUpdateSmCtxtResp4SerReq(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "no ue context.")
		return
	}

	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	if prcdCtxt.Order == prcdctxt.SerReqUpDataSmCtxtRespSecond {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"amf recevie update sm ctxt response second time for Server Request Procedure")

		//peg
		if ueCtxt.GetCmState() == types.CmIdle {
			pm.PegCounter(statistics.ActiveUserCounter)
			if pm.GetCounter(statistics.AmfStatisticRegistry, statistics.IdleUserCounterKey) > 0 {
				pm.UnPegCounter(statistics.IdleUserCounter)
			}
		}

		//cancel timer
		timerMgr, ok := ctxt.Value(types.ScTimerMgrCK).(*timermgr.TimerMgr)
		if !ok {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get timer manager")
			return
		}
		timerMgr.CancelTimer(prcdCtxt.TimerId)

		ServiceRequestFinished(ctxt)

		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"ueCtxt.TodoThing", ueCtxt.TodoThing)
		if ueCtxt.TodoThing == nas.DeregistrationRequestUeT {
			rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
				"start DeregistrationRequestUeT after service request")
			ueCtxt.SetProcCtxt(prcdctxt.NewDeRegistration())

			//  get the scnglayer
			sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
			if !ok {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to get ngap layer .")
				return
			}

			var deregistReqMsg nasmsg.DeRegistrationRequestUeTMsg
			deregistReqMsg.DeRegistrationType.SwithOff = false
			deregistReqMsg.DeRegistrationType.IsReRegistrationRequired = ueCtxt.IsReRegistrationRequired
			deregistReqMsg.DeRegistrationType.AccessType = nasie.ThreeGppAccess

			encodeMsg, err := deregistReqMsg.Encode()
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to encode deregister request msg,error %s", err)
				return
			}

			//msg counter
			pm.PegCounter(statistics.DeregistrationRequestUeCounter)

			//get gnb instance id for sending ngap msg
			gnbInstId, err := gctxt.GetGnbInstIdByAmfUeNgapId(ueCtxt.GetAmfUeNgapId())
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to get gnb instance id")
				return
			}

			// send the message in scng layer
			err, _ = sender.SendDownlinkNasTransport(ueCtxt, gnbInstId, encodeMsg)
			if err != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to send authentication request message")
				return
			}
		}
		return
	}

	//the first time
	counter := prcdCtxt.GetCounter()
	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "counter ", counter)
	if counter-1 > 0 {
		counter-- // continue to wait for the rest of the message
		// reset the state
		err := prcdCtxt.SetNextState(statetype.StateSrvReqWfUpSmCtxtResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
		prcdCtxt.SetCounter(counter)
		return
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"prcdCtxt.PsiIdx ", prcdCtxt.PsiIdx)
		// set pduSessCtxt Status
		for i := 0; i < len(prcdCtxt.PsiIdx); i++ {
			ueCtxt.GetPduSessCtxt(nas.PduSessID(prcdCtxt.PsiIdx[i])).Status = types3gpp.SessActived
		}
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"ueCtxt.GetUePsiStatus():", ueCtxt.GetUePsiStatus())

		//send initial ctxt setup response and service accetp to ue
		err := mmsender.SendServiceAccept(ctxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to send service accept msg")
			return
		}

		//clear n2 sm info in procedure context
		prcdCtxt.N2SmInfo = make(map[uint32][]byte)

		//set next status
		err = prcdCtxt.SetNextState(statetype.StateSrvReqWfInitCtxtSetupResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}
	}
}
