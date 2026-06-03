/*
* Copyright(C),2020-2022
* Author:  xiaoyun
* Date:    11/25/20 5:01 AM
* Description:
 */
package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	statetype "lite5gc/cmn/types/statetypes"
)

func (p *NasMgr) HandleSecModCmpMsg(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	plainNasMsg *bytes.Reader) error {
	td := []interface{}{ctx, n2connData}
	rlogger.FuncEntry(types.ModuleAmfNas, td)

	p.securityModeComplete.Reset()

	err := p.securityModeComplete.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td,
			"failed to decode security mode complete nas message")
		return fmt.Errorf("failed to decode security mode complete message")
	}

	//get ue context with amf ngap id
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "type assersion failed for ue ctext")
		return types.ErrFailFindUeCtxt
	}

	// finish the SMC procedure. set the security context establish flag in ue context
	ueCtxt.SecurityCtxt.IsSecCtxtEstb = true
	ueCtxt.SecurityCtxt.ForceAuthNeed = false //finished SMC procedure
	// set the security header to 2
	ueCtxt.SecHdrTypeIncomingMsg = nas.IntegrityPrtctCipher

	// TODO store the registration request message in procedure context

	//get procedure context
	switch ueCtxt.GetProcCtxt().(type) {
	case *prcdctxt.RegistrationPrcdCtxt:
		pCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
		if !ok {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return types.ErrFailGetProcedureCtxt
		}
		//trigger the FSM
		err := statemgr.TriggerFsm(ctx,
			statemgr.Register,
			pCtxt.GetCurrentState(),
			statetype.EventRegisterSecModCmp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	case *prcdctxt.ServiceRequestPrcdCtxt:
		pCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
		if !ok {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return types.ErrFailGetProcedureCtxt
		}
		err := statemgr.TriggerFsm(ctx,
			statemgr.ServiceRequest,
			pCtxt.GetCurrentState(),
			statetype.EventSrvReqSecModCmp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "unsupported procedure context")
		return fmt.Errorf("unsupported procedure context error(%s)", err)
	}
	return nil
}
