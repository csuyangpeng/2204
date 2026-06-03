package sessionmodification

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
)

type SessModProcFSM struct {
	fsm.BaseFsm
}

func NewSessModProcFSM() (*SessModProcFSM, error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	fsm := fsm.CreateFsm(st.StateSessModStart)

	anModFsm := &SessModProcFSM{}
	anModFsm.Bfsm = fsm

	err := anModFsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, "Failed to create AnRelProcFSM, err:", err)
		return nil, err
	}

	return anModFsm, err
}

func (p *SessModProcFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	// insert the an modification state model here
	stateModel := []fsm.StateModel{
		{
			Event:  st.EventSessModWf1stUpSmCtxtReq,
			Src:    st.StateSessModStart,
			Dest:   st.StateSessModWf1stUpSmCtxtResp,
			CbFunc: p.SessModWf1stUpSmCtxtReqCallback,
		},
		{
			Event:  st.EventSessModWfN2SessReq,
			Src:    st.StateSessModWf1stUpSmCtxtResp,
			Dest:   st.StateSessModWfN2SessResp,
			CbFunc: p.SessModWfN2SessReqCallback,
		},
		{
			Event:  st.EventSessModWf2ndUpSmCtxtReq,
			Src:    st.StateSessModWfN2SessResp,
			Dest:   st.StateSessModWf2ndUpSmCtxtResp,
			CbFunc: p.SessModWf2ndUpSmCtxtReqCallback,
		},
		{
			Event:  st.EventSessModWfN2NASUplinkTrans,
			Src:    st.StateSessModWf2ndUpSmCtxtResp,
			Dest:   st.StateSessModEnd,
			CbFunc: p.SessModWfN2NASUplinkTransCallback,
		},
		{
			Event:  st.EventSessModWf3rdUpSmCtxtReq,
			Src:    st.StateSessModWfN2NASUplinkTrans,
			Dest:   st.StateSessModWf3rdUpSmCtxtResp,
			CbFunc: p.SessModWf3rdUpSmCtxtReqCallback,
		},
		{
			Event:  st.EventSessModEnd,
			Src:    st.StateSessModWf3rdUpSmCtxtResp,
			Dest:   st.StateSessModEnd,
			CbFunc: p.SessModEndCallback,
		},
	}

	// add the Event / Source / Destination and CallBack Function into the fsm
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
