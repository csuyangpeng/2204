package anrelease

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
)

// AnRelFSM is a wrapper for Procedure State Machine for connection management process
type AnRelProcFSM struct {
	fsm.BaseFsm
}

// NewCmProcFSM return a CmProcFSM pointer after create
// for each sc go route, make sure only one fsm
func NewAnRelProcFSM() (*AnRelProcFSM, error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	fsm := fsm.CreateFsm(st.StateAnRelStart)

	anRelFsm := &AnRelProcFSM{}
	anRelFsm.Bfsm = fsm

	err := anRelFsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "Failed to create AnRelProcFSM, err:", err)
		return nil, err
	}

	return anRelFsm, err
}

// initStateModel of AnRelease and initialize the FMS
func (p *AnRelProcFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	// insert the an release state model here
	stateModel := []fsm.StateModel{
		{
			Event:  st.EventAnRelUeCtxtRelReq,
			Src:    st.StateAnRelStart,
			Dest:   st.StateAnRelWfRelCmp,
			CbFunc: p.UeCtxtRelReqCallback,
		},
		{
			Event:  st.EventAnRelUpdateSmCtxtAck,
			Src:    st.StateAnRelWfUpSmCtxtAck,
			Dest:   st.StateAnRelWfRelCmp,
			CbFunc: p.UpdateSmCtxtAckCallback,
		},
		{
			Event:  st.EventAnRelUeCtxtRelCmp,
			Src:    st.StateAnRelWfRelCmp,
			Dest:   st.StateAnRelWfEnd,
			CbFunc: p.UeCtxtRelCmpCallback,
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
