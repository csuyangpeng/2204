package sessestbprcd

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

type PduSessEstbPrcdFSM struct {
	fsm.BaseFsm
}

// NewSessionFSM return a PduSessionFSM pointer, which is a Session state machine
// for each sc goroute, make sure only one fsm
func NewPduSessEstbPrcdFSM() (*PduSessEstbPrcdFSM, error) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	createFsm := fsm.CreateFsm(statetype.StatePduSessEstbStart)

	sessFsm := &PduSessEstbPrcdFSM{}
	sessFsm.Bfsm = createFsm

	err := sessFsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleSmfState, rlogger.FATAL, nil, "failed to init state model, err: %s", err)
		return nil, err
	}

	return sessFsm, nil
}

func (p *PduSessEstbPrcdFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//insert the pdu session state model here
	stateModel := []fsm.StateModel{
		{
			Event:  statetype.EventPduSessEstbReq,
			Src:    statetype.StatePduSessEstbStart,
			Dest:   statetype.StatePduSessEstbWfN4SessEstbResp,
			CbFunc: p.PduSessEstbReqCb,
		},
		{
			Event:  statetype.EventPduSessEstbSmGetAck,
			Src:    statetype.StatePduSessEstbWfSmGet,
			Dest:   statetype.StatePduSessEstbWfN4SessEstbResp,
			CbFunc: p.PduSessEstbReqUdmSmDataCb,
		},
		{
			Event:  statetype.EventPduSessEstbN4SessEstbResp,
			Src:    statetype.StatePduSessEstbWfN4SessEstbResp,
			Dest:   statetype.StatePduSessEstbWfN1N2MsgTransferResp,
			CbFunc: p.PduSessN4SessEstbRespCb,
		},
		{
			Event:  statetype.EventPduSessEstbN1N2MsgTransferResp,
			Src:    statetype.StatePduSessEstbWfN1N2MsgTransferResp,
			Dest:   statetype.StatePduSessEstbWfUpdateSmCtxtReq,
			CbFunc: p.PduSessEstbN1N2MsgTransferRespCb,
		},
		{
			Event:  statetype.EventPduSessEstbUpdateSmCtxtReq,
			Src:    statetype.StatePduSessEstbWfUpdateSmCtxtReq,
			Dest:   statetype.StatePduSessEstbWfN4SessModResp,
			CbFunc: p.PduSessUpdateSmCtxtReqCb,
		},
		{
			Event:  statetype.EventPduSessEstbN4SessModResp,
			Src:    statetype.StatePduSessEstbWfN4SessModResp,
			Dest:   statetype.StatePduSessEstbEnd,
			CbFunc: p.PduSessN4SessModRespCb,
		},
	}

	//add the event, src, dest and callback into the fsm
	for _, sm := range stateModel {
		err = p.RegisterEvent(sm.Event,
			[]string{sm.Src},
			sm.Dest,
			sm.CbFunc)
		if err != nil {
			return
		}
	}

	return
}
