package anrelease

import (
	"context"
	"lite5gc/amf/sc/mobmgnt/procedure"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// UeCtxtRelReqCallback is callback function
func (p *AnRelProcFSM) UeCtxtRelReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleUeContextRelRequestMsg(ctxt)
}

func (p *AnRelProcFSM) UpdateSmCtxtAckCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleUpdateSmCtxtResp4AnRelease(ctxt)
}

func (p *AnRelProcFSM) UeCtxtRelCmpCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleUeContextRelCompleteMsg(ctxt)
}
