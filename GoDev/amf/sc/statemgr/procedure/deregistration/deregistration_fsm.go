package deregistration

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
)

// DeRegisterProcFSM is a wrapper for Procedure State Machine for ue DeRegister process
type DeRegisterProcFSM struct {
	fsm.BaseFsm
}

// NewRegisterProcFSM return a DeRegisterProcFSM pointer after create
// for each sc goroute, make sure only one fsm
func NewDeRegisterProcFSM() (*DeRegisterProcFSM, error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	fsm := fsm.CreateFsm(st.StateDeRegisterStart)

	deRegistfsm := &DeRegisterProcFSM{}
	deRegistfsm.Bfsm = fsm

	err := deRegistfsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "Failed to create DeRegisterFSM, err:", err)
		return nil, err
	}

	return deRegistfsm, err
}

// initStateModle register and initialize the FMS
func (p *DeRegisterProcFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	//insert the register management state model here
	stateModel := []fsm.StateModel{
		{
			Event:  st.EventDeRegisterRequest,
			Src:    st.StateDeRegisterStart,
			Dest:   st.StateDeRegisterWfRelSmCtxtResp,
			CbFunc: p.DeRegisterReqCallback,
		},
		{
			Event:  st.EventDeRegisterRelSmCtxtResp,
			Src:    st.StateDeRegisterWfRelSmCtxtResp,
			Dest:   st.StateDeRegisterWfUeCtxtRelCmp,
			CbFunc: p.DeRegisterRelSmCtxtCallback,
		},
		{
			Event:  st.EventDeRegisterUeCtxtRelCmp,
			Src:    st.StateDeRegisterWfUeCtxtRelCmp,
			Dest:   st.StateDeRegisterEnd,
			CbFunc: p.DeRegisterRelUeCtxtRelCmpCallback,
		},
		{
			Event:  st.EventDeRegisterRequestBySwitchOff,
			Src:    st.StateDeRegisterStart,
			Dest:   st.StateDeRegisterWfRelSmCtxtResp,
			CbFunc: p.DeRegisterReqBySwitchOffCallback,
		},
	}

	//add the eventdest and call back into the fsm
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
