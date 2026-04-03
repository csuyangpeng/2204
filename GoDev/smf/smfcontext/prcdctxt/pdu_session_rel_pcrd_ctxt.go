package prcdctxt

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

type PduSessionRelReqPrcdCtxt struct {
	BaseCtxt

	Imsi             types3gpp.Imsi
	PduSessId        nas.PduSessID
	Pti              nas.PrcdTransID
	Cause            nas.Sm5gCause
	IsCauseExist     bool
	IsDeRegisterPrcd bool

	TimerMgr                *timermgr.TimerMgr
	TimerId                 int64
	ReTransRelAccetpCoutner uint8

	SbiMessage *sbicmn.SbiHandlerMessage
}

func NewPduSessionRelReqPrcdCtxt(psi nas.PduSessID) *PduSessionRelReqPrcdCtxt {
	procCtxt := &PduSessionRelReqPrcdCtxt{}
	procCtxt.PduSessId = psi
	procCtxt.status = statetype.StatePduSessRelStart
	return procCtxt
}

// SetState set the procedure state
func (p *PduSessionRelReqPrcdCtxt) SetNextState(state string) error {
	err := statetype.PduSessRelCheckState(state)
	if err != nil {
		return err
	}
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, p, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}
