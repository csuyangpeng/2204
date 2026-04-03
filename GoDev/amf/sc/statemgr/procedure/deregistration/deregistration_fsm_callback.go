package deregistration

import (
	"context"
	"lite5gc/amf/sc/mobmgnt/procedure"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *DeRegisterProcFSM) DeRegisterReqCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	//handle deregistration request
	procedure.HandleDeRegistrationRequest(ctxt)
}

// RegisterRequestCallback is callback function for register request message event
func (p *DeRegisterProcFSM) DeRegisterRelSmCtxtCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	// handle release sm context response
	procedure.HandleRelSmCtxtResp4Deregist(ctxt)
}

// RegisterRequestCallback is callback function for register request message event
func (p *DeRegisterProcFSM) DeRegisterRelUeCtxtRelCmpCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	// handle release sm context response
	procedure.HandleUeCtxtRelCmp4Deregist(ctxt)
}

// RegisterRequestCallback is callback function for register request message event
func (p *DeRegisterProcFSM) DeRegisterReqBySwitchOffCallback(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	// handle release sm context response
	procedure.HandleUeCtxtReqBySwitchOff(ctxt)
}
