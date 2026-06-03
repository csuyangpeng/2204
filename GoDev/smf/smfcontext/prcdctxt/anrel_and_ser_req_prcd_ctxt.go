package prcdctxt

import (
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

type AnRelSerReqPrcdCtxt struct {
	BaseCtxt
	PduSessId           nas.PduSessID
	UpCnxState          n11msg.UpCnxState
	IsAnRelease         bool
	IsWfSecN4ModifyResp bool
	SbiMessage          *sbicmn.SbiHandlerMessage
}

func NewAnRelSerReqPrcdCtxt(psi nas.PduSessID) *AnRelSerReqPrcdCtxt {
	prcdCtxt := &AnRelSerReqPrcdCtxt{}
	prcdCtxt.PduSessId = psi
	return prcdCtxt
}

// SetState set the procedure state
func (p *AnRelSerReqPrcdCtxt) SetNextState(state string) error {
	err := statetype.PduSessSerReqCheckState(state)
	if err != nil {
		return err
	}
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, nil, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}
