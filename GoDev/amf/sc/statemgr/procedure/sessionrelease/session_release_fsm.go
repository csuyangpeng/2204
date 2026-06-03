package sessionrelease

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
)

type SessRelProcFSM struct {
	fsm.BaseFsm
}

func NewSessRelProcFSM() (*SessRelProcFSM, error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	fsm := fsm.CreateFsm(st.StateAnRelStart)

	anRelFsm := &SessRelProcFSM{}
	anRelFsm.Bfsm = fsm

	err := anRelFsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, "Failed to create AnRelProcFSM, err:", err)
		return nil, err
	}

	return anRelFsm, err
}

func (p *SessRelProcFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	// insert the an release state model here
	stateModel := []fsm.StateModel{
		{
			Event:  st.EventSessRelReq,
			Src:    st.StateSessRelStart,
			Dest:   st.StateSessRelWfUpSmCtxtResp, //smf->amf
			CbFunc: p.SessRelReqCallback,
		},
		{
			Event:  st.EventSessRelUpSmCtxtResp,
			Src:    st.StateSessRelWfUpSmCtxtResp,
			Dest:   st.StateSessRelWfUpSmCtxtReqSec, //ran->amf
			CbFunc: p.UpdateSmCtxtRespCallback,
		},
		//{
		//	Event:  st.EventSessRelN2ResRelResp,
		//	Src:    st.StateSessRelWfN2ResRelResp,
		//	Dest:   st.StateSessRelWfN2UplinkNasTrans,
		//	CbFunc: p.N2ResRelRespCallback,
		//},
		{
			Event:  st.EventSessRelUpSmCtxtReqSec,
			Src:    st.StateSessRelWfUpSmCtxtReqSec,
			Dest:   st.StateSessRelWfUpSmCtxtRespSec, //smf->amf
			CbFunc: p.UpSmCtxtRespSecCallback,
		},
		{
			Event:  st.EventSessRelEnd,
			Src:    st.StateSessRelWfUpSmCtxtRespSec,
			Dest:   st.StateSessRelEnd,
			CbFunc: p.EndCallback,
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
