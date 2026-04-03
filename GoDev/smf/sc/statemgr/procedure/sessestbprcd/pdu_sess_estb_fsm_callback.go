package sessestbprcd

import (
	"context"
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/sessmgnt/procedure"
)

func (p *PduSessEstbPrcdFSM) PduSessEstbReqCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandlePduSessEstbReqMsg(ctxt)
}

func (p *PduSessEstbPrcdFSM) PduSessEstbReqUdmSmDataCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandlePduSessEstbUdmGetRespMsg(ctxt)
}

func (p *PduSessEstbPrcdFSM) PduSessEstbN1N2MsgTransferRespCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleN1N2MsgTransferRespMsg4PduSessEstb(ctxt)
}

func (p *PduSessEstbPrcdFSM) PduSessN4SessEstbRespCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleN4SessEstbResp(ctxt)
}

func (p *PduSessEstbPrcdFSM) PduSessUpdateSmCtxtReqCb(e *fsm.Event) {
	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleUpdateSmCtxtRequestMsg4Estb(ctxt)
}

func (p *PduSessEstbPrcdFSM) PduSessN4SessModRespCb(e *fsm.Event) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//get the context from fsm.Event
	ctxt := e.Args[0].(context.Context)

	procedure.HandleN4SessModResp(ctxt)
}
