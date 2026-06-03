package naslayer

import (
	"bytes"
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

func (p *NasMgr) HandleRegistrationCompleteMsg(ctx context.Context, plainNasMsg *bytes.Reader) error {

	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	p.registrationComplete.Reset()

	//decode the message
	err := p.registrationComplete.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to decode "+
			"registration request nas message")
		return types.ErrWorngIpcMsg
	}

	//get ue context with amf ngap id
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to get ue ctxt")
		return types.ErrFailFindUeCtxt
	}

	//get procedure context
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return types.ErrFailGetProcedureCtxt
	}

	//trigger the FSM
	err = statemgr.TriggerFsm(ctx,
		statemgr.Register,
		prcdCtxt.GetCurrentState(),
		statetype.EventRegisterComplete)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
		return types.ErrFailTriggerFsm
	}

	return nil
}
