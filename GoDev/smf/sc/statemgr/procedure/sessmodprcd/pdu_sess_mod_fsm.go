package sessmodprcd

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

type PduSessModPrcdFSM struct {
	fsm.BaseFsm
}

func NewPduSessModPrcdFSM() (*PduSessModPrcdFSM, error) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	createFsm := fsm.CreateFsm(statetype.StatePduSessModStart)

	sessFsm := &PduSessModPrcdFSM{}
	sessFsm.Bfsm = createFsm

	err := sessFsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleSmfState, rlogger.FATAL, nil, "failed to init state model, err: %s", err)
		return nil, err
	}

	return sessFsm, nil
}

func (p *PduSessModPrcdFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	//insert the pdu session state model here
	stateModel := []fsm.StateModel{
		{
			Event:  statetype.EventPduSessMod1stUpdateSmCtxtResp,
			Src:    statetype.StatePduSessModWf1stN4SessModResp,
			Dest:   statetype.StatePduSessModWf2ndUpdateSMCtxtReq,
			CbFunc: p.PduSessN4SessMod1stN4SessModRespCb,
		},
		{
			Event:  statetype.EventPduSessMod1stUpdateN4SmCtxtReq,
			Src:    statetype.StatePduSessModWf1stN4SessModResp,
			Dest:   statetype.StatePduSessModWf2ndUpdateSMCtxtReq,
			CbFunc: p.PduSessN4SessMod1stUpdateN4SmCtxtReqCb,
		},
		{
			Event:  statetype.EventPduSessMod2ndUpdateSmCtxtResp,
			Src:    statetype.StatePduSessModWf2ndUpdateSMCtxtReq,
			Dest:   statetype.StatePduSessModWf2ndN4SessModResp,
			CbFunc: p.PduSessN4SessMod2ndN4SessModRespCb,
		},
		{
			Event:  statetype.EventPduSessMod2ndUpdateN4SmCtxtReq,
			Src:    statetype.StatePduSessModWf3rdUpdateSMCtxtReq,
			Dest:   statetype.StatePduSessModWf3rdN4SessModResp,
			CbFunc: p.PduSessN4SessMod2ndUpdateN4SmCtxtReqCb,
		},
		{
			Event:  statetype.EventPduSessMod3rdUpdateSmCtxtResp,
			Src:    statetype.StatePduSessModWf3rdN4SessModResp,
			Dest:   statetype.StatePduSessModEnd,
			CbFunc: p.PduSessN4SessMod3rdN4SessModRespCb,
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
