/*
* Copyright(C),2020-2022
* Author:  xiaoyun
* Date:    11/13/20 3:29 AM
* Description:
	处理初始注册请求消息，看是否需要进行鉴权流程
*/
package naslayer

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/mobmgnt/procedure"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types/statetypes"
)

func (p *NasMgr) handleInitRegisterRequest(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	ueCtxt *gctxt.UeContext) nas.Mm5gCause {

	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt.GetImsiPtr())

	// check rm state
	// 24.501 5.5.1.2.8 f)
	// If a REGISTRATION REQUEST message with 5GS registration type IE set to "initial registration"
	// is received in state 5GMM-REGISTERED the network may initiate the 5GMM common procedures;
	// if it turned out that the REGISTRATION REQUEST message was sent by a UE that has already been registered,
	// the 5GMM context, if any, are deleted and the new REGISTRATION REQUEST is progressed.
	if ueCtxt.GetRmState() == types.StateRmRegistered {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ueCtxt.GetImsiPtr(), "this ue has registered already")
		for i, v := range ueCtxt.GetPDUSessionCtxts() {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
				"release pdu sess psi(%i)", ueCtxt.GetPDUSessionCtxts()[i].Psi)
			mmsender.SendReleaseSMCtxtSBIMsg(ctx, v.Psi, ueCtxt.GetImsi(), n11msg.ReleaseSMContextRequestData{})
		}
		ueCtxt.SetPduSessCtxtNil()
	}

	// check security capability
	// 24.501 5.5.1.2.8  i)
	// If the REGISTRATION REQUEST message is received with invalid or unacceptable UE security capabilities
	// (e.g. no 5GS encryption algorithms (all bits zero), no 5GS integrity algorithms (all bits zero),
	// mandatory 5GS encryption algorithms not supported or mandatory 5GS integrity algorithms not supported, etc.),
	// the AMF shall return a REGISTRATION REJECT message
	configSecCap := configure.AmfConf.NAS.SecCap
	if !configSecCap.HaveCommonAlgo(p.registrationRequest.UeSecCapablity.GetNrIntPrctAlgo(),
		p.registrationRequest.UeSecCapablity.GetNrEncAlgo()) {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ueCtxt.GetImsiPtr(),
			"No collection of UE security algorithms")
		return nas.UeSecCapMismatch
	}

	// 24.501 5.5.1.2.8 e)
	if ueCtxt.GetProcCtxt() != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ueCtxt.GetImsiPtr(),
			"current procedure context is already exist")
		ueCtxt.SetProcCtxt(nil)
	}

	//allocate the procedure context
	ueRegistPrcdCtxt := prcdctxt.NewRegistrationPcrd(ueCtxt.GetImsiPtr())
	ueRegistPrcdCtxt.FillCtxtFromMsg(&p.registrationRequest)
	ueRegistPrcdCtxt.NeedInitCtxtSetupPrcd = n2connData.NeedInitCtxtSetupPrcd
	ueRegistPrcdCtxt.NgKSI = p.registrationRequest.NgKSI
	ueRegistPrcdCtxt.UeSecCapablity = p.registrationRequest.UeSecCapablity
	//todo add more info in procedure context here

	ueCtxt.SetProcCtxt(ueRegistPrcdCtxt)

	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ueCtxt.GetImsiPtr(),
		"ran_ue_ngap_id(%x)", ueCtxt.GetRanUeNgapId())

	PerformAuthCiphRegist(ctx, ueCtxt.SecurityCtxt, ueRegistPrcdCtxt)

	return nas.SuccessAccept
}

func PerformAuthCiphRegist(ctxt context.Context,
	secContext gctxt.SecurityCtxt,
	prcd *prcdctxt.RegistrationPrcdCtxt) nas.Mm5gCause {

	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	// check authentication
	authNeeded := procedure.CheckAuthentication(secContext, prcd.NgKSI, prcd.UeSecCapablity)

	//trigger the FSM
	if authNeeded {
		rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "authentication required")
		// initial registration with aka
		err := statemgr.TriggerFsm(ctxt,
			statemgr.Register,
			statetype.StateRegisterAuthStart,
			statetype.EventAuthRequest)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return nas.SystemFailure
		}
	} else {
		rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "no authentication required")
		// initial registration without aka
		err := statemgr.TriggerFsm(ctxt,
			statemgr.Register,
			statetype.StateRegisterStart,
			statetype.EventRegisterRequest)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return nas.SystemFailure
		}
	}
	return nas.SuccessAccept
}
