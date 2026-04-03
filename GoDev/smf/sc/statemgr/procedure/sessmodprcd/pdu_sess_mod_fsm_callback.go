package sessmodprcd

import (
	"context"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/sessmgnt/procedure"
)

func (p *PduSessModPrcdFSM) PduSessN4SessMod1stN4SessModRespCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.Handle1stN4SessModResp(ctxt)
}

func (p *PduSessModPrcdFSM) PduSessN4SessMod2ndN4SessModRespCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	//ctxt := e.Args[0].(context.Context)

	//procedure.Handle2ndN4SessModResp(ctxt)
}

func (p *PduSessModPrcdFSM) PduSessN4SessMod3rdN4SessModRespCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	//ctxt := e.Args[0].(context.Context)

	//procedure.Handle3rdN4SessModResp(ctxt)
}

func (p *PduSessModPrcdFSM) PduSessN4SessMod1stUpdateN4SmCtxtReqCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	//ctxt := e.Args[0].(context.Context)

	//procedure.Handle1stUpdateN4SmCtxtReq(ctxt)
}

func (p *PduSessModPrcdFSM) PduSessN4SessMod2ndUpdateN4SmCtxtReqCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	//ctxt := e.Args[0].(context.Context)

	//procedure.Handle2ndUpdateN4SmCtxtReq(ctxt)
}

func (p *PduSessModPrcdFSM) PduSessEndCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	//ctxt := e.Args[0].(context.Context)
}
