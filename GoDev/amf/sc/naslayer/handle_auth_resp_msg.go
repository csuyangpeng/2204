package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

func (p *NasMgr) HandleAuthRespMsg(ctx context.Context, plainNasMsg *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	p.authenticationResponse.Reset()

	err := p.authenticationResponse.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to decode "+
			"authentication response nas message")
		return fmt.Errorf("failed to decode authentication response message")
	}

	//get ue context with amf ngap id
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "type assersion failed for ue ctext")
		return types.ErrFailFindUeCtxt
	}

	// reset the flag
	ueCtxt.SecurityCtxt.SyncFailureProceeding = false
	//get procedure context
	prcdCtxt := ueCtxt.GetProcCtxt()

	if prcdCtxt != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, ueCtxt.GetImsiPtr(), "ueCtxt.GetProcCtxt() : %T ,CurrentState :%s", prcdCtxt, prcdCtxt.GetCurrentState())
		switch prcdCtxt.(type) {
		case *prcdctxt.RegistrationPrcdCtxt:
			//get procedure context
			prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
			if !ok {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
				return types.ErrFailGetProcedureCtxt
			} else {
				//store the info inout prcdCtxt
				prcdCtxt.ResStart = p.authenticationResponse.ResStar

				err := statemgr.TriggerFsm(ctx,
					statemgr.Register,
					prcdCtxt.GetCurrentState(),
					statetype.EventRegisterAuthResp)
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
					return types.ErrFailTriggerFsm
				}
			}
		case *prcdctxt.ServiceRequestPrcdCtxt:
			//get procedure context
			prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
			if !ok {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
					"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
				return types.ErrFailGetProcedureCtxt
			} else {
				//store the info inout prcdCtxt
				prcdCtxt.ResStart = p.authenticationResponse.ResStar

				err := statemgr.TriggerFsm(ctx,
					statemgr.ServiceRequest,
					prcdCtxt.GetCurrentState(),
					statetype.EventSrvReqAuthResp)
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
					return types.ErrFailTriggerFsm
				}
			}
		}
	} else {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "fail to get prcdCtxt from ueContext")
		return fmt.Errorf("fail to hand auth resp msg")
	}

	return nil
}
