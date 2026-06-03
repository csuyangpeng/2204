// Package registration responsible for Registration Procedure
package registration

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
)

// RegisterProcFSM is a wrapper for Procedure State Machine for Register process
type RegisterProcFSM struct {
	fsm.BaseFsm
}

// NewRegisterProcFSM return a RegisterProcFSM pointer after create
// for each sc goroutine, make sure only one fsm
func NewRegisterProcFSM() (*RegisterProcFSM, error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	fsm := fsm.CreateFsm(st.StateRegisterStart)

	registerFSM := &RegisterProcFSM{}
	registerFSM.Bfsm = fsm

	err := registerFSM.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "Failed to create RegisterProcFSM, err:", err)
		return nil, err
	}

	return registerFSM, err
}

// initStateModle register and initialize the FMS
func (p *RegisterProcFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)

	//insert the register management state modle here
	stateModel := []fsm.StateModel{
		{ //register request
			Event:  st.EventRegisterRequest,
			Src:    st.StateRegisterStart,
			Dest:   st.StateRegisterWfInitCtxtSetupResp,
			CbFunc: p.RegisterRequestCallback,
		},
		{ //udm sbi
			Event:  st.EventUdmGetAmData,
			Src:    st.StateRegisterWfUdmGetAmDataResp,
			Dest:   st.StateRegisterWfInitCtxtSetupResp,
			CbFunc: p.RegisterUdmGetAmDataCallback,
		},
		{ // identity request
			Event:  st.EventRegisterIdentityRequest,
			Src:    st.StateRegisterStart,
			Dest:   st.StateRegisterWfIdentityResp,
			CbFunc: p.RegisterIdentityReqCallback,
		},
		{ // register request on aka
			Event:  st.EventAuthRequest,
			Src:    st.StateRegisterAuthStart,
			Dest:   st.StateRegisterWfUdmGetAuth,
			CbFunc: p.RegisterUdmGetAuthCallback,
		},
		{ // register request on aka
			Event:  st.EventRegisterUdmGetAuth,
			Src:    st.StateRegisterWfUdmGetAuth,
			Dest:   st.StateRegisterWfAuthResp,
			CbFunc: p.RegisterUdmGetAuthRespCallback,
		},
		{ // auth response
			Event:  st.EventRegisterAuthResp,
			Src:    st.StateRegisterWfAuthResp,
			Dest:   st.StateRegisterWfSecModCmp,
			CbFunc: p.RegisterAuthRespCallback,
		},
		{ // security command complete
			Event:  st.EventRegisterSecModCmp,
			Src:    st.StateRegisterWfSecModCmp,
			Dest:   st.StateRegisterSecDone,
			CbFunc: p.RegisterSecModCmpCallback,
		},
		{ //initial context setup response
			Event:  st.EventRegisterInitCtxtSetupResp,
			Src:    st.StateRegisterWfInitCtxtSetupResp,
			Dest:   st.StateRegisterWfRegisterCmp,
			CbFunc: p.RegisterInitCtxtSetupRespCallback,
		},
		{ //register complete
			Event:  st.EventRegisterComplete,
			Src:    st.StateRegisterWfRegisterCmp,
			Dest:   st.StateRegisterEnd,
			CbFunc: p.RegisterCmpCallback,
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
