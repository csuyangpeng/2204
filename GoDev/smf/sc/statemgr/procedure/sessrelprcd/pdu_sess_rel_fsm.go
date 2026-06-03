package sessrelprcd

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

type PduSessRelPrcdFSM struct {
	fsm.BaseFsm
}

// NewSessionFSM return a PduSessionFSM pointer, which is a Session state machine
// for each sc goroute, make sure only one fsm
func NewPduSessRelPrcdFSM() (*PduSessRelPrcdFSM, error) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	createFsm := fsm.CreateFsm(statetype.StatePduSessRelStart)

	sessFsm := &PduSessRelPrcdFSM{}
	sessFsm.Bfsm = createFsm

	err := sessFsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleSmfState, rlogger.FATAL, nil, "failed to init state modle, err: %s", err)
		return nil, err
	}

	return sessFsm, nil
}

func (p *PduSessRelPrcdFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//insert the pdu session state model here
	stateModel := []fsm.StateModel{
		{
			Event:  statetype.EventPduSessRelReq,
			Src:    statetype.StatePduSessRelStart,
			Dest:   statetype.StatePduSessRelWfUpdateSmCtxtReqSec,
			CbFunc: p.PduSessRelReqCb,
		},
		{
			Event:  statetype.EventPduSessRelUpdateSmCtxtReqSec,
			Src:    statetype.StatePduSessRelWfUpdateSmCtxtReqSec,
			Dest:   statetype.StatePduSessRelEnd,
			CbFunc: p.PduSessRelUpdateSmCtxtReqSecCb,
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
