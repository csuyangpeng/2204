package sessserreqprcd

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

type PduSessSerReqPrcdFSM struct {
	fsm.BaseFsm
}

// NewSessionFSM return a PduSessionFSM pointer, which is a Session state machine
// for each sc goroute, make sure only one fsm
func NewPduSessSerReqPrcdFSM() (*PduSessSerReqPrcdFSM, error) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	createFsm := fsm.CreateFsm(statetype.StatePduSessSerReqStart)

	sessFsm := &PduSessSerReqPrcdFSM{}
	sessFsm.Bfsm = createFsm

	err := sessFsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleSmfState, rlogger.FATAL, nil, "failed to init state modle, err: %s", err)
		return nil, err
	}

	return sessFsm, nil
}

func (p *PduSessSerReqPrcdFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//insert the pdu session state model here
	stateModel := []fsm.StateModel{
		{
			Event:  statetype.EventPduSessSerReqStart,
			Src:    statetype.StatePduSessSerReqStart,
			Dest:   statetype.StatePduSessSerReqWfN4ModResp,
			CbFunc: p.PduSessSerReqStartCb,
		},
		{
			Event:  statetype.EventPduSessSerReqN4ModifyResp,
			Src:    statetype.StatePduSessSerReqWfN4ModResp,
			Dest:   statetype.StatePduSessSerReqWfUpdateSmCtxtReqSec,
			CbFunc: p.PduSessSerReqN4ModRespCb,
		},
		{
			Event:  statetype.EventPduSessSerReqUpdateSmCtxtReqSec,
			Src:    statetype.StatePduSessSerReqWfUpdateSmCtxtReqSec,
			Dest:   statetype.StatePduSessSerReqWfN4ModRespSec,
			CbFunc: p.PduSessSerReqUpdateSmCtxtReqSecCb,
		},
		{
			Event:  statetype.EventPduSessSerReqN4ModifyRespSec,
			Src:    statetype.StatePduSessSerReqWfN4ModRespSec,
			Dest:   statetype.StatePduSessSerReqEnd,
			CbFunc: p.PduSessSerReqN4ModRespSecCb,
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
