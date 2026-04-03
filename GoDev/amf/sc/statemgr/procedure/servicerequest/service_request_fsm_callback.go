package servicerequest

import (
	"context"
	"lite5gc/amf/sc/mobmgnt/procedure"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// PagingReqCallback is callback function
func (p *SerReqProcFSM) PagingReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	//TODO
}

// ServiceReqCallback is callback function   1:no psi
func (p *SerReqProcFSM) SrvReqServiceReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	//TODO
	ctxt := e.Args[0].(context.Context)
	procedure.HandleServiceRequestMsg(ctxt)
}

// InitCtxtSetupRespCallback is callback function
func (p *SerReqProcFSM) SrvReqUpSmCtxtRespCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)
	procedure.HandleUpdateSmCtxtResp4SerReq(ctxt)
}

// send init ctxt setup request && service accetp msg out
func (p *SerReqProcFSM) SrvReqInitCtxtSetupRespCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)
	procedure.HandleInitCtxtSetupRespMsg4SerReq(ctxt)
}

// UpdateSmCtxtRespCallback is callback function
func (p *SerReqProcFSM) SrvReqUpSmCtxtRespSecCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)
	procedure.HandleUpdateSmCtxtResp4SerReq(ctxt)
}

func (p *SerReqProcFSM) SrvReqFinishedCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)
	procedure.ServiceRequestFinished(ctxt)
}

func (p *SerReqProcFSM) SrvReqAuthCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)
	procedure.HandleAuthRequest(ctxt)
}

func (p *SerReqProcFSM) SrvReqAuthRespCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)
	procedure.HandleAuthResponseMsgSerReq(ctxt)
}

func (p *SerReqProcFSM) SrvReqSecModCmpCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)
	procedure.HandleSecModCmpMsg(ctxt)
}
