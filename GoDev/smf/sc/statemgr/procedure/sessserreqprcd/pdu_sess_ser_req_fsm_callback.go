package sessserreqprcd

import (
	"context"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/sessmgnt/procedure"
)

func (p *PduSessSerReqPrcdFSM) PduSessSerReqStartCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleUpdateSmCtxtRequestMsg4SerReqPrcd(ctxt)
}

func (p *PduSessSerReqPrcdFSM) PduSessSerReqN4ModRespCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleN4SessModResp4SerReq(ctxt)
}

func (p *PduSessSerReqPrcdFSM) PduSessSerReqUpdateSmCtxtReqSecCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleUpdateSmCtxtRequestSecMsg4SerReqPrcd(ctxt)
}

func (p *PduSessSerReqPrcdFSM) PduSessSerReqN4ModRespSecCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleN4SessModRespSec4SerReq(ctxt)
}
