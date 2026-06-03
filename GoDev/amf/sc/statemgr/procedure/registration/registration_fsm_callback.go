package registration

import (
	"context"
	"lite5gc/amf/sc/mobmgnt/procedure"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// RegisterRequestCallback is callback function for register request message event
func (p *RegisterProcFSM) RegisterRequestCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleRegistrationMsg(ctxt)
}

func (p *RegisterProcFSM) RegisterUdmGetAmDataCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleRegistrationUdmGetAmDataMsg(ctxt)
}

func (p *RegisterProcFSM) RegisterInitCtxtSetupRespCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleInitCtxtSetupRespMsg4Register(ctxt)
}

// RegisterCmpCallback is callback function for Register complete message event
func (p *RegisterProcFSM) RegisterCmpCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleRegistrationCmpMsg(ctxt)

}

func (p *RegisterProcFSM) RegisterUdmGetAuthCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleAuthRequest(ctxt)
}

func (p *RegisterProcFSM) RegisterUdmGetAuthRespCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleAuthRequestForUdmResp(ctxt)
}

func (p *RegisterProcFSM) RegisterAuthRespCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleAuthResponseMsgRegistration(ctxt)
}

func (p *RegisterProcFSM) RegisterSecModCmpCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleSecModCmpMsg(ctxt)
}

func (p *RegisterProcFSM) RegisterIdentityReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	ctxt := e.Args[0].(context.Context)

	procedure.HandleIdentityReqMsg(ctxt)
}
