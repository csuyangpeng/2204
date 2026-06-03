package n11layer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	statetype "lite5gc/cmn/types/statetypes"
)

func HandleReleaseSmCtxtResponse(ctx context.Context, smResp *sbicmn.SbiPostReleaseSmContext) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)
	//get the ueCtxt from ctxt
	imsi, _, err := nas.GetSmfKeys(smResp.SmContextRef)
	if err != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil, "failed to get imsi and psi from url(%s)", smResp.SmContextRef)
		return fmt.Errorf("failed to get imsi and psi from url(%s)", smResp.SmContextRef)
	}

	ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil, "failed to find the ue context with AmfUeNgApId %s,error %s",
			imsi, err)
		return fmt.Errorf("failed to find the ue context with AmfUeNgApId %s,error %s",
			imsi, err)
	}

	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	//get the context Mgr
	stateMgr, ok := ctx.Value(types.StateMgrCK).(*statemgr.StateMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, ueCtxt, "get state manager failed")
		return fmt.Errorf("get state manager failed")
	}

	//get procedure context
	prcdCtxt := ueCtxt.GetProcCtxt()
	//TODO switch case the prcdCtxt to deregister procedrue and others
	if prcdCtxt != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, ueCtxt, "prcdCtxt.GetCurrentState()", prcdCtxt.GetCurrentState())
		//trigger the FSM
		stateMgr.DeRegisterFsm.Bfsm.SetState(prcdCtxt.GetCurrentState())
		stateMgr.DeRegisterFsm.Bfsm.Event(statetype.EventDeRegisterRelSmCtxtResp, ctx)
	}
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "release resource done")

	return nil
}
