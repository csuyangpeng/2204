package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/mobmgnt/mmutils"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

func HandleServiceRequestMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, ctxt)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ctxt, "no ue context found.")
		//TODO trigger service reject
		return
	}

	// get procedure context
	procCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return
	}

	// check the rm state
	if ueCtxt.GetRmState() != types.StateRmRegistered {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"this ue has not registered")
		//TODO trigger service reject
		return
	}

	// check the ue state
	//if ueCtxt.GetCmState() != statetype.CmIdle {
	if len(ueCtxt.GetPsiList(types3gpp.SessActived)) != 0 {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"this ue is connected state")

		// means the service request is trigger by initial ue request msg
		// trigger an release for old n2 connection
		if ueCtxt.GetAmfUeNgapId() != ueCtxt.GetAmfUeNgapId() {
			if mmsender.SendUeContextReleaseCmdSim(ctxt, types3gpp.Radiok_ue_context_transfer) != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to send an release for old n2 connection, AmfNgapId(%d)", ueCtxt.GetAmfUeNgapId())
			}
		}
	}

	if procCtxt.IeFlags.Test(nasmsg.IeidServicereqAllowedpdusessioinstaus) {
		// TODO non 3GPP case
	}

	var UppsiIdx, PdupsiIdx, psiIdx []byte

	if procCtxt.IeFlags.Test(nasmsg.IeidServicereqUplinkdatastatus) || procCtxt.IeFlags.Test(nasmsg.IeidServicereqPdusessionstatus) {
		// count how many UpdateSmCtxtReq messages sent to the SMF
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"Ieid_ServiceReq_UplinkDataStatus PsiA (%x), PsiB(%x)",
			procCtxt.UplinkDataStatus.PsiA, procCtxt.UplinkDataStatus.PsiB)
		mmutils.GetPsiIdx(procCtxt.UplinkDataStatus.PsiA, 0, &UppsiIdx)
		mmutils.GetPsiIdx(procCtxt.UplinkDataStatus.PsiB, 1, &UppsiIdx)
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"UppsiIdx from Ieid_ServiceReq_UplinkDataStatus, %v", UppsiIdx)

		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"Ieid_ServiceReq_PDUSessionStatus PsiA (%x), PsiB(%x)",
			procCtxt.PDUSessionStatus.PsiA, procCtxt.PDUSessionStatus.PsiB)
		mmutils.GetPsiIdx(procCtxt.PDUSessionStatus.PsiA, 0, &PdupsiIdx)
		mmutils.GetPsiIdx(procCtxt.PDUSessionStatus.PsiB, 1, &PdupsiIdx)
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"PdupsiIdx from Ieid_ServiceReq_PDUSessionStatus, %v", PdupsiIdx)

		psiIdx = mmutils.Union(UppsiIdx, PdupsiIdx)
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "Union psi idx , %v", psiIdx)
	} else {
		//get from ue context and trigger update sm context
		psiIdx = ueCtxt.GetPsiList(types3gpp.SessDeactive)
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
			"psi idx from ueCtxt.GetPsiList(types3gpp.SessDeactive), %v", psiIdx)
	}

	procCtxt.PsiIdx = psiIdx

	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "len(psiIdx): ", len(psiIdx))

	if len(psiIdx) <= 0 {
		ctxt = context.WithValue(ctxt, types.UeContextCK, ueCtxt)
		//send initial ctxt setup response and service accetp to ue
		err := mmsender.SendServiceAccept(ctxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to send service accept msg")
		}
		// set the next status
		err = procCtxt.SetNextState(statetype.StateSrvReqWfInitCtxtSetupResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to set state")
			return
		}

		var counter int
		for _, i := range psiIdx {
			if i == 0 {
				continue
			}
			if ueCtxt.GetPDUSessionCtxts()[nas.PduSessID(i)] != nil {
				rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "psi: ", i)
				updateSmCtxtReq := n11msg.UpdateSMContextRequestData{}
				updateSmCtxtReq.UpCnxState = n11msg.ACTIVATED
				updateSmCtxtReq.IeFlags.Set(n11msg.Ieid_upCnxState)
				mmsender.SendUpdateSMCtxtSBIMsg(ctxt, nas.PduSessID(i), ueCtxt.GetImsi(), updateSmCtxtReq)
				counter++
				procCtxt.SetCounter(counter)
				rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "counter: ", counter)
				// set the order
				procCtxt.Order = prcdctxt.SerReqUpDataSmCtxtRespFirst
				// set the next status
				err = procCtxt.SetNextState(statetype.StateSrvReqWfUpSmCtxtResp)
				if err != nil {
					rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
						"failed to set state")
					return
				}
			}
		}
	}
}
