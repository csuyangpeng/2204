package sessionrelease

import (
	"context"
	"lite5gc/amf/sc/mobmgnt/procedure"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *SessRelProcFSM) SessRelReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessRelReqMsg(ctxt)
}

func (p *SessRelProcFSM) UpdateSmCtxtRespCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessUpdateSmCtxtResp(ctxt)
}

//func (p *SessRelProcFSM) N2ResRelRespCallback(e *fsm.Event) {
//	rlogger.FuncEntry(types.ModuleAmfState, nil)
//
//	ctxt := e.Args[0].(context.Context)
//
//	procedure.HandleN2ResRelResp(ctxt)
//}

func (p *SessRelProcFSM) UpSmCtxtRespSecCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.UpdateSmCtxt2ndRespMsg(ctxt)
}

func (p *SessRelProcFSM) EndCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSessRelEnd(ctxt)
}
