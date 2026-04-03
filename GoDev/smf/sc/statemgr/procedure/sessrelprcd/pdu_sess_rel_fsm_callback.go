package sessrelprcd

import (
	"context"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/sessmgnt/procedure"
)

func (p *PduSessRelPrcdFSM) PduSessRelReqCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandlePduSessRelReqMsg(ctxt)
}

func (p *PduSessRelPrcdFSM) PduSessRelUpdateSmCtxtReqSecCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandlePduSessRelUpdateSmCtxtReqSecond(ctxt)
}
