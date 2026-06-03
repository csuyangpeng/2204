package sessionmodification

import (
	"context"
	"lite5gc/amf/sc/mobmgnt/procedure"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *SessModProcFSM) SessModWf1stUpSmCtxtReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessMod1stUpSmCtxtReqMsg(ctxt)
}

func (p *SessModProcFSM) SessModWfN2SessReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessModN2SessReqMsg(ctxt)
}

func (p *SessModProcFSM) SessModWf2ndUpSmCtxtReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessMod2ndUpSmCtxtReqMsg(ctxt)
}

func (p *SessModProcFSM) SessModWfN2NASUplinkTransCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessModN2NASUplinkTransMsg(ctxt)
}

func (p *SessModProcFSM) SessModWf3rdUpSmCtxtReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessMod3rdUpSmCtxtReqMsg(ctxt)
}

func (p *SessModProcFSM) SessModEndCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessModEnd(ctxt)
}
