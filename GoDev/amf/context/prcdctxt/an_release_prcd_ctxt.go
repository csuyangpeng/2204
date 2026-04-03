package prcdctxt

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

// DeRegistration is procedure context for DeRegistration procedure
type AnReleasePrcdCtxt struct {
	BaseCtxt
	psiCounter     int
	ReleasePsiList []uint8
	UpCnxState     types3gpp.PduSessStatus
}

func (p *AnReleasePrcdCtxt) SetCounter(c int) {
	p.psiCounter = c
}

func (p *AnReleasePrcdCtxt) GetCounter() int {
	return p.psiCounter
}

// SetState set the procedure state
func (p *AnReleasePrcdCtxt) SetNextState(state string) error {
	err := st.AnRelCheckState(state)
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "state (%s) is not belong to AnReleasePrcdCtxt", state)
		return err
	}
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}

func NewAnReleasePrcdCtxtUETrigger() (anReleaseCtxt *AnReleasePrcdCtxt) {
	anReleaseCtxt = &AnReleasePrcdCtxt{}
	anReleaseCtxt.status = st.StateAnRelStart
	return anReleaseCtxt
}

func NewAnReleasePrcdCtxtNetTrigger() (anReleaseCtxt *AnReleasePrcdCtxt) {
	anReleaseCtxt = &AnReleasePrcdCtxt{}
	anReleaseCtxt.status = st.StateAnRelWfUpSmCtxtAck
	return anReleaseCtxt
}

func NewAnReleasePrcdCtxtAfterReg() (anReleaseCtxt *AnReleasePrcdCtxt) {
	anReleaseCtxt = &AnReleasePrcdCtxt{}
	anReleaseCtxt.status = st.StateAnRelWfRelCmp
	return anReleaseCtxt
}
