package servicerequest

import (
	"lite5gc/cmn/fsm"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)
import (
	st "lite5gc/cmn/types/statetypes"
)

// SerReqFSM is a wrapper for Procedure State Machine for service request procedure
type SerReqProcFSM struct {
	fsm.BaseFsm
}

// NewSerReqProcFSM return a SerReqProcFSM pointer after create
// for each sc goroute, make sure only one fsm
func NewSerReqProcFSM() (*SerReqProcFSM, error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	fsm := fsm.CreateFsm(st.StateSrvReqStart)

	serReqFsm := &SerReqProcFSM{}
	serReqFsm.Bfsm = fsm

	err := serReqFsm.initStateModel()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "Failed to create SerReqProcFSM, err:", err)
		return nil, err
	}

	return serReqFsm, err
}

// initStateModle register and initialize the FMS
func (p *SerReqProcFSM) initStateModel() (err error) {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	//insert the service request state model here
	stateModel := []fsm.StateModel{
		// 1 paging
		{ //paging request
			Event:  st.EventPagingReq,
			Src:    st.StateSrvReqStart,
			Dest:   st.StateSrvReqWfSerReq,
			CbFunc: p.PagingReqCallback,
		},

		{ //service request triggered by paging
			Event:  st.EventSrvReqServiceReqPaging,
			Src:    st.StateSrvReqWfSerReq,
			Dest:   st.StateSrvReqWfUpSmCtxtResp,
			CbFunc: p.SrvReqServiceReqCallback,
		},

		{ //service request triggered by ue
			Event:  st.EventSrvReqServiceReq,
			Src:    st.StateSrvReqStart,
			Dest:   st.StateSrvReqWfUpSmCtxtResp,
			CbFunc: p.SrvReqServiceReqCallback,
		},

		{ // service request on aka
			Event:  st.EventSrvReqServiceReqAuth,
			Src:    st.StateSrvReqAuthStart,
			Dest:   st.StateSrvReqWfAuthResp,
			CbFunc: p.SrvReqAuthCallback,
		},
		{ // auth response
			Event:  st.EventSrvReqAuthResp,
			Src:    st.StateSrvReqWfAuthResp,
			Dest:   st.StateSrvReqWfSecModCmp,
			CbFunc: p.SrvReqAuthRespCallback,
		},
		{ // security command complete
			Event:  st.EventSrvReqSecModCmp,
			Src:    st.StateSrvReqWfSecModCmp,
			Dest:   st.StateSrvReqSecDone,
			CbFunc: p.SrvReqSecModCmpCallback,
		},

		{ //update sm ctxt response
			Event:  st.EventSrvReqUpdateSmCtxtResp,
			Src:    st.StateSrvReqWfUpSmCtxtResp,
			Dest:   st.StateSrvReqWfInitCtxtSetupResp,
			CbFunc: p.SrvReqUpSmCtxtRespCallback,
		},

		{ //initial context setup response
			Event:  st.EventSrvReqInitCtxtSetupResp,
			Src:    st.StateSrvReqWfInitCtxtSetupResp,
			Dest:   st.StateSrvReqWfUpSmCtxtRespSec,
			CbFunc: p.SrvReqInitCtxtSetupRespCallback,
		},
		{ //update sm ctxt response second
			Event:  st.EventSrvReqUpdateSmCtxtRespSec,
			Src:    st.StateSrvReqWfUpSmCtxtRespSec,
			Dest:   st.StateSrvReqEnd,
			CbFunc: p.SrvReqUpSmCtxtRespSecCallback,
		},
		{
			Event:  st.EventSrvReqEnd,
			Src:    st.StateSrvReqEnd,
			Dest:   st.EventSrvReqEnd,
			CbFunc: p.SrvReqFinishedCallback,
		},
		{ //initial context setup response
			Event:  st.EventSrvReqInitCtxtSetupRespNoPsi,
			Src:    st.StateSrvReqWfInitCtxtSetupResp,
			Dest:   st.StateSrvReqEnd,
			CbFunc: p.SrvReqInitCtxtSetupRespCallback,
		},
	}
	//add the event dest and call back into the fsm
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
